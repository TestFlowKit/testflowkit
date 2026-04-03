package security

import (
	"testing"

	"testflowkit/internal/config"
)

func makeConfig(defaultSecurity string, schemes map[string]config.SecurityScheme) *config.Config {
	return &config.Config{
		SecuritySchemes: schemes,
		DefaultSecurity: defaultSecurity,
	}
}

func TestResolve_NoSecurityAtAnyLevel(t *testing.T) {
	cfg := makeConfig("", nil)
	got := Resolve(cfg, nil, nil)
	if got.Disabled {
		t.Error("expected Disabled=false when no security configured")
	}
	if got.SchemeName != "" {
		t.Errorf("expected empty SchemeName, got %q", got.SchemeName)
	}
}

func TestResolve_ProjectDefault(t *testing.T) {
	bearer := config.SecurityScheme{Type: config.SecurityTypeBearer, Token: "tok"}
	cfg := makeConfig("default_bearer", map[string]config.SecurityScheme{
		"default_bearer": bearer,
	})
	got := Resolve(cfg, nil, nil)
	if got.Disabled {
		t.Error("expected Disabled=false")
	}
	if got.SchemeName != "default_bearer" {
		t.Errorf("expected SchemeName=default_bearer, got %q", got.SchemeName)
	}
	if got.Scheme.Token != "tok" {
		t.Errorf("expected token=tok, got %q", got.Scheme.Token)
	}
}

func TestResolve_ProjectDefault_None(t *testing.T) {
	cfg := makeConfig("none", nil)
	got := Resolve(cfg, nil, nil)
	if !got.Disabled {
		t.Error("expected Disabled=true for default_security: none")
	}
}

func TestResolve_APILevelRef(t *testing.T) {
	apiScheme := config.SecurityScheme{Type: config.SecurityTypeBasic, Username: "user", Password: "pass"}
	cfg := makeConfig("", map[string]config.SecurityScheme{
		"api_scheme": apiScheme,
	})
	apiDef := &config.APIDefinition{
		SecurityRef: &config.SecurityRef{Name: "api_scheme"},
	}
	got := Resolve(cfg, apiDef, nil)
	if got.SchemeName != "api_scheme" {
		t.Errorf("expected SchemeName=api_scheme, got %q", got.SchemeName)
	}
	if got.Scheme.Username != "user" {
		t.Errorf("expected Username=user, got %q", got.Scheme.Username)
	}
}

func TestResolve_APILevelRef_None(t *testing.T) {
	cfg := makeConfig("some_default", map[string]config.SecurityScheme{
		"some_default": {Type: config.SecurityTypeBearer, Token: "t"},
	})
	apiDef := &config.APIDefinition{
		SecurityRef: &config.SecurityRef{Name: "none"},
	}
	got := Resolve(cfg, apiDef, nil)
	if !got.Disabled {
		t.Error("expected Disabled=true for api-level security: none")
	}
}

func TestResolve_RequestLevelRef_OverridesAll(t *testing.T) {
	apiScheme := config.SecurityScheme{Type: config.SecurityTypeBasic, Username: "api_user"}
	reqScheme := config.SecurityScheme{Type: config.SecurityTypeBearer, Token: "req_token"}
	cfg := makeConfig("", map[string]config.SecurityScheme{
		"api_scheme": apiScheme,
		"req_scheme": reqScheme,
	})
	apiDef := &config.APIDefinition{
		SecurityRef: &config.SecurityRef{Name: "api_scheme"},
	}
	requestRef := &config.SecurityRef{Name: "req_scheme"}
	got := Resolve(cfg, apiDef, requestRef)
	if got.SchemeName != "req_scheme" {
		t.Errorf("expected SchemeName=req_scheme, got %q", got.SchemeName)
	}
	if got.Scheme.Token != "req_token" {
		t.Errorf("expected Token=req_token, got %q", got.Scheme.Token)
	}
}

func TestResolve_APILevelOverrides_Applied(t *testing.T) {
	base := config.SecurityScheme{
		Type:     config.SecurityTypeOAuth2,
		Scopes:   []string{"read"},
		Audience: "aud1",
	}
	cfg := makeConfig("base_oauth", map[string]config.SecurityScheme{
		"base_oauth": base,
	})
	apiDef := &config.APIDefinition{
		SecurityRef: &config.SecurityRef{Name: "base_oauth"},
		SecurityOverrides: &config.SecurityOverrides{
			Scopes:   []string{"read", "write"},
			Audience: "aud2",
		},
	}
	got := Resolve(cfg, apiDef, nil)
	if len(got.Scheme.Scopes) != 2 {
		t.Errorf("expected 2 scopes, got %d", len(got.Scheme.Scopes))
	}
	if got.Scheme.Audience != "aud2" {
		t.Errorf("expected Audience=aud2, got %q", got.Scheme.Audience)
	}
}

func TestResolve_ProjectDefault_WithAPIOverrides(t *testing.T) {
	base := config.SecurityScheme{
		Type:   config.SecurityTypeOAuth2,
		Scopes: []string{"read"},
	}
	cfg := makeConfig("proj_oauth", map[string]config.SecurityScheme{
		"proj_oauth": base,
	})
	// API has no SecurityRef but has SecurityOverrides
	apiDef := &config.APIDefinition{
		SecurityOverrides: &config.SecurityOverrides{
			Scopes: []string{"read", "admin"},
		},
	}
	got := Resolve(cfg, apiDef, nil)
	if got.SchemeName != "proj_oauth" {
		t.Errorf("expected SchemeName=proj_oauth, got %q", got.SchemeName)
	}
	if len(got.Scheme.Scopes) != 2 {
		t.Errorf("expected 2 scopes from override, got %d: %v", len(got.Scheme.Scopes), got.Scheme.Scopes)
	}
}
