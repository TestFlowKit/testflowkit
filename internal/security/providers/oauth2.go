package providers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testflowkit/internal/config"
	"time"
)

// OAuth2Provider handles type: oauth2.
// The method used to send client credentials to the token endpoint is
// determined by TokenEndpointAuthMethod (required) and resolved via the
// tokenEndpointAuthHandler strategy.
type OAuth2Provider struct{}

type oauth2TokenResponse struct {
	AccessToken string `json:"access_token"` //nolint:gosec // OAuth2 response schema uses this field name.
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

const tokenRequestTimeout = 30 * time.Second

// tokenEndpointAuthHandler is the strategy interface for injecting OAuth2
// client credentials. Each method owns exactly one concern:
//   - ApplyToForm adds credential fields to the form body (before encoding).
//   - ApplyToRequest adds credential headers to the prepared HTTP request.
//
// Implementations must be stateless.
type tokenEndpointAuthHandler interface {
	ApplyToForm(form url.Values, scheme config.SecurityScheme)
	ApplyToRequest(req *http.Request, scheme config.SecurityScheme)
}

// clientSecretPostHandler implements client_secret_post (RFC 6749 §2.3.1):
// client_id and client_secret are sent as form body parameters.
type clientSecretPostHandler struct{}

func (clientSecretPostHandler) ApplyToForm(form url.Values, scheme config.SecurityScheme) {
	form.Set("client_id", scheme.ClientID)
	form.Set("client_secret", scheme.ClientSecret)
}

func (clientSecretPostHandler) ApplyToRequest(_ *http.Request, _ config.SecurityScheme) {}

// clientSecretBasicHandler implements client_secret_basic (RFC 6749 §2.3.1 alt):
// credentials are sent in the Authorization header as HTTP Basic; the body
// carries no credential fields at all (strict RFC compliance).
type clientSecretBasicHandler struct{}

func (clientSecretBasicHandler) ApplyToForm(_ url.Values, _ config.SecurityScheme) {}

func (clientSecretBasicHandler) ApplyToRequest(req *http.Request, scheme config.SecurityScheme) {
	raw := scheme.ClientID + ":" + scheme.ClientSecret
	encoded := base64.StdEncoding.EncodeToString([]byte(raw))
	req.Header.Set("Authorization", "Basic "+encoded)
}

// authHandlerFor returns the tokenEndpointAuthHandler for the given method.
// Validation guarantees the method is always a known value when this is called.
func authHandlerFor(method config.OAuth2TokenAuthMethod) tokenEndpointAuthHandler {
	if method == config.OAuth2TokenAuthMethodBasic {
		return clientSecretBasicHandler{}
	}
	return clientSecretPostHandler{}
}

func (*OAuth2Provider) Authenticate(ctx context.Context, scheme config.SecurityScheme) (*TokenResult, error) {
	client, err := buildHTTPClient(scheme.ProxyURL)
	if err != nil {
		return nil, fmt.Errorf("oauth2: build http client: %w", err)
	}

	// Build base form — only protocol fields, no credentials yet.
	form := url.Values{}
	form.Set("grant_type", grantType(scheme))
	if len(scheme.Scopes) > 0 {
		form.Set("scope", strings.Join(scheme.Scopes, " "))
	}
	if scheme.Audience != "" {
		form.Set("audience", scheme.Audience)
	}

	// Resolve the credential strategy and inject into form body.
	h := authHandlerFor(scheme.TokenEndpointAuthMethod)
	h.ApplyToForm(form, scheme)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, scheme.TokenURL,
		strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("oauth2: create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Inject any request-level credential headers.
	h.ApplyToRequest(req, scheme)

	resp, err := client.Do(req) //nolint:gosec // Token endpoint URL is user-configured and intentionally requested.
	if err != nil {
		return nil, fmt.Errorf("oauth2: token request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("oauth2: read token response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("oauth2: token endpoint returned %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp oauth2TokenResponse
	decodeErr := json.Unmarshal(body, &tokenResp)
	if decodeErr != nil {
		return nil, fmt.Errorf("oauth2: parse token response: %w", decodeErr)
	}
	if tokenResp.AccessToken == "" {
		return nil, errors.New("oauth2: token response missing access_token")
	}

	tokenType := tokenResp.TokenType
	if tokenType == "" {
		tokenType = "Bearer"
	}

	return &TokenResult{
		AccessToken: tokenResp.AccessToken,
		TokenType:   tokenType,
		ExpiresIn:   tokenResp.ExpiresIn,
		HeaderValue: tokenType + " " + tokenResp.AccessToken,
	}, nil
}

func grantType(scheme config.SecurityScheme) string {
	if scheme.GrantType != "" {
		return scheme.GrantType
	}
	return "client_credentials"
}

func buildHTTPClient(proxyURL string) (*http.Client, error) {
	transport := &http.Transport{}
	if proxyURL != "" {
		parsed, err := url.Parse(proxyURL)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy_url '%s': %w", proxyURL, err)
		}
		transport.Proxy = http.ProxyURL(parsed)
	}
	return &http.Client{
		Transport: transport,
		Timeout:   tokenRequestTimeout,
	}, nil
}
