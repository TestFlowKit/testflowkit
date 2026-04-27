package scenario

import (
	"context"
	"testflowkit/internal/security"
	"testflowkit/pkg/graphql"
	"testflowkit/pkg/variables"
	"time"
)

// APIProtocol defines the interface for different API protocol implementations
// This is defined here to avoid import cycles with the protocol package.
type APIProtocol interface {
	PrepareRequest(ctx context.Context, apiName string, requestName string) (context.Context, error)

	// SendRequest executes the prepared request and stores the response
	SendRequest(ctx context.Context) (context.Context, error)

	GetCURLCommand(ctx context.Context) (string, error)

	// GetResponseBody returns the raw response body as bytes
	GetResponseBody(ctx context.Context) ([]byte, error)

	GetStatusCode(ctx context.Context) (int, error)

	HasErrors(ctx context.Context) bool

	GetProtocolName() string
}

type PrepareRequestParams struct {
	Ctx         context.Context
	APIName     string
	RequestName string
}

// BackendContext is the unified context for both GraphQL and REST API testing.
type BackendContext struct {
	// Shared fields
	Headers  map[string]string
	Response *UnifiedResponse
	Protocol APIProtocol
	parser   *variables.Parser

	// ResolvedSecurity holds the auth context computed by the security resolver
	// during request preparation. The transport layer uses it to inject
	// credentials and handle retry_on_401.
	ResolvedSecurity security.ResolvedSecurity

	Timeout time.Duration

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
	Endpoint  string
	Headers   map[string]string
}

type UnifiedResponse struct {
	StatusCode    int
	Body          []byte
	Headers       map[string]string
	GraphQLErrors []graphql.Error
}
