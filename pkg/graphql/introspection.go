package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
)

// IntrospectionClient handles GraphQL schema introspection
type IntrospectionClient struct {
	client *Client
	cache  *schemaCache
}

// schemaCache provides caching for introspection results
type schemaCache struct {
	mu      sync.RWMutex
	schemas map[string]*cachedSchema
	ttl     time.Duration
	enabled bool
}

// cachedSchema represents a cached schema with expiration
type cachedSchema struct {
	schema    *Schema
	fetchedAt time.Time
}

// Schema represents a GraphQL schema from introspection
type Schema struct {
	QueryType        *Type  `json:"queryType"`
	MutationType     *Type  `json:"mutationType,omitempty"`
	SubscriptionType *Type  `json:"subscriptionType,omitempty"`
	Types            []Type `json:"types"`
	Directives       []Type `json:"directives"`
}

// Type represents a GraphQL type
type Type struct {
	Kind          string       `json:"kind"`
	Name          string       `json:"name,omitempty"`
	Description   string       `json:"description,omitempty"`
	Fields        []Field      `json:"fields,omitempty"`
	InputFields   []InputValue `json:"inputFields,omitempty"`
	Interfaces    []Type       `json:"interfaces,omitempty"`
	EnumValues    []EnumValue  `json:"enumValues,omitempty"`
	PossibleTypes []Type       `json:"possibleTypes,omitempty"`
	OfType        *Type        `json:"ofType,omitempty"`
}

// Field represents a GraphQL field
type Field struct {
	Name              string       `json:"name"`
	Description       string       `json:"description,omitempty"`
	Args              []InputValue `json:"args"`
	Type              Type         `json:"type"`
	IsDeprecated      bool         `json:"isDeprecated"`
	DeprecationReason string       `json:"deprecationReason,omitempty"`
}

// InputValue represents a GraphQL input value (argument or input field)
type InputValue struct {
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Type         Type   `json:"type"`
	DefaultValue string `json:"defaultValue,omitempty"`
}

// EnumValue represents a GraphQL enum value
type EnumValue struct {
	Name              string `json:"name"`
	Description       string `json:"description,omitempty"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason,omitempty"`
}

// IntrospectionOptions configures introspection behavior
type IntrospectionOptions struct {
	CacheEnabled bool
	CacheTTL     time.Duration
}

// DefaultIntrospectionOptions returns default introspection options
func DefaultIntrospectionOptions() IntrospectionOptions {
	return IntrospectionOptions{
		CacheEnabled: true,
		CacheTTL:     5 * time.Minute,
	}
}

// NewIntrospectionClient creates a new introspection client
func NewIntrospectionClient(client *Client, options ...IntrospectionOptions) *IntrospectionClient {
	opts := DefaultIntrospectionOptions()
	if len(options) > 0 {
		opts = options[0]
	}

	cache := &schemaCache{
		schemas: make(map[string]*cachedSchema),
		ttl:     opts.CacheTTL,
		enabled: opts.CacheEnabled,
	}

	return &IntrospectionClient{
		client: client,
		cache:  cache,
	}
}

// GetSchema fetches the GraphQL schema using introspection
func (ic *IntrospectionClient) GetSchema(ctx context.Context) (*Schema, error) {
	endpoint := ic.client.GetEndpoint()

	// Check cache first
	if ic.cache.enabled {
		if schema := ic.cache.get(endpoint); schema != nil {
			return schema, nil
		}
	}

	// Build introspection query
	query := ic.buildIntrospectionQuery()

	// Execute introspection query
	response, err := ic.client.Query(ctx, query, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to execute introspection query: %w", err)
	}

	// Check for GraphQL errors
	if len(response.Errors) > 0 {
		return nil, fmt.Errorf("introspection query returned errors: %v", response.Errors)
	}

	// Parse schema from response
	schema, err := ic.parseSchemaResponse(response.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse introspection response: %w", err)
	}

	// Cache the result
	if ic.cache.enabled {
		ic.cache.set(endpoint, schema)
	}

	return schema, nil
}

// buildIntrospectionQuery builds the standard GraphQL introspection query
func (ic *IntrospectionClient) buildIntrospectionQuery() string {
	return `
		query IntrospectionQuery {
			__schema {
				queryType { name }
				mutationType { name }
				subscriptionType { name }
				types {
					...FullType
				}
				directives {
					name
					description
					locations
					args {
						...InputValue
					}
				}
			}
		}

		fragment FullType on __Type {
			kind
			name
			description
			fields(includeDeprecated: true) {
				name
				description
				args {
					...InputValue
				}
				type {
					...TypeRef
				}
				isDeprecated
				deprecationReason
			}
			inputFields {
				...InputValue
			}
			interfaces {
				...TypeRef
			}
			enumValues(includeDeprecated: true) {
				name
				description
				isDeprecated
				deprecationReason
			}
			possibleTypes {
				...TypeRef
			}
		}

		fragment InputValue on __InputValue {
			name
			description
			type { ...TypeRef }
			defaultValue
		}

		fragment TypeRef on __Type {
			kind
			name
			ofType {
				kind
				name
				ofType {
					kind
					name
					ofType {
						kind
						name
						ofType {
							kind
							name
							ofType {
								kind
								name
								ofType {
									kind
									name
									ofType {
										kind
										name
									}
								}
							}
						}
					}
				}
			}
		}
	`
}

// parseSchemaResponse parses the introspection response into a Schema
func (ic *IntrospectionClient) parseSchemaResponse(data json.RawMessage) (*Schema, error) {
	var introspectionResult struct {
		Schema Schema `json:"__schema"`
	}

	if err := json.Unmarshal(data, &introspectionResult); err != nil {
		return nil, fmt.Errorf("failed to unmarshal introspection response: %w", err)
	}

	return &introspectionResult.Schema, nil
}

// ValidateOperation validates a GraphQL operation against the schema
func (ic *IntrospectionClient) ValidateOperation(ctx context.Context, operation string) error {
	schema, err := ic.GetSchema(ctx)
	if err != nil {
		// Gracefully handle introspection failures (Requirement 6.5)
		return fmt.Errorf("schema validation unavailable: %w", err)
	}

	// Create validator and validate operation
	validator := NewSchemaValidator(schema)
	result := validator.ValidateOperation(operation)

	if !result.Valid {
		return ic.formatValidationErrors(result.Errors)
	}

	return nil
}

// ValidateOperationDetailed validates a GraphQL operation and returns detailed validation results
func (ic *IntrospectionClient) ValidateOperationDetailed(ctx context.Context, operation string) (*ValidationResult, error) {
	schema, err := ic.GetSchema(ctx)
	if err != nil {
		// Gracefully handle introspection failures (Requirement 6.5)
		result := &ValidationResult{Valid: false}
		result.AddError(ErrorTypeSchemaUnavailable, fmt.Sprintf("Schema introspection failed: %v", err), "")
		return result, nil
	}

	// Create validator and validate operation
	validator := NewSchemaValidator(schema)
	return validator.ValidateOperation(operation), nil
}

// ValidateFieldPath validates a specific field path against the schema
func (ic *IntrospectionClient) ValidateFieldPath(ctx context.Context, path string, operationType string) error {
	schema, err := ic.GetSchema(ctx)
	if err != nil {
		// Gracefully handle introspection failures (Requirement 6.5)
		return fmt.Errorf("schema validation unavailable: %w", err)
	}

	// Create validator and validate field path
	validator := NewSchemaValidator(schema)
	result := validator.ValidateFieldPath(path, operationType)

	if !result.Valid {
		return ic.formatValidationErrors(result.Errors)
	}

	return nil
}

// ValidateFieldPathDetailed validates a field path and returns detailed validation results
func (ic *IntrospectionClient) ValidateFieldPathDetailed(ctx context.Context, path string, operationType string) (*ValidationResult, error) {
	schema, err := ic.GetSchema(ctx)
	if err != nil {
		// Gracefully handle introspection failures (Requirement 6.5)
		result := &ValidationResult{Valid: false}
		result.AddError(ErrorTypeSchemaUnavailable, fmt.Sprintf("Schema introspection failed: %v", err), "")
		return result, nil
	}

	// Create validator and validate field path
	validator := NewSchemaValidator(schema)
	return validator.ValidateFieldPath(path, operationType), nil
}

// formatValidationErrors formats validation errors into a single error message
func (ic *IntrospectionClient) formatValidationErrors(errors []ValidationError) error {
	if len(errors) == 0 {
		return nil
	}

	var messages []string
	for _, err := range errors {
		messages = append(messages, err.Error())
	}

	return fmt.Errorf("schema validation failed: %s", strings.Join(messages, "; "))
}

// GetTypeByName finds a type in the schema by name
func (s *Schema) GetTypeByName(name string) *Type {
	for i, t := range s.Types {
		if t.Name == name {
			return &s.Types[i]
		}
	}
	return nil
}

// GetFieldByName finds a field in a type by name
func (t *Type) GetFieldByName(name string) *Field {
	for i, f := range t.Fields {
		if f.Name == name {
			return &t.Fields[i]
		}
	}
	return nil
}

// IsScalarType checks if the type is a scalar type
func (t *Type) IsScalarType() bool {
	return t.Kind == "SCALAR"
}

// IsObjectType checks if the type is an object type
func (t *Type) IsObjectType() bool {
	return t.Kind == "OBJECT"
}

// IsListType checks if the type is a list type
func (t *Type) IsListType() bool {
	return t.Kind == "LIST"
}

// IsNonNullType checks if the type is a non-null type
func (t *Type) IsNonNullType() bool {
	return t.Kind == "NON_NULL"
}

// Cache methods

// get retrieves a schema from cache if it exists and is not expired
func (c *schemaCache) get(endpoint string) *Schema {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cached, exists := c.schemas[endpoint]
	if !exists {
		return nil
	}

	// Check if cache entry is expired
	if time.Since(cached.fetchedAt) > c.ttl {
		return nil
	}

	return cached.schema
}

// set stores a schema in cache
func (c *schemaCache) set(endpoint string, schema *Schema) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.schemas[endpoint] = &cachedSchema{
		schema:    schema,
		fetchedAt: time.Now(),
	}
}

// clear removes all cached schemas
func (c *schemaCache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.schemas = make(map[string]*cachedSchema)
}

// ClearCache clears the schema cache
func (ic *IntrospectionClient) ClearCache() {
	ic.cache.clear()
}

// SetCacheEnabled enables or disables schema caching
func (ic *IntrospectionClient) SetCacheEnabled(enabled bool) {
	ic.cache.mu.Lock()
	defer ic.cache.mu.Unlock()
	ic.cache.enabled = enabled
}

// SetCacheTTL sets the cache time-to-live duration
func (ic *IntrospectionClient) SetCacheTTL(ttl time.Duration) {
	ic.cache.mu.Lock()
	defer ic.cache.mu.Unlock()
	ic.cache.ttl = ttl
}

