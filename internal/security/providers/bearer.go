package providers

import (
	"context"
	"testflowkit/internal/config"
)

// BearerProvider handles type: bearer – a static pre-issued access token.
type BearerProvider struct{}

func (*BearerProvider) Authenticate(_ context.Context, scheme config.SecurityScheme) (*TokenResult, error) {
	return &TokenResult{
		AccessToken: scheme.Token,
		TokenType:   BearerTokenType,
		HeaderValue: "Bearer " + scheme.Token,
	}, nil
}
