package providers

import (
	"context"
	"encoding/base64"
	"testflowkit/internal/config"
)

// BasicProvider handles type: basic – HTTP Basic authentication.
type BasicProvider struct{}

func (*BasicProvider) Authenticate(_ context.Context, scheme config.SecurityScheme) (*TokenResult, error) {
	raw := scheme.Username + ":" + scheme.Password
	encoded := base64.StdEncoding.EncodeToString([]byte(raw))
	return &TokenResult{
		AccessToken: encoded,
		TokenType:   "Basic",
		HeaderValue: "Basic " + encoded,
	}, nil
}
