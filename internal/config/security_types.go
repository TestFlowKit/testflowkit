package config

import "time"

// SecuritySchemeType enumerates supported auth mechanisms.
type SecuritySchemeType string

const (
	// SecurityTypeBearer uses a static Bearer token in the Authorization header.
	SecurityTypeBearer SecuritySchemeType = "bearer"

	// SecurityTypeBasic uses HTTP Basic auth (username + password).
	SecurityTypeBasic SecuritySchemeType = "basic"

	// SecurityTypeAPIKey injects a key into a header, query param, or cookie.
	SecurityTypeAPIKey SecuritySchemeType = "apikey"

	// SecurityTypeOAuth2 obtains a token from a token endpoint (client credentials).
	SecurityTypeOAuth2 SecuritySchemeType = "oauth2"

	// SecurityTypeOIDC obtains tokens via OpenID Connect (not yet fully implemented).
	SecurityTypeOIDC SecuritySchemeType = "oidc"

	// SecurityTypeCertificate uses mutual TLS (cert_file + key_file).
	SecurityTypeCertificate SecuritySchemeType = "certificate"

	// SecurityTypeNone disables inherited auth (use as security_ref.name: "none").
	SecurityTypeNone SecuritySchemeType = "none"
)

// APIKeyPlacement controls where the API key is attached to the request.
type APIKeyPlacement string

const (
	// APIKeyPlacementHeader injects the key as an HTTP header (default).
	APIKeyPlacementHeader APIKeyPlacement = "header"

	// APIKeyPlacementQuery appends the key as a URL query parameter.
	APIKeyPlacementQuery APIKeyPlacement = "query"

	// APIKeyPlacementCookie sends the key as a cookie.
	APIKeyPlacementCookie APIKeyPlacement = "cookie"
)

// OAuth2TokenAuthMethod controls how client credentials are sent to the token endpoint.
type OAuth2TokenAuthMethod string

const (
	// OAuth2TokenAuthMethodPost sends credentials as form body params.
	OAuth2TokenAuthMethodPost OAuth2TokenAuthMethod = "client_secret_post"

	// OAuth2TokenAuthMethodBasic sends credentials via HTTP Basic auth header.
	OAuth2TokenAuthMethodBasic OAuth2TokenAuthMethod = "client_secret_basic"
)

// SecurityScheme is a reusable named auth config in security_schemes.
type SecurityScheme struct {
	// Type is the auth mechanism: bearer, basic, apikey, oauth2, oidc, certificate, none.
	Type SecuritySchemeType `yaml:"type" json:"type"`

	// Token is the Bearer/access token for type "bearer".
	Token string `yaml:"token" json:"token"`

	// Username is the account ID for HTTP Basic auth.
	Username string `yaml:"username" json:"username"`

	// Password is the credential for HTTP Basic auth.
	Password string `yaml:"password" json:"password"`

	// Key is the API key value for type "apikey".
	Key string `yaml:"key" json:"key"`

	// Placement is where the key is injected: header (default), query, or cookie.
	Placement APIKeyPlacement `yaml:"placement" json:"placement"`

	// HeaderName is the HTTP header for the API key (default: Authorization).
	HeaderName string `yaml:"header_name" json:"header_name"`

	// QueryParam is the URL query parameter name for the API key.
	QueryParam string `yaml:"query_param" json:"query_param"`

	// GrantType is the OAuth 2.0 grant type, e.g. "client_credentials".
	GrantType string `yaml:"grant_type" json:"grant_type"`

	// TokenURL is the OAuth 2.0 / OIDC token endpoint URL.
	TokenURL string `yaml:"token_url" json:"token_url"`

	// ClientID is the OAuth 2.0 application identifier.
	ClientID string `yaml:"client_id" json:"client_id"`

	// ClientSecret is the OAuth 2.0 application secret.
	ClientSecret string `yaml:"client_secret" json:"client_secret"`

	// Scopes is the list of OAuth 2.0 / OIDC scopes to request.
	Scopes []string `yaml:"scopes" json:"scopes"`

	// Audience is the intended token recipient (the "aud" claim).
	Audience string `yaml:"audience" json:"audience"`

	// TokenEndpointAuthMethod is how credentials are sent: client_secret_post or client_secret_basic.
	TokenEndpointAuthMethod OAuth2TokenAuthMethod `yaml:"token_endpoint_auth_method" json:"token_endpoint_auth_method"`

	// CertFile is the path to a PEM TLS client certificate.
	CertFile string `yaml:"cert_file" json:"cert_file"`

	// KeyFile is the path to the PEM private key paired with cert_file.
	KeyFile string `yaml:"key_file" json:"key_file"`

	// ProxyURL is an optional HTTP/HTTPS proxy for this scheme.
	ProxyURL string `yaml:"proxy_url" json:"proxy_url"`

	// Duration overrides the token TTL, e.g. "1h" or "30m".
	Duration string `yaml:"duration" json:"duration"`

	// RetryOn401 retries once after invalidating the cached token on HTTP 401.
	RetryOn401 bool `yaml:"retry_on_401" json:"retry_on_401"`
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

// SecurityRef selects the auth scheme for an API, endpoint, or operation.
type SecurityRef struct {
	// Name references a scheme in security_schemes; use "none" to disable auth.
	Name string `yaml:"name"`

	// Inline defines an anonymous scheme at the use-site.
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

// SecurityOverrides patches fields of the resolved security scheme for one API.
type SecurityOverrides struct {
	// Scopes replaces the base scheme's OAuth scope list for this API.
	Scopes []string `yaml:"scopes"`

	// Audience replaces the base scheme's OAuth audience for this API.
	Audience string `yaml:"audience"`

	// ProxyURL replaces the base scheme's proxy URL for this API.
	ProxyURL string `yaml:"proxy_url"`
}
