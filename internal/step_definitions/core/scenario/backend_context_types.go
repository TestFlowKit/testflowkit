package scenario

import (
	"context"
	"testflowkit/pkg/graphql"
	"testflowkit/pkg/variables"
)

// APIProtocol defines the interface for different API protocol implementations
// This is defined here to avoid import cycles with the protocol package.
type APIProtocol interface {
	PrepareRequest(ctx context.Context, name string) (context.Context, error)

	// SendRequest executes the prepared request and stores the response
	SendRequest(ctx context.Context) (context.Context, error)

	// GetResponseBody returns the raw response body as bytes
	GetResponseBody(ctx context.Context) ([]byte, error)

	GetStatusCode(ctx context.Context) (int, error)

	HasErrors(ctx context.Context) bool

	GetProtocolName() string
}

// BackendContext is the unified context for both GraphQL and REST API testing.
type BackendContext struct {
	// Shared fields
	Headers  map[string]string
	Response *UnifiedResponse
	Protocol APIProtocol
	parser   *variables.Parser

	Rest    RestContext
	GraphQL GraphQLContext
}

type RestContext struct {
	Endpoint    *EndpointEnricher
	RequestBody []byte
}

type GraphQLContext struct {
	Request   *graphql.Request
	Variables map[string]any
}

type UnifiedResponse struct {
	StatusCode    int
	Body          []byte
	Headers       map[string]string
	GraphQLErrors []graphql.Error
}
