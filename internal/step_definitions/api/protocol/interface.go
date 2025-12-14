package protocol

import (
	"context"
)

// APIProtocol defines the interface for different API protocol implementations
// This allows GraphQL and REST API to share common step definitions while
// maintaining protocol-specific behavior.
type APIProtocol interface {
	// PrepareRequest prepares a request by looking up the operation/endpoint in config
	PrepareRequest(ctx context.Context, name string) (context.Context, error)

	// SendRequest executes the prepared request and stores the response
	SendRequest(ctx context.Context) (context.Context, error)

	// GetResponseBody returns the raw response body as bytes
	GetResponseBody(ctx context.Context) ([]byte, error)

	// GetStatusCode returns the HTTP status code of the last response
	GetStatusCode(ctx context.Context) (int, error)

	// HasErrors returns true if the response contains protocol-level errors
	// For GraphQL: checks the errors array
	// For REST: checks if status code indicates error
	HasErrors(ctx context.Context) bool

	// GetProtocolName returns the name of the protocol (e.g., "GraphQL", "REST API")
	GetProtocolName() string
}
