package security

import (
	"testing"

	"testflowkit/internal/config"
)

func TestSchemeHash_Deterministic(t *testing.T) {
	s := config.SecurityScheme{
		Type:         config.SecurityTypeOAuth2,
		ClientID:     "client_id",
		ClientSecret: "secret",
		TokenURL:     "https://auth.example.com/token",
		Scopes:       []string{"read", "write"},
	}
	h1, err1 := SchemeHash(s)
	h2, err2 := SchemeHash(s)
	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected errors: %v, %v", err1, err2)
	}
	if h1 != h2 {
		t.Errorf("SchemeHash is not deterministic: %q != %q", h1, h2)
	}
}

func TestSchemeHash_DifferentSchemes(t *testing.T) {
	a := config.SecurityScheme{Type: config.SecurityTypeBearer, Token: "token-a"}
	b := config.SecurityScheme{Type: config.SecurityTypeBearer, Token: "token-b"}
	ha, _ := SchemeHash(a)
	hb, _ := SchemeHash(b)
	if ha == hb {
		t.Error("different schemes should produce different hashes")
	}
}

func TestSchemeHash_EmptyScheme(t *testing.T) {
	h, err := SchemeHash(config.SecurityScheme{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h == "" {
		t.Error("hash should be non-empty even for empty scheme")
	}
}

func TestSchemeHash_Length(t *testing.T) {
	s := config.SecurityScheme{Type: config.SecurityTypeBasic, Username: "u", Password: "p"}
	h, err := SchemeHash(s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// SHA-256 hex is 64 characters
	const want = 64
	if len(h) != want {
		t.Errorf("expected hash length %d, got %d", want, len(h))
	}
}
