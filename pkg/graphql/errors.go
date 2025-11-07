package graphql

import (
	"fmt"
	"strings"
)

// ErrorType represents different types of GraphQL client errors.
type ErrorType string

const (
	// ErrorTypeConfiguration indicates a client configuration error.
	ErrorTypeConfiguration ErrorType = "configuration"
	// ErrorTypeSyntax indicates a GraphQL syntax error.
	ErrorTypeSyntax ErrorType = "syntax"
	// ErrorTypeNetwork indicates a network-related error.
	ErrorTypeNetwork ErrorType = "network"
	// ErrorTypeGraphQL indicates a GraphQL execution error.
	ErrorTypeGraphQL ErrorType = "graphql"
	// ErrorTypeSchema indicates a schema-related error.
	ErrorTypeSchema ErrorType = "schema"
)

// ClientError represents a GraphQL client error with additional context.
type ClientError struct {
	Type    ErrorType              `json:"type"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}

// NewConfigurationError creates a configuration error.
func NewConfigurationError(message string, details map[string]interface{}) *ClientError {
	return &ClientError{
		Type:    ErrorTypeConfiguration,
		Message: message,
		Details: details,
	}
}

// NewNetworkError creates a network error.
func NewNetworkError(message string, details map[string]interface{}) *ClientError {
	return &ClientError{
		Type:    ErrorTypeNetwork,
		Message: message,
		Details: details,
	}
}

// NewGraphQLError creates a GraphQL error.
func NewGraphQLError(message string, details map[string]interface{}) *ClientError {
	return &ClientError{
		Type:    ErrorTypeGraphQL,
		Message: message,
		Details: details,
	}
}

// NewSyntaxError creates a syntax error.
func NewSyntaxError(message string, details map[string]interface{}) *ClientError {
	return &ClientError{
		Type:    ErrorTypeSyntax,
		Message: message,
		Details: details,
	}
}

// NewSchemaError creates a schema error.
func NewSchemaError(message string, details map[string]interface{}) *ClientError {
	return &ClientError{
		Type:    ErrorTypeSchema,
		Message: message,
		Details: details,
	}
}

// HasErrors checks if the response contains GraphQL errors.
func (r *Response) HasErrors() bool {
	return len(r.Errors) > 0
}

// GetErrorMessages returns all error messages from the response.
func (r *Response) GetErrorMessages() []string {
	messages := make([]string, len(r.Errors))
	for i, err := range r.Errors {
		messages[i] = err.Message
	}
	return messages
}

// GetErrorsAsString returns all errors as a single formatted string.
func (r *Response) GetErrorsAsString() string {
	if !r.HasErrors() {
		return ""
	}

	messages := r.GetErrorMessages()
	return strings.Join(messages, "; ")
}

// HasData checks if the response contains data.
func (r *Response) HasData() bool {
	return len(r.Data) > 0
}

// IsSuccessful checks if the response is considered successful (has data and no errors).
func (r *Response) IsSuccessful() bool {
	return r.HasData() && !r.HasErrors()
}

// GetDetailedErrorInfo returns detailed information about all errors.
func (r *Response) GetDetailedErrorInfo() []map[string]interface{} {
	if !r.HasErrors() {
		return []map[string]interface{}{}
	}

	details := make([]map[string]interface{}, len(r.Errors))
	for i, err := range r.Errors {
		details[i] = map[string]interface{}{
			"message":    err.Message,
			"locations":  err.Locations,
			"path":       err.Path,
			"extensions": err.Extensions,
		}
	}
	return details
}

// GetFirstError returns the first error if any exist.
func (r *Response) GetFirstError() *Error {
	if len(r.Errors) > 0 {
		return &r.Errors[0]
	}
	return nil
}

// GetErrorsWithPath returns errors that have a specific path element.
func (r *Response) GetErrorsWithPath(pathElement interface{}) []Error {
	filteredErrors := make([]Error, 0, len(r.Errors))
	for _, err := range r.Errors {
		for _, pathItem := range err.Path {
			if pathItem == pathElement {
				filteredErrors = append(filteredErrors, err)
				break
			}
		}
	}
	return filteredErrors
}

// GetErrorsAtLocation returns errors at a specific line/column location.
func (r *Response) GetErrorsAtLocation(line, column int) []Error {
	filteredErrors := make([]Error, 0, len(r.Errors))
	for _, err := range r.Errors {
		for _, location := range err.Locations {
			if location.Line == line && location.Column == column {
				filteredErrors = append(filteredErrors, err)
				break
			}
		}
	}
	return filteredErrors
}

// HasErrorsWithClassification checks if response has errors of a specific classification.
func (r *Response) HasErrorsWithClassification(classification string) bool {
	for _, err := range r.Errors {
		if class, exists := err.Extensions["classification"]; exists {
			if classStr, ok := class.(string); ok && classStr == classification {
				return true
			}
		}
	}
	return false
}

// HasErrorsWithSeverity checks if response has errors of a specific severity.
func (r *Response) HasErrorsWithSeverity(severity string) bool {
	for _, err := range r.Errors {
		if sev, exists := err.Extensions["severity"]; exists {
			if sevStr, ok := sev.(string); ok && sevStr == severity {
				return true
			}
		}
	}
	return false
}
