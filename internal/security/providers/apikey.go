package providers

import (
	"context"
	"testflowkit/internal/config"
)

// APIKeyProvider handles type: apikey – a static API key delivered via header,
// query parameter, or cookie.
type APIKeyProvider struct{}

func (*APIKeyProvider) Authenticate(_ context.Context, scheme config.SecurityScheme) (*TokenResult, error) {
	placement := scheme.Placement
	if placement == "" {
		placement = config.APIKeyPlacementHeader
	}

	result := &TokenResult{
		AccessToken: scheme.Key,
		Placement:   placement,
	}

	switch placement {
	case config.APIKeyPlacementHeader:
		headerName := scheme.HeaderName
		if headerName == "" {
			headerName = "Authorization"
		}
		result.HeaderName = headerName
		result.HeaderValue = scheme.Key
	case config.APIKeyPlacementQuery:
		param := scheme.QueryParam
		if param == "" {
			param = DefaultAPIKeyQueryParam
		}
		result.QueryParam = param
	case config.APIKeyPlacementCookie:
		// Cookie placement is handled by the transport layer using HeaderName
		// as the cookie name.
		cookieName := scheme.HeaderName
		if cookieName == "" {
			cookieName = DefaultAPIKeyQueryParam
		}
		result.HeaderName = cookieName
	}

	return result, nil
}
