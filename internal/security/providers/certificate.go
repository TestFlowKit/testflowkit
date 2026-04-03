package providers

import (
	"context"
	"testflowkit/internal/config"
)

// CertificateProvider is a placeholder for a future mTLS implementation.
// It satisfies the Provider interface but always returns ErrNotImplemented.
type CertificateProvider struct{}

func (*CertificateProvider) Authenticate(_ context.Context, _ config.SecurityScheme) (*TokenResult, error) {
	return nil, ErrNotImplemented
}
