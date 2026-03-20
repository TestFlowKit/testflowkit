package config

import "time"

// SecuritySchemeType enumerates all supported auth mechanisms.
type SecuritySchemeType string

const (
	SecurityTypeBearer      SecuritySchemeType = "bearer"
	SecurityTypeBasic       SecuritySchemeType = "basic"
	SecurityTypeAPIKey      SecuritySchemeType = "apikey"
	SecurityTypeOAuth2      SecuritySchemeType = "oauth2"
	SecurityTypeOIDC        SecuritySchemeType = "oidc"
	SecurityTypeCertificate SecuritySchemeType = "certificate"
	// SecurityTypeNone is a sentinel value that disables all inherited security
	// for a specific endpoint or API definition.
	SecurityTypeNone SecuritySchemeType = "none"
)

// APIKeyPlacement controls where the API key is attached to the request.
type APIKeyPlacement string

const (
	APIKeyPlacementHeader APIKeyPlacement = "header"
	APIKeyPlacementQuery  APIKeyPlacement = "query"
	APIKeyPlacementCookie APIKeyPlacement = "cookie"
)

// SecurityScheme is a reusable, named authentication configuration stored in
// the root security_schemes registry. All {{ env.* }} references are resolved
// before this struct is populated, making its field values suitable for
// deterministic SHA-256 hashing.
type SecurityScheme struct {
	// Type is the auth mechanism (required).
	Type SecuritySchemeType `yaml:"type"`

	// --- Bearer / static token ---
	// Token is the raw bearer/access token for type: bearer.
	Token string `yaml:"token"`

	// --- Basic auth ---
	Username string `yaml:"username"`
	Password string `yaml:"password"`

	// --- API Key ---
	Key        string          `yaml:"key"`
	Placement  APIKeyPlacement `yaml:"placement"`   // header | query | cookie (default: header)
	HeaderName string          `yaml:"header_name"` // e.g. "X-Api-Key" (default: Authorization)
	QueryParam string          `yaml:"query_param"` // e.g. "api_key" (for placement: query)

	// --- OAuth2 / OIDC ---
	GrantType    string   `yaml:"grant_type"` // client_credentials | password | …
	TokenURL     string   `yaml:"token_url"`
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	Scopes       []string `yaml:"scopes"`
	Audience     string   `yaml:"audience"`

	// --- Certificate / mTLS ---
	// CertFile and KeyFile are paths to PEM certificate and private-key files.
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`

	// --- Transport ---
	// ProxyURL is an optional per-scheme HTTP/HTTPS proxy (e.g. enterprise egress).
	// Overrides any system proxy for token-fetch requests.
	ProxyURL string `yaml:"proxy_url"`

	// --- Persistence & lifecycle ---
	// Persist controls whether the obtained token is written to lock file.
	Persist bool `yaml:"persist"`
	// Duration overrides the token TTL reported by the IDP (e.g. "1h", "30m").
	// Accepts any Go duration string understood by time.ParseDuration.
	Duration string `yaml:"duration"`
	// RetryOn401 enables a single invalidate-and-retry cycle when the API
	// returns 401. Disabled by default.
	RetryOn401 bool `yaml:"retry_on_401"`
}

// ParsedDuration parses the Duration field into a time.Duration.
// Returns 0 and false when Duration is empty or unparseable.
func (s *SecurityScheme) ParsedDuration() (time.Duration, bool) {
	if s.Duration == "" {
		return 0, false
	}
	d, err := time.ParseDuration(s.Duration)
	if err != nil {
		return 0, false
	}
	return d, true
}

// SecurityRef is an inline reference to a named SecurityScheme.
// Only one of Name or Inline should be populated per call site.
type SecurityRef struct {
	// Name references an entry in Config.SecuritySchemes.
	Name string `yaml:"name"`
	// Inline allows an anonymous scheme defined directly at the use-site.
	// Prefer named references for reusability.
	Inline *SecurityScheme `yaml:"inline"`
}

// IsNone reports whether this ref explicitly disables inherited security.
func (r *SecurityRef) IsNone() bool {
	return r != nil && r.Name == string(SecurityTypeNone)
}

// IsEmpty reports whether no security configuration is provided.
func (r *SecurityRef) IsEmpty() bool {
	return r == nil || (r.Name == "" && r.Inline == nil)
}

// SecurityOverrides carries API-level adjustments applied on top of a
// referenced SecurityScheme. They do not replace the base scheme; they patch
// specific fields (e.g. restricting scopes for one API without a new scheme).
type SecurityOverrides struct {
	Scopes   []string `yaml:"scopes"`
	Audience string   `yaml:"audience"`
	ProxyURL string   `yaml:"proxy_url"`
}
