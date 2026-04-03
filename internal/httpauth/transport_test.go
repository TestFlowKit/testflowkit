package httpauth

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"testflowkit/internal/config"
	"testflowkit/internal/security"
)

type recordingRoundTripper struct {
	requests  []*http.Request
	bodies    []string
	responses []*http.Response
}

func (r *recordingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	r.requests = append(r.requests, req.Clone(req.Context()))

	bodyText := ""
	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		bodyText = string(body)
		req.Body = io.NopCloser(bytes.NewReader(body))
	}
	r.bodies = append(r.bodies, bodyText)

	if len(r.responses) == 0 {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("")),
		}, nil
	}

	resp := r.responses[0]
	r.responses = r.responses[1:]
	if resp.Header == nil {
		resp.Header = make(http.Header)
	}
	if resp.Body == nil {
		resp.Body = io.NopCloser(strings.NewReader(""))
	}
	return resp, nil
}

func TestAuthTransport_RoundTrip_BypassesAuthWhenDisabled(t *testing.T) {
	base := &recordingRoundTripper{}
	transport := &AuthTransport{
		Base: base,
		Resolved: security.ResolvedSecurity{
			Disabled: true,
		},
	}

	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}

	resp, err := transport.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	if len(base.requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(base.requests))
	}
	if got := base.requests[0].Header.Get("Authorization"); got != "" {
		t.Fatalf("expected no Authorization header, got %q", got)
	}
}

func TestAuthTransport_RoundTrip_InjectsBearerHeader(t *testing.T) {
	base := &recordingRoundTripper{}
	transport := &AuthTransport{
		Base: base,
		Resolved: security.ResolvedSecurity{
			Scheme: config.SecurityScheme{
				Type:  config.SecurityTypeBearer,
				Token: "secret-token",
			},
		},
	}

	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}

	_, err = transport.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(base.requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(base.requests))
	}
	if got := base.requests[0].Header.Get("Authorization"); got != "Bearer secret-token" {
		t.Fatalf("expected bearer header to be injected, got %q", got)
	}
	if got := req.Header.Get("Authorization"); got != "" {
		t.Fatalf("expected original request to remain unchanged, got %q", got)
	}
}

func TestAuthTransport_RoundTrip_InjectsAPIKeyQueryParam(t *testing.T) {
	base := &recordingRoundTripper{}
	transport := &AuthTransport{
		Base: base,
		Resolved: security.ResolvedSecurity{
			Scheme: config.SecurityScheme{
				Type:       config.SecurityTypeAPIKey,
				Key:        "api-secret",
				Placement:  config.APIKeyPlacementQuery,
				QueryParam: "token",
			},
		},
	}

	req, err := http.NewRequest(http.MethodGet, "http://example.com/items", nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}

	_, err = transport.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(base.requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(base.requests))
	}
	if got := base.requests[0].URL.Query().Get("token"); got != "api-secret" {
		t.Fatalf("expected query token to be injected, got %q", got)
	}
}

func TestAuthTransport_RetryOnUnauthorized_RetriesOnceWithBody(t *testing.T) {
	base := &recordingRoundTripper{
		responses: []*http.Response{
			{StatusCode: http.StatusUnauthorized},
			{StatusCode: http.StatusOK},
		},
	}
	transport := &AuthTransport{
		Base: base,
		Resolved: security.ResolvedSecurity{
			Scheme: config.SecurityScheme{
				Type:       config.SecurityTypeBearer,
				Token:      "retry-token",
				RetryOn401: true,
			},
		},
	}

	req, err := http.NewRequest(http.MethodPost, "http://example.com/retry", bytes.NewBufferString(`{"hello":"world"}`))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}

	resp, err := transport.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected final status %d, got %d", http.StatusOK, resp.StatusCode)
	}
	if len(base.requests) != 2 {
		t.Fatalf("expected 2 requests after retry, got %d", len(base.requests))
	}
	if got := base.requests[0].Header.Get("Authorization"); got != "Bearer retry-token" {
		t.Fatalf("expected first request Authorization header, got %q", got)
	}
	if got := base.requests[1].Header.Get("Authorization"); got != "Bearer retry-token" {
		t.Fatalf("expected retry request Authorization header, got %q", got)
	}
	if len(base.bodies) != 2 || base.bodies[0] != `{"hello":"world"}` || base.bodies[1] != `{"hello":"world"}` {
		t.Fatalf("expected request body to be restored on retry, got %v", base.bodies)
	}
}

func TestNewBaseTransport_UsesProxyFromEnvironmentByDefault(t *testing.T) {
	t.Setenv("HTTP_PROXY", "http://proxy.example:8080")
	t.Setenv("HTTPS_PROXY", "")
	t.Setenv("NO_PROXY", "")

	tr, err := NewBaseTransport("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tr.Proxy == nil {
		t.Fatal("expected proxy function to be configured")
	}

	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}

	proxyURL, err := tr.Proxy(req)
	if err != nil {
		t.Fatalf("unexpected proxy resolution error: %v", err)
	}
	if proxyURL == nil || proxyURL.String() != "http://proxy.example:8080" {
		t.Fatalf("expected env proxy to be used, got %v", proxyURL)
	}
}

func TestNewBaseTransport_UsesExplicitProxyWhenProvided(t *testing.T) {
	t.Setenv("HTTP_PROXY", "http://proxy.example:8080")
	t.Setenv("HTTPS_PROXY", "")
	t.Setenv("NO_PROXY", "")

	tr, err := NewBaseTransport("http://custom-proxy:9090")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tr.Proxy == nil {
		t.Fatal("expected proxy function to be configured")
	}

	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}

	proxyURL, err := tr.Proxy(req)
	if err != nil {
		t.Fatalf("unexpected proxy resolution error: %v", err)
	}
	if proxyURL == nil || proxyURL.String() != "http://custom-proxy:9090" {
		t.Fatalf("expected explicit proxy to be used, got %v", proxyURL)
	}
}
