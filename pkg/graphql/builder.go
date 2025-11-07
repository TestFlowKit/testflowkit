package graphql

import "errors"

// RequestBuilder helps build GraphQL requests with a fluent API.
type RequestBuilder struct {
	query     string
	variables map[string]interface{}
}

// NewRequestBuilder creates a new request builder.
func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		variables: make(map[string]interface{}),
	}
}

// WithQuery sets the GraphQL query/mutation.
func (rb *RequestBuilder) WithQuery(query string) *RequestBuilder {
	rb.query = query
	return rb
}

// WithVariable sets a single variable.
func (rb *RequestBuilder) WithVariable(name string, value interface{}) *RequestBuilder {
	rb.variables[name] = value
	return rb
}

// WithVariables sets multiple variables at once.
func (rb *RequestBuilder) WithVariables(variables map[string]interface{}) *RequestBuilder {
	for name, value := range variables {
		rb.variables[name] = value
	}
	return rb
}

// ClearVariables removes all variables.
func (rb *RequestBuilder) ClearVariables() *RequestBuilder {
	rb.variables = make(map[string]interface{})
	return rb
}

// GetVariables returns a copy of the current variables.
func (rb *RequestBuilder) GetVariables() map[string]interface{} {
	variables := make(map[string]interface{})
	for name, value := range rb.variables {
		variables[name] = value
	}
	return variables
}

// GetQuery returns the current query.
func (rb *RequestBuilder) GetQuery() string {
	return rb.query
}

// Build creates the final GraphQL request.
func (rb *RequestBuilder) Build() (Request, error) {
	if rb.query == "" {
		return Request{}, errors.New("GraphQL query is required")
	}

	return Request{
		Query:     rb.query,
		Variables: rb.variables,
	}, nil
}

// Clone creates a copy of the request builder.
func (rb *RequestBuilder) Clone() *RequestBuilder {
	clone := NewRequestBuilder()
	clone.query = rb.query
	for name, value := range rb.variables {
		clone.variables[name] = value
	}
	return clone
}
