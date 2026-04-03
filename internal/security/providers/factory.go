package providers

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
)

// Factory returns the Provider for the given SecuritySchemeType.
// Returns an error for unknown or not-implemented types.
func Factory(t config.SecuritySchemeType) (Provider, error) {
	switch t {
	case config.SecurityTypeBearer:
		return &BearerProvider{}, nil
	case config.SecurityTypeBasic:
		return &BasicProvider{}, nil
	case config.SecurityTypeAPIKey:
		return &APIKeyProvider{}, nil
	case config.SecurityTypeOAuth2:
		return &OAuth2Provider{}, nil
	case config.SecurityTypeOIDC:
		return &OIDCProvider{}, nil
	case config.SecurityTypeCertificate:
		return &CertificateProvider{}, nil
	case config.SecurityTypeNone, "":
		// Caller should have checked for the none sentinel before reaching here.
		return nil, fmt.Errorf("factory: security type '%s' does not require a provider", t)
	default:
		return nil, fmt.Errorf("factory: unknown security type '%s'", t)
	}
}

// Authenticate is a convenience wrapper that resolves the provider for the
// scheme type and calls Authenticate in a single step.
func Authenticate(ctx context.Context, scheme config.SecurityScheme) (*TokenResult, error) {
	p, err := Factory(scheme.Type)
	if err != nil {
		return nil, err
	}
	return p.Authenticate(ctx, scheme)
}
