package graphql

import (
	"encoding/json"
)

// Request represents a GraphQL request.
type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// Response represents a GraphQL response.
type Response struct {
	Data       json.RawMessage        `json:"data,omitempty"`
	Errors     []Error                `json:"errors,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	StatusCode int                    `json:"-"` // HTTP status code (not part of GraphQL spec)
}

// Error represents a GraphQL error.
type Error struct {
	Message    string                 `json:"message"`
	Locations  []ErrorLocation        `json:"locations,omitempty"`
	Path       []interface{}          `json:"path,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// ErrorLocation represents the location of a GraphQL error.
type ErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// ErrorSummary provides a structured overview of errors in a GraphQL response.
type ErrorSummary struct {
	TotalErrors      int            `json:"total_errors"`
	ErrorsByType     map[string]int `json:"errors_by_type"`
	ErrorsBySeverity map[string]int `json:"errors_by_severity"`
	Messages         []string       `json:"messages"`
}

// GetParser returns a ResponseParser for this response.
func (r *Response) GetParser() *ResponseParser {
	return NewResponseParser(r)
}
