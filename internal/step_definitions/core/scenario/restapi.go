package scenario

import (
	"encoding/json"
	"fmt"
	"maps"
	"testflowkit/internal/config"
)

func (c *Context) GetResponse() *HTTPResponse {
	resp := c.backend.GetResponse()
	if resp == nil {
		return nil
	}
	return &HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       resp.Body,
	}
}

func (c *Context) SetEndpoint(
	configuredBaseURL string,
	endpoint config.Endpoint,
) {
	c.backend.SetEndpoint(&EndpointEnricher{
		ConfiguredBaseURL: configuredBaseURL,
		QueryParams:       make(map[string]string),
		PathParams:        make(map[string]string),
		Endpoint:          endpoint,
	})
}

func (c *Context) GetEndpoint() EndpointEnricher {
	return *c.backend.GetEndpoint()
}

func (c *Context) AddPathParam(param string, s string) {
	param = ReplaceVariablesInString(c, param)
	s = ReplaceVariablesInString(c, s)
	c.backend.AddRESTPathParam(param, s)
}

func (c *Context) AddQueryParam(key, value string) {
	key = ReplaceVariablesInString(c, key)
	value = ReplaceVariablesInString(c, value)
	c.backend.AddRESTQueryParam(key, value)
}

func (c *Context) SetRequestBody(body []byte) error {
	c.backend.SetRESTRequestBody(body)
	return nil
}

func (c *Context) SetRequestBodyAsJSON(body []byte) error {
	var jsonTest interface{}
	if err := json.Unmarshal(body, &jsonTest); err != nil {
		return fmt.Errorf("invalid JSON in request body: %w", err)
	}
	c.backend.SetRESTRequestBody(body)
	return nil
}

func (c *Context) AddHeader(key, value string) {
	key = ReplaceVariablesInString(c, key)
	value = ReplaceVariablesInString(c, value)
	c.backend.SetHeader(key, value)
}

func (c *Context) GetRESTRequestBody() []byte {
	return c.backend.GetRESTRequestBody()
}

func (c *Context) GetRequestHeaders() map[string]string {
	headers := c.config.Backend.DefaultHeaders
	maps.Copy(headers, c.backend.GetHeaders())
	return headers
}

func (c *Context) SetResponse(statusCode int, body []byte) {
	c.backend.SetResponse(&UnifiedResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    make(map[string]string),
	})
}
