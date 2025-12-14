package scenario

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/pkg/variables"
)

// APIProtocol defines the interface for different API protocol implementations
// This is defined here to avoid import cycles with the protocol package.
type APIProtocol interface {
	// PrepareRequest prepares a request by looking up the operation/endpoint in config
	PrepareRequest(ctx context.Context, name string) (context.Context, error)

	// SendRequest executes the prepared request and stores the response
	SendRequest(ctx context.Context) (context.Context, error)

	// GetResponseBody returns the raw response body as bytes
	GetResponseBody(ctx context.Context) ([]byte, error)

	// GetStatusCode returns the HTTP status code of the last response
	GetStatusCode(ctx context.Context) (int, error)

	// HasErrors returns true if the response contains protocol-level errors
	HasErrors(ctx context.Context) bool

	GetProtocolName() string
}

type BackendContext struct {
	// Shared fields
	Headers   map[string]string
	Variables map[string]any
	Response  *UnifiedResponse
	Protocol  APIProtocol
	parser    *variables.Parser

	// REST-specific fields (used when Protocol is RESTAPIAdapter)
	Endpoint    *EndpointEnricher
	RequestBody []byte
}

type UnifiedResponse struct {
	StatusCode int
	Body       []byte
	Headers    map[string]string
}

// NewBackendContext creates a new unified backend context.
func NewBackendContext() *BackendContext {
	ctx := &BackendContext{
		Headers:   make(map[string]string),
		Variables: make(map[string]any),
		Response:  nil,
	}
	ctx.parser = variables.NewParser(ctx)
	return ctx
}

func (bc *BackendContext) SetProtocol(p APIProtocol) {
	bc.Protocol = p
}

// GetProtocol returns the current protocol adapter.
func (bc *BackendContext) GetProtocol() APIProtocol {
	return bc.Protocol
}

// IsREST returns true if the current protocol is REST.
func (bc *BackendContext) IsREST() bool {
	return bc.Protocol != nil && bc.Protocol.GetProtocolName() == "REST"
}

// === Header Management ===

func (bc *BackendContext) GetHeader(name string) (string, bool) {
	value, exists := bc.Headers[name]
	return value, exists
}

func (bc *BackendContext) SetHeader(name, value string) {
	bc.Headers[name] = value
}

func (bc *BackendContext) GetHeaders() map[string]string {
	return bc.Headers
}

func (bc *BackendContext) ClearHeaders() {
	bc.Headers = make(map[string]string)
}

// GetVariable gets a variable by name.
func (bc *BackendContext) GetVariable(name string) (any, bool) {
	value, exists := bc.Variables[name]
	return value, exists
}

// SetVariable sets a variable with a pre-parsed value.
func (bc *BackendContext) SetVariable(name string, value any) {
	bc.Variables[name] = value
}

// SetVariableFromString sets a variable by parsing a string value
// The string value will be parsed to determine the appropriate type (array, object, primitive).
func (bc *BackendContext) SetVariableFromString(name, value string) error {
	parsedValue, err := bc.parser.ParseValue(value)
	if err != nil {
		return fmt.Errorf("failed to set variable '%s': %w", name, err)
	}
	bc.Variables[name] = parsedValue
	return nil
}

// SetVariablesFromStrings sets multiple variables by parsing string values.
func (bc *BackendContext) SetVariablesFromStrings(variables map[string]string) error {
	parsedVariables, err := bc.parser.ParseVariables(variables)
	if err != nil {
		return err
	}
	// Merge with existing variables
	for k, v := range parsedVariables {
		bc.Variables[k] = v
	}
	return nil
}

// ClearVariables clears all variables.
func (bc *BackendContext) ClearVariables() {
	bc.Variables = make(map[string]any)
}

// === Response Management ===

func (bc *BackendContext) SetResponse(response *UnifiedResponse) {
	bc.Response = response
}

func (bc *BackendContext) GetResponse() *UnifiedResponse {
	return bc.Response
}

func (bc *BackendContext) HasResponse() bool {
	return bc.Response != nil
}

// GetResponseBody returns the raw response body.
func (bc *BackendContext) GetResponseBody() []byte {
	if bc.Response == nil {
		return nil
	}
	return bc.Response.Body
}

// GetStatusCode returns the HTTP status code.
func (bc *BackendContext) GetStatusCode() int {
	if bc.Response == nil {
		return 0
	}
	return bc.Response.StatusCode
}

// === REST-specific methods ===

// SetEndpoint sets the REST endpoint configuration.
func (bc *BackendContext) SetEndpoint(endpoint *EndpointEnricher) {
	bc.Endpoint = endpoint
}

// GetEndpoint returns the REST endpoint configuration.
func (bc *BackendContext) GetEndpoint() *EndpointEnricher {
	return bc.Endpoint
}

// SetRequestBody sets the REST request body.
func (bc *BackendContext) SetRequestBody(body []byte) {
	bc.RequestBody = body
}

func (bc *BackendContext) AddPathParam(param, value string) {
	if bc.Endpoint == nil {
		bc.Endpoint = &EndpointEnricher{
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		}
	}
	bc.Endpoint.PathParams[param] = value
}

func (bc *BackendContext) AddQueryParam(key, value string) {
	if bc.Endpoint == nil {
		bc.Endpoint = &EndpointEnricher{
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		}
	}
	bc.Endpoint.QueryParams[key] = value
}

// GetRequestBody returns the REST request body.
func (bc *BackendContext) GetRequestBody() []byte {
	return bc.RequestBody
}

// === Reset ===

// Reset clears all context data.
func (bc *BackendContext) Reset() {
	bc.ClearHeaders()
	bc.ClearVariables()
	bc.Response = nil
	bc.Endpoint = nil
	bc.RequestBody = nil
	bc.Protocol = nil
}

// === variables.Store interface implementation (for parser) ===

// Get implements the variables.Store interface for variable substitution.
func (bc *BackendContext) Get(name string) (any, bool) {
	return bc.GetVariable(name)
}

// Set implements the variables.Store interface for variable substitution.
func (bc *BackendContext) Set(name string, value any) {
	bc.SetVariable(name, value)
}

// === EndpointEnricher (REST-specific) ===

type EndpointEnricher struct {
	config.Endpoint
	QueryParams       map[string]string
	PathParams        map[string]string
	ConfiguredBaseURL string
}

func (e *EndpointEnricher) GetFullURL() string {
	fullURL, err := e.getSimpleURL()
	if err != nil {
		return ""
	}

	if len(e.PathParams) > 0 {
		for name, value := range e.PathParams {
			placeholder := fmt.Sprintf("{%s}", name)
			fullURL = strings.ReplaceAll(fullURL, placeholder, value)
		}
	}

	if len(e.QueryParams) > 0 {
		values := url.Values{}
		for key, value := range e.QueryParams {
			values.Set(key, value)
		}
		encoded := values.Encode()
		if encoded != "" {
			fullURL += "?" + encoded
		}
	}

	return fullURL
}

func (e *EndpointEnricher) getSimpleURL() (string, error) {
	// Simple URL parsing
	path := e.Path
	if path == "" {
		return "", errors.New("endpoint path is empty")
	}

	// Check if it's an absolute URL (contains ://)
	if idx := strings.Index(path, "://"); idx > 0 {
		return path, nil
	}

	// Join with base URL
	baseURL := e.ConfiguredBaseURL
	if baseURL == "" {
		return path, nil
	}

	// Remove trailing slash from base URL
	if len(baseURL) > 0 && baseURL[len(baseURL)-1] == '/' {
		baseURL = baseURL[:len(baseURL)-1]
	}

	// Add leading slash to path if missing
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}

	return baseURL + path, nil
}

// SubstituteVariables replaces all variables in the backend context using the provided context.
func (bc *BackendContext) SubstituteVariables(ctx *Context) error {
	// 1. Headers
	newHeaders := make(map[string]string)
	for k, v := range bc.Headers {
		newKey := ReplaceVariablesInString(ctx, k)
		newHeaders[newKey] = ReplaceVariablesInString(ctx, v)
	}
	bc.Headers = newHeaders

	// 2. Endpoint (REST)
	if bc.Endpoint != nil {
		newQueryParams := make(map[string]string)
		for k, v := range bc.Endpoint.QueryParams {
			newKey := ReplaceVariablesInString(ctx, k)
			newQueryParams[newKey] = ReplaceVariablesInString(ctx, v)
		}
		bc.Endpoint.QueryParams = newQueryParams

		newPathParams := make(map[string]string)
		for k, v := range bc.Endpoint.PathParams {
			newKey := ReplaceVariablesInString(ctx, k)
			newPathParams[newKey] = ReplaceVariablesInString(ctx, v)
		}
		bc.Endpoint.PathParams = newPathParams
	}

	// 3. RequestBody (REST)
	if len(bc.RequestBody) > 0 {
		bodyStr := string(bc.RequestBody)
		newBody := ReplaceVariablesInString(ctx, bodyStr)
		bc.RequestBody = []byte(newBody)
	}

	if len(bc.Variables) > 0 {
		// Use JSON roundtrip to handle nested variables and type conversion
		err := bc.marshalAndSubstituteVariables(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bc *BackendContext) marshalAndSubstituteVariables(ctx *Context) error {
	jsonData, err := json.Marshal(bc.Variables)
	if err != nil {
		return fmt.Errorf("failed to marshal variables for substitution: %w", err)
	}

	substitutedJSON := ReplaceVariablesInString(ctx, string(jsonData))

	var newVariables map[string]any
	if errJSONDecode := json.Unmarshal([]byte(substitutedJSON), &newVariables); errJSONDecode != nil {
		return fmt.Errorf("failed to unmarshal variables after substitution: %w", errJSONDecode)
	}
	bc.Variables = newVariables

	// Post-processing: Try to parse strings that look like JSON or booleans/numbers
	// This ensures that variables that were substituted into complex types (arrays/objects)
	// are correctly parsed into their Go types.
	for k, v := range bc.Variables {
		if strVal, ok := v.(string); ok {
			if parsed, errParse := bc.parser.ParseValue(strVal); errParse == nil {
				bc.Variables[k] = parsed
			}
		}
	}
	return nil
}
