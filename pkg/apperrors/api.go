package apperrors

import (
	"errors"
	"fmt"
)

// ErrNoAPIsConfigured is returned when no APIs are defined in the configuration.
var ErrNoAPIsConfigured = errors.New("no APIs configured")

// ErrNoGraphQLEndpoint is returned when a GraphQL API has no endpoint configured.
var ErrNoGraphQLEndpoint = errors.New("no GraphQL endpoint configured")

// ErrWrongAPIType is a sentinel that can be wrapped to indicate an API type mismatch.
var ErrWrongAPIType = errors.New("wrong API type")

// ErrEndpointPathEmpty is returned when an endpoint path is not set.
var ErrEndpointPathEmpty = errors.New("endpoint path is empty")

// APINotFoundError is returned when a named API does not exist in the configuration.
type APINotFoundError struct {
	Name string
}

func (e *APINotFoundError) Error() string {
	return fmt.Sprintf("API '%s' not found in configuration", e.Name)
}

// EndpointNotFoundError is returned when a named endpoint does not exist within an API.
type EndpointNotFoundError struct {
	API      string
	Endpoint string
}

func (e *EndpointNotFoundError) Error() string {
	return fmt.Sprintf("endpoint '%s' not found in API '%s'", e.Endpoint, e.API)
}

// OperationNotFoundError is returned when a named GraphQL operation does not exist within an API.
type OperationNotFoundError struct {
	API       string
	Operation string
}

func (e *OperationNotFoundError) Error() string {
	return fmt.Sprintf("operation '%s' not found in API '%s'", e.Operation, e.API)
}
