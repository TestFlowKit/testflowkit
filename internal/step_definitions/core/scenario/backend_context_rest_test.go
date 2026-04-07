package scenario

import (
	"testflowkit/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpointEnricher_GetFullURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		enricher EndpointEnricher
		want     string
	}{
		{
			name: "joins base URL with path params and query params",
			enricher: EndpointEnricher{
				Endpoint: config.Endpoint{
					Method:      "GET",
					Path:        "users/{id}/orders",
					Description: "list user orders",
				},
				ConfiguredBaseURL: "https://api.example.com/v1",
				PathParams: map[string]string{
					"id": "john doe",
				},
				QueryParams: map[string]string{
					"status": "open",
					"page":   "2",
				},
			},
			want: "https://api.example.com/v1/users/john%20doe/orders?page=2&status=open",
		},
		{
			name: "preserves absolute URL and merges query params",
			enricher: EndpointEnricher{
				Endpoint: config.Endpoint{
					Method:      "GET",
					Path:        "https://service.example.com/search?q=golang",
					Description: "search",
				},
				QueryParams: map[string]string{
					"page": "1",
				},
			},
			want: "https://service.example.com/search?page=1&q=golang",
		},
		{
			name: "keeps query string without trailing slash before question mark",
			enricher: EndpointEnricher{
				Endpoint: config.Endpoint{
					Method:      "GET",
					Path:        "/path/?ok=true",
					Description: "path with query string",
				},
			},
			want: "/path?ok=true",
		},
		{
			name: "returns relative path when no base URL is configured",
			enricher: EndpointEnricher{
				Endpoint: config.Endpoint{
					Method:      "GET",
					Path:        "/health",
					Description: "health check",
				},
			},
			want: "/health",
		},
		{
			name: "returns empty string when endpoint path is blank",
			enricher: EndpointEnricher{
				Endpoint: config.Endpoint{
					Method:      "GET",
					Path:        "   ",
					Description: "invalid",
				},
				ConfiguredBaseURL: "https://api.example.com",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.enricher.GetFullURL())
		})
	}
}
