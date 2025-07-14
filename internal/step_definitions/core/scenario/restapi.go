package scenario

import (
	"encoding/json"
	"fmt"
	"testflowkit/internal/config"
)

func (c *Context) GetResponse() *HTTPResponse {
	return c.http.Response
}

func (c *Context) SetEndpoint(
	configuredBaseURL string,
	endpoint config.Endpoint,
) {
	c.http.Endpoint = &EndpointEnricher{
		ConfiguredBaseURL: configuredBaseURL,
		QueryParams:       make(map[string]string),
		PathParams:        make(map[string]string),
		Endpoint:          endpoint,
	}
}

func (c *Context) GetEndpoint() EndpointEnricher {
	return *c.http.Endpoint
}

func (c *Context) AddPathParam(param string, s string) {
	param = ReplaceVariablesInString(c, param)
	s = ReplaceVariablesInString(c, s)
	c.http.Endpoint.PathParams[param] = s
}

func (c *Context) AddQueryParam(key, value string) {
	key = ReplaceVariablesInString(c, key)
	value = ReplaceVariablesInString(c, value)
	c.http.Endpoint.QueryParams[key] = value
}

func (c *Context) SetRequestBody(body []byte) error {
	var jsonTest interface{}
	if err := json.Unmarshal(body, &jsonTest); err != nil {
		return fmt.Errorf("invalid JSON in request body: %w", err)
	}
	c.http.RequestBody = body
	return nil
}

func (c *Context) AddHeader(key, value string) {
	key = ReplaceVariablesInString(c, key)
	value = ReplaceVariablesInString(c, value)
	c.http.RequestHeaders[key] = value
}

func (c *Context) GetRequestBody() []byte {
	return c.http.RequestBody
}

func (c *Context) GetRequestHeaders() map[string]string {
	return c.http.RequestHeaders
}

func (c *Context) SetResponse(statusCode int, body []byte) {
	c.http.Response = &HTTPResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}
