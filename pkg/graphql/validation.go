package graphql

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidationError represents a schema validation error
type ValidationError struct {
	Type    ValidationErrorType    `json:"type"`
	Message string                 `json:"message"`
	Field   string                 `json:"field,omitempty"`
	Path    string                 `json:"path,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// ValidationErrorType represents the type of validation error
type ValidationErrorType string

const (
	ErrorTypeFieldNotFound     ValidationErrorType = "field_not_found"
	ErrorTypeTypeNotFound      ValidationErrorType = "type_not_found"
	ErrorTypeIncompatibleType  ValidationErrorType = "incompatible_type"
	ErrorTypeInvalidOperation  ValidationErrorType = "invalid_operation"
	ErrorTypeSchemaUnavailable ValidationErrorType = "schema_unavailable"
	ErrorTypeSyntaxError       ValidationErrorType = "syntax_error"
)

// Error implements the error interface
func (ve ValidationError) Error() string {
	if ve.Field != "" {
		return fmt.Sprintf("%s: %s (field: %s)", ve.Type, ve.Message, ve.Field)
	}
	return fmt.Sprintf("%s: %s", ve.Type, ve.Message)
}

// ValidationResult represents the result of schema validation
type ValidationResult struct {
	Valid  bool              `json:"valid"`
	Errors []ValidationError `json:"errors,omitempty"`
}

// AddError adds a validation error to the result
func (vr *ValidationResult) AddError(errorType ValidationErrorType, message string, field string) {
	vr.Valid = false
	vr.Errors = append(vr.Errors, ValidationError{
		Type:    errorType,
		Message: message,
		Field:   field,
	})
}

// AddDetailedError adds a validation error with additional details
func (vr *ValidationResult) AddDetailedError(errorType ValidationErrorType, message string, field string, details map[string]interface{}) {
	vr.Valid = false
	vr.Errors = append(vr.Errors, ValidationError{
		Type:    errorType,
		Message: message,
		Field:   field,
		Details: details,
	})
}

// SchemaValidator provides GraphQL schema validation functionality
type SchemaValidator struct {
	schema *Schema
}

// NewSchemaValidator creates a new schema validator
func NewSchemaValidator(schema *Schema) *SchemaValidator {
	return &SchemaValidator{
		schema: schema,
	}
}

// ValidateOperation validates a GraphQL operation against the schema
func (sv *SchemaValidator) ValidateOperation(operation string) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if sv.schema == nil {
		result.AddError(ErrorTypeSchemaUnavailable, "Schema is not available for validation", "")
		return result
	}

	// Parse operation type and validate
	operationType, err := sv.extractOperationType(operation)
	if err != nil {
		result.AddError(ErrorTypeSyntaxError, err.Error(), "")
		return result
	}

	// Validate operation type exists in schema
	if err := sv.validateOperationType(operationType, result); err != nil {
		return result
	}

	// Extract and validate fields
	fields, err := sv.extractFields(operation)
	if err != nil {
		result.AddError(ErrorTypeSyntaxError, err.Error(), "")
		return result
	}

	// Validate each field against the schema
	sv.validateFields(fields, operationType, result)

	return result
}

// extractOperationType extracts the operation type (query, mutation, subscription) from the operation
func (sv *SchemaValidator) extractOperationType(operation string) (string, error) {
	// Remove comments and normalize whitespace
	cleaned := sv.cleanOperation(operation)

	// Match operation type
	operationRegex := regexp.MustCompile(`^\s*(query|mutation|subscription)\s*`)
	matches := operationRegex.FindStringSubmatch(cleaned)

	if len(matches) > 1 {
		return strings.ToLower(matches[1]), nil
	}

	// Default to query if no operation type specified
	return "query", nil
}

// validateOperationType validates that the operation type exists in the schema
func (sv *SchemaValidator) validateOperationType(operationType string, result *ValidationResult) error {
	switch operationType {
	case "query":
		if sv.schema.QueryType == nil {
			result.AddError(ErrorTypeTypeNotFound, "Schema does not support query operations", "")
			return fmt.Errorf("query type not available")
		}
	case "mutation":
		if sv.schema.MutationType == nil {
			result.AddError(ErrorTypeTypeNotFound, "Schema does not support mutation operations", "")
			return fmt.Errorf("mutation type not available")
		}
	case "subscription":
		if sv.schema.SubscriptionType == nil {
			result.AddError(ErrorTypeTypeNotFound, "Schema does not support subscription operations", "")
			return fmt.Errorf("subscription type not available")
		}
	default:
		result.AddError(ErrorTypeInvalidOperation, fmt.Sprintf("Unknown operation type: %s", operationType), "")
		return fmt.Errorf("invalid operation type")
	}
	return nil
}

// extractFields extracts root-level field names from the GraphQL operation
func (sv *SchemaValidator) extractFields(operation string) ([]string, error) {
	cleaned := sv.cleanOperation(operation)

	// Find the selection set (content between first { and matching })
	openBrace := strings.Index(cleaned, "{")
	if openBrace == -1 {
		return nil, fmt.Errorf("invalid GraphQL operation: missing selection set")
	}

	// Find matching closing brace
	braceCount := 0
	closeBrace := -1
	for i := openBrace; i < len(cleaned); i++ {
		if cleaned[i] == '{' {
			braceCount++
		} else if cleaned[i] == '}' {
			braceCount--
			if braceCount == 0 {
				closeBrace = i
				break
			}
		}
	}

	if closeBrace == -1 {
		return nil, fmt.Errorf("invalid GraphQL operation: unmatched braces")
	}

	selectionSet := cleaned[openBrace+1 : closeBrace]
	return sv.parseRootFields(selectionSet), nil
}

// parseRootFields parses only root-level field names from a selection set
func (sv *SchemaValidator) parseRootFields(selectionSet string) []string {
	var fields []string

	// Parse root-level fields by looking for field names that are not inside nested braces
	i := 0
	for i < len(selectionSet) {
		// Skip whitespace
		for i < len(selectionSet) && (selectionSet[i] == ' ' || selectionSet[i] == '\t' || selectionSet[i] == '\n') {
			i++
		}

		if i >= len(selectionSet) {
			break
		}

		// Extract field name
		start := i
		for i < len(selectionSet) && (isAlphaNumeric(selectionSet[i]) || selectionSet[i] == '_') {
			i++
		}

		if i > start {
			fieldName := selectionSet[start:i]
			if !sv.isGraphQLKeyword(fieldName) {
				fields = append(fields, fieldName)
			}
		}

		// Skip arguments if present
		for i < len(selectionSet) && selectionSet[i] == ' ' {
			i++
		}
		if i < len(selectionSet) && selectionSet[i] == '(' {
			parenCount := 1
			i++
			for i < len(selectionSet) && parenCount > 0 {
				switch selectionSet[i] {
				case '(':
					parenCount++
				case ')':
					parenCount--
				}
				i++
			}
		}

		// Skip nested selection set if present
		for i < len(selectionSet) && selectionSet[i] == ' ' {
			i++
		}
		if i < len(selectionSet) && selectionSet[i] == '{' {
			braceCount := 1
			i++
			for i < len(selectionSet) && braceCount > 0 {
				switch selectionSet[i] {
				case '{':
					braceCount++
				case '}':
					braceCount--
				}
				i++
			}
		}

		// Skip to next field
		for i < len(selectionSet) && (selectionSet[i] == ' ' || selectionSet[i] == ',' || selectionSet[i] == '\n') {
			i++
		}
	}

	return sv.removeDuplicates(fields)
}

// isAlphaNumeric checks if a character is alphanumeric
func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

// validateFields validates that all fields exist in the appropriate schema type
func (sv *SchemaValidator) validateFields(fields []string, operationType string, result *ValidationResult) {
	var rootType *Type

	switch operationType {
	case "query":
		rootType = sv.schema.QueryType
	case "mutation":
		rootType = sv.schema.MutationType
	case "subscription":
		rootType = sv.schema.SubscriptionType
	}

	if rootType == nil {
		return
	}

	// Find the actual type definition
	actualType := sv.schema.GetTypeByName(rootType.Name)
	if actualType == nil {
		result.AddError(ErrorTypeTypeNotFound, fmt.Sprintf("Root type '%s' not found in schema", rootType.Name), "")
		return
	}

	// Validate each field (these should now be root-level fields only)
	for _, fieldName := range fields {
		if !sv.fieldExistsInType(fieldName, actualType) {
			result.AddDetailedError(
				ErrorTypeFieldNotFound,
				fmt.Sprintf("Field '%s' does not exist on type '%s'", fieldName, actualType.Name),
				fieldName,
				map[string]interface{}{
					"type":             actualType.Name,
					"available_fields": sv.getFieldNames(actualType),
				},
			)
		}
	}
}

// fieldExistsInType checks if a field exists in the given type
func (sv *SchemaValidator) fieldExistsInType(fieldName string, typ *Type) bool {
	for _, field := range typ.Fields {
		if field.Name == fieldName {
			return true
		}
	}
	return false
}

// getFieldNames returns all field names for a type
func (sv *SchemaValidator) getFieldNames(typ *Type) []string {
	var names []string
	for _, field := range typ.Fields {
		names = append(names, field.Name)
	}
	return names
}

// cleanOperation removes comments and normalizes whitespace
func (sv *SchemaValidator) cleanOperation(operation string) string {
	// Remove single-line comments
	lines := strings.Split(operation, "\n")
	var cleanedLines []string

	for _, line := range lines {
		// Remove comments (simple approach - doesn't handle comments in strings)
		if commentIndex := strings.Index(line, "#"); commentIndex != -1 {
			line = line[:commentIndex]
		}
		cleanedLines = append(cleanedLines, line)
	}

	cleaned := strings.Join(cleanedLines, " ")

	// Normalize whitespace
	spaceRegex := regexp.MustCompile(`\s+`)
	return spaceRegex.ReplaceAllString(strings.TrimSpace(cleaned), " ")
}

// isGraphQLKeyword checks if a string is a GraphQL keyword
func (sv *SchemaValidator) isGraphQLKeyword(word string) bool {
	keywords := map[string]bool{
		"query":        true,
		"mutation":     true,
		"subscription": true,
		"fragment":     true,
		"on":           true,
		"true":         true,
		"false":        true,
		"null":         true,
	}
	return keywords[strings.ToLower(word)]
}

// removeDuplicates removes duplicate strings from a slice
func (sv *SchemaValidator) removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// ValidateFieldPath validates a specific field path against the schema
func (sv *SchemaValidator) ValidateFieldPath(path string, operationType string) *ValidationResult {
	result := &ValidationResult{Valid: true}

	if sv.schema == nil {
		result.AddError(ErrorTypeSchemaUnavailable, "Schema is not available for validation", "")
		return result
	}

	// Get root type
	var rootType *Type
	switch operationType {
	case "query":
		rootType = sv.schema.QueryType
	case "mutation":
		rootType = sv.schema.MutationType
	case "subscription":
		rootType = sv.schema.SubscriptionType
	default:
		result.AddError(ErrorTypeInvalidOperation, fmt.Sprintf("Unknown operation type: %s", operationType), "")
		return result
	}

	if rootType == nil {
		result.AddError(ErrorTypeTypeNotFound, fmt.Sprintf("Schema does not support %s operations", operationType), "")
		return result
	}

	// Validate the field path
	sv.validateFieldPathRecursive(path, rootType.Name, result)

	return result
}

// validateFieldPathRecursive validates a field path recursively
func (sv *SchemaValidator) validateFieldPathRecursive(path string, typeName string, result *ValidationResult) {
	if path == "" {
		return
	}

	// Split path into first field and remaining path
	parts := strings.SplitN(path, ".", 2)
	fieldName := parts[0]
	remainingPath := ""
	if len(parts) > 1 {
		remainingPath = parts[1]
	}

	// Find the type
	typ := sv.schema.GetTypeByName(typeName)
	if typ == nil {
		result.AddError(ErrorTypeTypeNotFound, fmt.Sprintf("Type '%s' not found in schema", typeName), fieldName)
		return
	}

	// Find the field
	field := typ.GetFieldByName(fieldName)
	if field == nil {
		result.AddDetailedError(
			ErrorTypeFieldNotFound,
			fmt.Sprintf("Field '%s' does not exist on type '%s'", fieldName, typeName),
			fieldName,
			map[string]interface{}{
				"type":             typeName,
				"available_fields": sv.getFieldNames(typ),
			},
		)
		return
	}

	// If there's more path to validate, continue with the field's type
	if remainingPath != "" {
		fieldTypeName := sv.getBaseTypeName(&field.Type)
		if fieldTypeName != "" {
			sv.validateFieldPathRecursive(remainingPath, fieldTypeName, result)
		}
	}
}

// getBaseTypeName extracts the base type name from a type (handling NON_NULL and LIST wrappers)
func (sv *SchemaValidator) getBaseTypeName(typ *Type) string {
	if typ == nil {
		return ""
	}

	// Handle wrapped types (NON_NULL, LIST)
	if typ.Kind == "NON_NULL" || typ.Kind == "LIST" {
		if typ.OfType != nil {
			return sv.getBaseTypeName(typ.OfType)
		}
	}

	return typ.Name
}
