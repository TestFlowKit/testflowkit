// Package providers contains the auth provider interface and all concrete
// implementations for TestFlowKit's security framework.
package providers

import (
	"context"
	"errors"
	"testflowkit/internal/config"
)

// ErrNotImplemented is returned by stub providers that are reserved for a
// future implementation phase (e.g. OIDC, certificate).
var ErrNotImplemented = errors.New("auth provider: not implemented in this release")

// TokenResult is the auth material produced by a provider.
// The transport layer uses it to attach credentials to outgoing requests.
type TokenResult struct {
	// AccessToken is the raw token/key value.
	AccessToken string
	// TokenType is e.g. "Bearer" (used as the Authorization header prefix).
	TokenType string
	// ExpiresIn is the TTL reported by the IDP.  Zero means unknown/no expiry.
	ExpiresIn int64
	// HeaderName is the custom header name for API-key schemes;
	// empty defaults to "Authorization".
	HeaderName string
	// HeaderValue is the fully-formatted header value, e.g. "Bearer <token>"
	// or the raw key.  When set, AccessToken and TokenType are ignored for
	// header injection.
	HeaderValue string
	// QueryParam is set by API-key providers with placement: query.
	QueryParam string
	// Placement tells the transport where to attach the credential.
	Placement config.APIKeyPlacement
}

// Provider is the interface every auth backend must implement.
type Provider interface {
	// Authenticate acquires (or refreshes) credentials for the given scheme
	// and returns a TokenResult.  Implementations must respect ctx cancellation.
	Authenticate(ctx context.Context, scheme config.SecurityScheme) (*TokenResult, error)
}
