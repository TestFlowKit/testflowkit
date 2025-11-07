package scenario

import "testflowkit/pkg/graphql"

// GetGraphQLContext returns the GraphQL context
func (c *Context) GetGraphQLContext() *GraphQLContext {
	return c.graphql
}

// SetGraphQLClient sets the GraphQL client
func (c *Context) SetGraphQLClient(client *graphql.Client) {
	c.graphql.SetClient(client)
}

// GetGraphQLClient returns the current GraphQL client
func (c *Context) GetGraphQLClient() *graphql.Client {
	return c.graphql.GetClient()
}

// SetGraphQLRequest sets the current GraphQL request
func (c *Context) SetGraphQLRequest(request *graphql.Request) {
	c.graphql.SetCurrentRequest(request)
}

// GetGraphQLRequest returns the current GraphQL request
func (c *Context) GetGraphQLRequest() *graphql.Request {
	return c.graphql.GetCurrentRequest()
}

// SetGraphQLResponse sets the last GraphQL response
func (c *Context) SetGraphQLResponse(response *graphql.Response) {
	c.graphql.SetLastResponse(response)
}

// GetGraphQLResponse returns the last GraphQL response
func (c *Context) GetGraphQLResponse() *graphql.Response {
	return c.graphql.GetLastResponse()
}

// GetGraphQLVariable gets a GraphQL variable
func (c *Context) GetGraphQLVariable(name string) (interface{}, bool) {
	return c.graphql.GetVariable(name)
}

// SetGraphQLVariable sets a single GraphQL variable
func (c *Context) SetGraphQLVariable(name string, value interface{}) {
	c.graphql.SetVariable(name, value)
}

// SetGraphQLVariableFromString sets a GraphQL variable by parsing a string value
func (c *Context) SetGraphQLVariableFromString(name, value string) error {
	return c.graphql.SetVariableFromString(name, value)
}

// SetGraphQLVariables sets multiple GraphQL variables
func (c *Context) SetGraphQLVariables(variables map[string]interface{}) {
	c.graphql.SetVariables(variables)
}

// SetGraphQLVariablesFromStrings sets multiple GraphQL variables by parsing string values
func (c *Context) SetGraphQLVariablesFromStrings(variables map[string]string) error {
	return c.graphql.SetVariablesFromStrings(variables)
}

// SetGraphQLArrayVariable sets a GraphQL variable to an array value from a JSON string
func (c *Context) SetGraphQLArrayVariable(name, arrayValue string) error {
	return c.graphql.SetArrayVariable(name, arrayValue)
}

// SetGraphQLObjectVariable sets a GraphQL variable to an object value from a JSON string
func (c *Context) SetGraphQLObjectVariable(name, objectValue string) error {
	return c.graphql.SetObjectVariable(name, objectValue)
}

// GetGraphQLVariables returns all GraphQL variables
func (c *Context) GetGraphQLVariables() map[string]interface{} {
	return c.graphql.GetVariables()
}

// GetGraphQLVariableType returns the type of a GraphQL variable
func (c *Context) GetGraphQLVariableType(name string) string {
	return c.graphql.GetVariableType(name)
}

// GetGraphQLVariableAsString returns a string representation of a GraphQL variable
func (c *Context) GetGraphQLVariableAsString(name string) (string, error) {
	return c.graphql.GetVariableAsString(name)
}

// HasGraphQLVariable checks if a GraphQL variable exists
func (c *Context) HasGraphQLVariable(name string) bool {
	return c.graphql.HasVariable(name)
}

// RemoveGraphQLVariable removes a GraphQL variable
func (c *Context) RemoveGraphQLVariable(name string) {
	c.graphql.RemoveVariable(name)
}

// GetGraphQLVariableNames returns a list of all GraphQL variable names
func (c *Context) GetGraphQLVariableNames() []string {
	return c.graphql.GetVariableNames()
}

// GetGraphQLVariableCount returns the number of GraphQL variables
func (c *Context) GetGraphQLVariableCount() int {
	return c.graphql.GetVariableCount()
}

// GetGraphQLHeader gets a GraphQL request header
func (c *Context) GetGraphQLHeader(name string) (string, bool) {
	return c.graphql.GetHeader(name)
}

// SetGraphQLHeaders sets multiple GraphQL request headers
func (c *Context) SetGraphQLHeaders(headers map[string]string) {
	c.graphql.SetHeaders(headers)
}

// GetGraphQLHeaders returns all GraphQL request headers
func (c *Context) GetGraphQLHeaders() map[string]string {
	return c.graphql.GetHeaders()
}
