// Package graphql provides a comprehensive GraphQL client library for Go applications.
// It supports query execution, response parsing, error handling, and request building.
package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client represents a GraphQL HTTP client with configurable options.
type Client struct {
	httpClient *http.Client
	endpoint   string
	headers    map[string]string
}

// ClientOption represents a configuration option for the GraphQL client.
type ClientOption func(*Client)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets the request timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		if c.httpClient == nil {
			c.httpClient = &http.Client{}
		}
		c.httpClient.Timeout = timeout
	}
}

// WithHeaders sets default headers for all requests.
func WithHeaders(headers map[string]string) ClientOption {
	return func(c *Client) {
		if c.headers == nil {
			c.headers = make(map[string]string)
		}
		for key, value := range headers {
			c.headers[key] = value
		}
	}
}

// NewClient creates a new GraphQL client with the specified endpoint and options.
func NewClient(endpoint string, options ...ClientOption) *Client {
	const clientTimeout = 30
	client := &Client{
		endpoint: endpoint,
		headers:  make(map[string]string),
		httpClient: &http.Client{
			Timeout: clientTimeout * time.Second,
		},
	}

	// Apply options
	for _, option := range options {
		option(client)
	}

	return client
}

// SetHeaders sets multiple headers for all requests.
func (c *Client) SetHeaders(headers map[string]string) {
	for key, value := range headers {
		c.headers[key] = value
	}
}

// GetEndpoint returns the GraphQL endpoint URL.
func (c *Client) GetEndpoint() string {
	return c.endpoint
}

// GetHeaders returns a copy of the default headers.
func (c *Client) GetHeaders() map[string]string {
	headers := make(map[string]string)
	for key, value := range c.headers {
		headers[key] = value
	}
	return headers
}

// Execute sends a GraphQL request and returns the response.
func (c *Client) Execute(ctx context.Context, req Request) (*Response, error) {
	requestBody, err := c.getReqBody(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, requestBody)
	if err != nil {
		return nil, NewNetworkError(
			"failed to create HTTP request",
			map[string]interface{}{
				"endpoint": c.endpoint,
				"error":    err.Error(),
			},
		)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	for key, value := range c.headers {
		httpReq.Header.Set(key, value)
	}

	// Execute the HTTP request
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, NewNetworkError(
			"failed to execute HTTP request",
			map[string]interface{}{
				"endpoint": c.endpoint,
				"error":    err.Error(),
			},
		)
	}
	defer httpResp.Body.Close()

	// Read response body
	responseBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, NewNetworkError(
			"failed to read response body",
			map[string]interface{}{
				"status_code": httpResp.StatusCode,
				"error":       err.Error(),
			},
		)
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, NewNetworkError(
			fmt.Sprintf("HTTP request failed with status %d", httpResp.StatusCode),
			map[string]interface{}{
				"status_code":   httpResp.StatusCode,
				"response_body": string(responseBody),
			},
		)
	}

	// Parse GraphQL response
	graphqlResp, err := parseResponse(responseBody, httpResp.StatusCode)
	if err != nil {
		return nil, err
	}

	graphqlResp.Headers = make(map[string]string)
	for key, values := range httpResp.Header {
		graphqlResp.Headers[key] = strings.Join(values, ", ")
	}

	return graphqlResp, nil
}

func (*Client) getReqBody(req Request) (*bytes.Buffer, error) {
	if req.Query == "" {
		return nil, NewConfigurationError("GraphQL query is required", nil)
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, NewConfigurationError(
			"failed to marshal GraphQL request",
			map[string]interface{}{"marshal_error": err.Error()},
		)
	}
	return bytes.NewBuffer(requestBody), nil
}

// ExecuteWithBuilder executes a request using a RequestBuilder.
func (c *Client) ExecuteWithBuilder(ctx context.Context, builder *RequestBuilder) (*Response, error) {
	req, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build GraphQL request: %w", err)
	}

	return c.Execute(ctx, req)
}

// Query is a convenience method for executing GraphQL queries.
func (c *Client) Query(ctx context.Context, query string, variables map[string]interface{}) (*Response, error) {
	req := Request{
		Query:     query,
		Variables: variables,
	}
	return c.Execute(ctx, req)
}

// Mutate is a convenience method for executing GraphQL mutations.
func (c *Client) Mutate(ctx context.Context, mutation string, variables map[string]interface{}) (*Response, error) {
	req := Request{
		Query:     mutation,
		Variables: variables,
	}
	return c.Execute(ctx, req)
}
