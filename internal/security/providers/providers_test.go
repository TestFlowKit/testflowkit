package providers

import (
	"context"
	"errors"
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
		APIKey:     "key123",
		Placement:  config.APIKeyHeader,
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
		APIKey:     "key123",
		Placement:  config.APIKeyQuery,
		QueryParam: "api_key",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.QueryParam != "api_key" {
		t.Errorf("expected QueryParam=api_key, got %q", result.QueryParam)
	}
	if result.HeaderValue != "key123" {
		t.Errorf("expected HeaderValue=key123, got %q", result.HeaderValue)
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
