// Package security provides the inheritance resolver for TestFlowKit's
// unified security framework.  It translates the three-level config hierarchy
// (endpoint → API → project) into a single, flat ResolvedSecurity value that
// the auth providers and transport layer consume.
package security

import (
	"testflowkit/internal/config"
)

// ResolvedSecurity is the fully-merged security context for one request.
// It is computed once during request preparation and is immutable afterwards.
type ResolvedSecurity struct {
	// Scheme is the effective SecurityScheme after applying overrides.
	Scheme config.SecurityScheme
	// SchemeName is the name of the root scheme (empty for inline/anonymous).
	SchemeName string
	// Disabled is true when security: none was encountered at any level.
	Disabled bool
}

// Resolve returns the effective SecurityScheme for a specific API endpoint or
// GraphQL operation, applying the resolution order:
//
//  1. Request level  – endpoint/operation SecurityRef
//  2. API level      – APIDefinition SecurityRef + SecurityOverrides
//  3. Project level  – Config.DefaultSecurity
//
// A SecurityRef with Name == "none" at any level stops the chain and returns
// a ResolvedSecurity with Disabled == true.
func Resolve(
	cfg *config.Config,
	apiDef *config.APIDefinition,
	requestRef *config.SecurityRef,
) ResolvedSecurity {
	// 1. Request-level ref takes absolute precedence.
	if !requestRef.IsEmpty() {
		if requestRef.IsNone() {
			return ResolvedSecurity{Disabled: true}
		}
		base, name := resolveRef(cfg, requestRef)
		return ResolvedSecurity{Scheme: base, SchemeName: name}
	}

	// 2. API-level ref (+ optional overrides).
	if apiDef != nil && !apiDef.SecurityRef.IsEmpty() {
		if apiDef.SecurityRef.IsNone() {
			return ResolvedSecurity{Disabled: true}
		}
		base, name := resolveRef(cfg, apiDef.SecurityRef)
		if apiDef.SecurityOverrides != nil {
			base = applyOverrides(base, apiDef.SecurityOverrides)
		}
		return ResolvedSecurity{Scheme: base, SchemeName: name}
	}

	// 3. Project default.
	if cfg.DefaultSecurity != "" {
		if cfg.DefaultSecurity == string(config.SecurityTypeNone) {
			return ResolvedSecurity{Disabled: true}
		}
		scheme, ok := cfg.GetSecurityScheme(cfg.DefaultSecurity)
		if ok {
			// API-level overrides still apply even when falling back to project default.
			if apiDef != nil && apiDef.SecurityOverrides != nil {
				scheme = applyOverrides(scheme, apiDef.SecurityOverrides)
			}
			return ResolvedSecurity{Scheme: scheme, SchemeName: cfg.DefaultSecurity}
		}
	}

	// No security configured at any level.
	return ResolvedSecurity{}
}

// resolveRef extracts the concrete SecurityScheme and name from a SecurityRef.
func resolveRef(cfg *config.Config, ref *config.SecurityRef) (config.SecurityScheme, string) {
	if ref.Inline != nil {
		return *ref.Inline, ""
	}
	scheme, _ := cfg.GetSecurityScheme(ref.Name)
	return scheme, ref.Name
}

// applyOverrides creates a shallow copy of base and patches the fields present
// in overrides.  Only non-zero override values are applied so that partial
// overrides (e.g. only scopes) leave the rest of the base scheme intact.
func applyOverrides(base config.SecurityScheme, ov *config.SecurityOverrides) config.SecurityScheme {
	if ov == nil {
		return base
	}
	if len(ov.Scopes) > 0 {
		base.Scopes = ov.Scopes
	}
	if ov.Audience != "" {
		base.Audience = ov.Audience
	}
	if ov.ProxyURL != "" {
		base.ProxyURL = ov.ProxyURL
	}
	return base
}
