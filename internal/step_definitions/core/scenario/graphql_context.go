package scenario

import (
	"fmt"
	"testflowkit/pkg/graphql"
)

// GraphQLContext manages GraphQL-specific state during test execution
type GraphQLContext struct {
	Client         *graphql.Client
	CurrentRequest *graphql.Request
	LastResponse   *graphql.Response
	Variables      map[string]interface{}
	Headers        map[string]string
	parser         *VariableParser
}

// NewGraphQLContext creates a new GraphQL context
func NewGraphQLContext() *GraphQLContext {
	return &GraphQLContext{
		Variables: make(map[string]interface{}),
		Headers:   make(map[string]string),
		parser:    NewVariableParser(),
	}
}

// SetClient sets the GraphQL client for this context
func (gc *GraphQLContext) SetClient(client *graphql.Client) {
	gc.Client = client
}

// GetClient returns the current GraphQL client
func (gc *GraphQLContext) GetClient() *graphql.Client {
	return gc.Client
}

// SetCurrentRequest sets the current GraphQL request
func (gc *GraphQLContext) SetCurrentRequest(request *graphql.Request) {
	gc.CurrentRequest = request
}

// GetCurrentRequest returns the current GraphQL request
func (gc *GraphQLContext) GetCurrentRequest() *graphql.Request {
	return gc.CurrentRequest
}

// SetLastResponse sets the last GraphQL response
func (gc *GraphQLContext) SetLastResponse(response *graphql.Response) {
	gc.LastResponse = response
}

// GetLastResponse returns the last GraphQL response
func (gc *GraphQLContext) GetLastResponse() *graphql.Response {
	return gc.LastResponse
}

// GetVariable gets a GraphQL variable
func (gc *GraphQLContext) GetVariable(name string) (interface{}, bool) {
	value, exists := gc.Variables[name]
	return value, exists
}

// SetVariable sets a single GraphQL variable with the given value
func (gc *GraphQLContext) SetVariable(name string, value interface{}) {
	gc.Variables[name] = value
}

// SetVariableFromString sets a GraphQL variable by parsing a string value
// The string value will be parsed to determine the appropriate type (array, object, primitive)
func (gc *GraphQLContext) SetVariableFromString(name, value string) error {
	parsedValue, err := gc.parser.ParseValue(value)
	if err != nil {
		return fmt.Errorf("failed to set variable '%s': %w", name, err)
	}
	gc.Variables[name] = parsedValue
	return nil
}

// SetVariables sets multiple GraphQL variables
func (gc *GraphQLContext) SetVariables(variables map[string]interface{}) {
	for name, value := range variables {
		gc.Variables[name] = value
	}
}

// SetVariablesFromStrings sets multiple GraphQL variables by parsing string values
func (gc *GraphQLContext) SetVariablesFromStrings(variables map[string]string) error {
	parsedVariables, err := gc.parser.ParseVariables(variables)
	if err != nil {
		return err
	}
	gc.SetVariables(parsedVariables)
	return nil
}

// SetArrayVariable sets a GraphQL variable to an array value from a JSON string
func (gc *GraphQLContext) SetArrayVariable(name, arrayValue string) error {
	if err := gc.parser.ValidateArrayValue(arrayValue); err != nil {
		return fmt.Errorf("invalid array value for variable '%s': %w", name, err)
	}

	return gc.SetVariableFromString(name, arrayValue)
}

// SetObjectVariable sets a GraphQL variable to an object value from a JSON string
func (gc *GraphQLContext) SetObjectVariable(name, objectValue string) error {
	if err := gc.parser.ValidateObjectValue(objectValue); err != nil {
		return fmt.Errorf("invalid object value for variable '%s': %w", name, err)
	}

	return gc.SetVariableFromString(name, objectValue)
}

// GetVariables returns all GraphQL variables
func (gc *GraphQLContext) GetVariables() map[string]interface{} {
	return gc.Variables
}

// GetVariableType returns the type of a GraphQL variable
func (gc *GraphQLContext) GetVariableType(name string) string {
	if value, exists := gc.Variables[name]; exists {
		return gc.parser.GetVariableType(value)
	}
	return "undefined"
}

// GetVariableAsString returns a string representation of a GraphQL variable
func (gc *GraphQLContext) GetVariableAsString(name string) (string, error) {
	value, exists := gc.Variables[name]
	if !exists {
		return "", fmt.Errorf("variable '%s' not found", name)
	}

	return gc.parser.SerializeValue(value)
}

// HasVariable checks if a GraphQL variable exists
func (gc *GraphQLContext) HasVariable(name string) bool {
	_, exists := gc.Variables[name]
	return exists
}

// RemoveVariable removes a GraphQL variable
func (gc *GraphQLContext) RemoveVariable(name string) {
	delete(gc.Variables, name)
}

// ClearVariables clears all GraphQL variables
func (gc *GraphQLContext) ClearVariables() {
	gc.Variables = make(map[string]interface{})
}

// GetVariableNames returns a list of all variable names
func (gc *GraphQLContext) GetVariableNames() []string {
	names := make([]string, 0, len(gc.Variables))
	for name := range gc.Variables {
		names = append(names, name)
	}
	return names
}

// GetVariableCount returns the number of variables
func (gc *GraphQLContext) GetVariableCount() int {
	return len(gc.Variables)
}

// GetHeader gets a GraphQL request header
func (gc *GraphQLContext) GetHeader(name string) (string, bool) {
	value, exists := gc.Headers[name]
	return value, exists
}

// SetHeaders sets multiple GraphQL request headers
func (gc *GraphQLContext) SetHeaders(headers map[string]string) {
	for name, value := range headers {
		gc.Headers[name] = value
	}
}

// GetHeaders returns all GraphQL request headers
func (gc *GraphQLContext) GetHeaders() map[string]string {
	return gc.Headers
}

// ClearHeaders clears all GraphQL request headers
func (gc *GraphQLContext) ClearHeaders() {
	gc.Headers = make(map[string]string)
}

// Reset resets the GraphQL context to initial state
func (gc *GraphQLContext) Reset() {
	gc.CurrentRequest = nil
	gc.LastResponse = nil
	gc.ClearVariables()
	gc.ClearHeaders()
	// Note: We don't reset the client as it should persist across requests
}
