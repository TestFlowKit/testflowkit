package scenario

import "testflowkit/pkg/graphql"

func (c *Context) SetGraphQLRequest(request *graphql.Request) {
	c.backend.SetGraphQLRequest(request)
}

func (c *Context) GetGraphQLRequest() *graphql.Request {
	return c.backend.GetGraphQLRequest()
}

func (c *Context) SetGraphQLEndpoint(endpoint string) {
	endpoint = ReplaceVariablesInString(c, endpoint)
	c.backend.SetGraphQLEndpoint(endpoint)
}

func (c *Context) GetGraphQLEndpoint() string {
	return c.backend.GetGraphQLEndpoint()
}

func (c *Context) SetGraphQLHeader(key, value string) {
	key = ReplaceVariablesInString(c, key)
	value = ReplaceVariablesInString(c, value)
	c.backend.SetGraphQLHeader(key, value)
}

func (c *Context) SetGraphQLResponse(response *graphql.Response) {
	unifiedResp := &UnifiedResponse{
		StatusCode:    response.StatusCode,
		Body:          response.Data,
		Headers:       make(map[string]string),
		GraphQLErrors: response.Errors,
	}
	c.backend.SetResponse(unifiedResp)
}

func (c *Context) GetGraphQLResponse() *graphql.Response {
	resp := c.backend.GetResponse()
	if resp == nil {
		return nil
	}
	return &graphql.Response{
		Data:       resp.Body,
		Errors:     resp.GraphQLErrors,
		StatusCode: resp.StatusCode,
	}
}

func (c *Context) GetGraphQLVariable(name string) (any, bool) {
	return c.backend.GetGraphQLVariable(name)
}

func (c *Context) SetGraphQLVariablesFromStrings(variables map[string]string) error {
	variables = c.replaceVariablesInMap(variables)
	return c.backend.SetVariablesFromStrings(variables)
}

func (c *Context) GetGraphQLVariables() map[string]any {
	return c.backend.GetGraphQLVariables()
}

func (c *Context) SetGraphQLHeaders(headers map[string]string) {
	headers = c.replaceVariablesInMap(headers)
	for name, value := range headers {
		c.backend.SetHeader(name, value)
	}
}

func (c *Context) GetGraphQLHeaders() map[string]string {
	return c.backend.GetGraphQLHeaders()
}

func (c *Context) replaceVariablesInMap(input map[string]string) map[string]string {
	output := make(map[string]string)
	for key, value := range input {
		newKey := ReplaceVariablesInString(c, key)
		newValue := ReplaceVariablesInString(c, value)
		output[newKey] = newValue
	}
	return output
}
