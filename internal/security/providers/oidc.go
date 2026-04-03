package providers

import (
	"context"
	"testflowkit/internal/config"
)

// OIDCProvider is a placeholder for a future OIDC / PKCE implementation.
// It satisfies the Provider interface but always returns ErrNotImplemented.
type OIDCProvider struct{}

func (*OIDCProvider) Authenticate(_ context.Context, _ config.SecurityScheme) (*TokenResult, error) {
	return nil, ErrNotImplemented
}
