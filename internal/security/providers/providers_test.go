package providers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"testflowkit/internal/config"
)

func TestBearerProvider(t *testing.T) {
	p := &BearerProvider{}
	result, err := p.Authenticate(context.Background(), config.SecurityScheme{
		Type:  config.SecurityTypeBearer,
		Token: "my-token",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.HeaderValue != "Bearer my-token" {
		t.Errorf("expected 'Bearer my-token', got %q", result.HeaderValue)
	}
}

func TestBasicProvider(t *testing.T) {
	p := &BasicProvider{}
	result, err := p.Authenticate(context.Background(), config.SecurityScheme{
		Type:     config.SecurityTypeBasic,
		Username: "user",
		Password: "pass",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// "user:pass" base64 = "dXNlcjpwYXNz"
	if result.HeaderValue != "Basic dXNlcjpwYXNz" {
		t.Errorf("unexpected basic auth value: %q", result.HeaderValue)
	}
}

func TestAPIKeyProvider_Header(t *testing.T) {
	p := &APIKeyProvider{}
	result, err := p.Authenticate(context.Background(), config.SecurityScheme{
		Type:       config.SecurityTypeAPIKey,
		Key:        "key123",
		Placement:  config.APIKeyPlacementHeader,
		HeaderName: "X-API-Key",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.HeaderName != "X-API-Key" {
		t.Errorf("expected HeaderName=X-API-Key, got %q", result.HeaderName)
	}
	if result.HeaderValue != "key123" {
		t.Errorf("expected HeaderValue=key123, got %q", result.HeaderValue)
	}
}

func TestAPIKeyProvider_Query(t *testing.T) {
	p := &APIKeyProvider{}
	result, err := p.Authenticate(context.Background(), config.SecurityScheme{
		Type:       config.SecurityTypeAPIKey,
		Key:        "key123",
		Placement:  config.APIKeyPlacementQuery,
		QueryParam: "api_key",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.QueryParam != "api_key" {
		t.Errorf("expected QueryParam=api_key, got %q", result.QueryParam)
	}
	if result.AccessToken != "key123" {
		t.Errorf("expected AccessToken=key123, got %q", result.AccessToken)
	}
}

func TestOIDCProvider_NotImplemented(t *testing.T) {
	p := &OIDCProvider{}
	_, err := p.Authenticate(context.Background(), config.SecurityScheme{
		Type: config.SecurityTypeOIDC,
	})
	if !errors.Is(err, ErrNotImplemented) {
		t.Errorf("expected ErrNotImplemented, got %v", err)
	}
}

func TestCertificateProvider_NotImplemented(t *testing.T) {
	p := &CertificateProvider{}
	_, err := p.Authenticate(context.Background(), config.SecurityScheme{
		Type: config.SecurityTypeCertificate,
	})
	if !errors.Is(err, ErrNotImplemented) {
		t.Errorf("expected ErrNotImplemented, got %v", err)
	}
}

func TestFactory_UnknownType(t *testing.T) {
	_, err := Factory(config.SecuritySchemeType("unknown"))
	if err == nil {
		t.Error("expected error for unknown provider type")
	}
}

func TestFactory_AllKnownTypes(t *testing.T) {
	types := []config.SecuritySchemeType{
		config.SecurityTypeBearer,
		config.SecurityTypeBasic,
		config.SecurityTypeAPIKey,
		config.SecurityTypeOAuth2,
		config.SecurityTypeOIDC,
		config.SecurityTypeCertificate,
	}
	for _, typ := range types {
		p, err := Factory(typ)
		if err != nil {
			t.Errorf("Factory(%s) returned unexpected error: %v", typ, err)
		}
		if p == nil {
			t.Errorf("Factory(%s) returned nil provider", typ)
		}
	}
}

// makeTokenServer creates an httptest.Server that responds with a token JSON
// payload and captures the incoming request for later assertion.
func makeTokenServer(t *testing.T, captured **http.Request) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		// Re-attach body so the caller can also parse it via ParseForm.
		r.Body = io.NopCloser(strings.NewReader(string(body)))
		r.PostForm, _ = url.ParseQuery(string(body))
		*captured = r
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": "test-access-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
}

func TestOAuth2Provider_ClientSecretPost(t *testing.T) {
	var captured *http.Request
	srv := makeTokenServer(t, &captured)
	defer srv.Close()

	p := &OAuth2Provider{}
	result, err := p.Authenticate(context.Background(), config.SecurityScheme{
		Type:                    config.SecurityTypeOAuth2,
		TokenURL:                srv.URL,
		ClientID:                "my-client",
		ClientSecret:            "my-secret",
		TokenEndpointAuthMethod: config.OAuth2TokenAuthMethodPost,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.AccessToken != "test-access-token" {
		t.Errorf("expected access token, got %q", result.AccessToken)
	}

	// Credentials must be in the form body.
	if got := captured.PostForm.Get("client_id"); got != "my-client" {
		t.Errorf("expected client_id in body, got %q", got)
	}
	if got := captured.PostForm.Get("client_secret"); got != "my-secret" {
		t.Errorf("expected client_secret in body, got %q", got)
	}
	// Must NOT set Authorization header for this method.
	if h := captured.Header.Get("Authorization"); h != "" {
		t.Errorf("expected no Authorization header for client_secret_post, got %q", h)
	}
}

func TestOAuth2Provider_ClientSecretBasic(t *testing.T) {
	var captured *http.Request
	srv := makeTokenServer(t, &captured)
	defer srv.Close()

	p := &OAuth2Provider{}
	_, err := p.Authenticate(context.Background(), config.SecurityScheme{
		Type:                    config.SecurityTypeOAuth2,
		TokenURL:                srv.URL,
		ClientID:                "my-client",
		ClientSecret:            "my-secret",
		TokenEndpointAuthMethod: config.OAuth2TokenAuthMethodBasic,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Authorization header must be Basic base64(client_id:client_secret).
	wantEncoded := base64.StdEncoding.EncodeToString([]byte("my-client:my-secret"))
	wantHeader := "Basic " + wantEncoded
	if got := captured.Header.Get("Authorization"); got != wantHeader {
		t.Errorf("expected Authorization=%q, got %q", wantHeader, got)
	}
	// Strict RFC: neither client_id nor client_secret should be in the form body.
	if got := captured.PostForm.Get("client_id"); got != "" {
		t.Errorf("client_id must not appear in form body for client_secret_basic, got %q", got)
	}
	if got := captured.PostForm.Get("client_secret"); got != "" {
		t.Errorf("client_secret must not appear in form body for client_secret_basic, got %q", got)
	}
}
