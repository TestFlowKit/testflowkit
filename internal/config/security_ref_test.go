package config

import (
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
)

func TestSecurityRefUnmarshal_ObjectName(t *testing.T) {
	var cfg struct {
		SecurityRef *SecurityRef `yaml:"security_ref"`
	}

	err := yaml.Unmarshal([]byte("security_ref:\n  name: my_scheme\n"), &cfg)
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if cfg.SecurityRef == nil {
		t.Fatal("expected security_ref to be populated")
	}
	if cfg.SecurityRef.Name != "my_scheme" {
		t.Fatalf("expected name my_scheme, got %q", cfg.SecurityRef.Name)
	}
}

func TestSecurityRefUnmarshal_NoneName(t *testing.T) {
	var cfg struct {
		SecurityRef *SecurityRef `yaml:"security_ref"`
	}

	err := yaml.Unmarshal([]byte("security_ref:\n  name: none\n"), &cfg)
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if cfg.SecurityRef == nil {
		t.Fatal("expected security_ref to be populated")
	}
	if !cfg.SecurityRef.IsNone() {
		t.Fatal("expected security_ref to resolve as none sentinel")
	}
}

func TestSecurityRefUnmarshal_Inline(t *testing.T) {
	var cfg struct {
		SecurityRef *SecurityRef `yaml:"security_ref"`
	}

	err := yaml.Unmarshal([]byte("security_ref:\n  inline:\n    type: bearer\n    token: abc\n"), &cfg)
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if cfg.SecurityRef == nil || cfg.SecurityRef.Inline == nil {
		t.Fatal("expected inline security_ref")
	}
	if cfg.SecurityRef.Inline.Type != SecurityTypeBearer {
		t.Fatalf("expected inline type bearer, got %q", cfg.SecurityRef.Inline.Type)
	}
	if cfg.SecurityRef.Inline.Token != "abc" {
		t.Fatalf("expected inline token abc, got %q", cfg.SecurityRef.Inline.Token)
	}
}

func TestValidateSecurityRef_RejectsNameAndInline(t *testing.T) {
	conf := &Config{
		SecuritySchemes: map[string]SecurityScheme{
			"my_scheme": {Type: SecurityTypeBearer, Token: "token"},
		},
	}

	err := conf.validateSecurityRef(&SecurityRef{
		Name: "my_scheme",
		Inline: &SecurityScheme{
			Type:  SecurityTypeBearer,
			Token: "inline-token",
		},
	}, "API 'sample'")

	if err == nil {
		t.Fatal("expected validation error when both name and inline are set")
	}
	if !strings.Contains(err.Error(), "cannot set both name and inline") {
		t.Fatalf("unexpected error: %v", err)
	}
}
