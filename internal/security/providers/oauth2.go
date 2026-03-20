package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testflowkit/internal/config"
	"time"
)

// OAuth2Provider handles type: oauth2 with grant_type: client_credentials.
// It posts to the token_url and returns a bearer token.
// Per-scheme proxy support is applied by dialling through scheme.ProxyURL when set.
type OAuth2Provider struct{}

type oauth2TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

func (*OAuth2Provider) Authenticate(ctx context.Context, scheme config.SecurityScheme) (*TokenResult, error) {
	client, err := buildHTTPClient(scheme.ProxyURL)
	if err != nil {
		return nil, fmt.Errorf("oauth2: build http client: %w", err)
	}

	form := url.Values{}
	form.Set("grant_type", grantType(scheme))
	form.Set("client_id", scheme.ClientID)
	form.Set("client_secret", scheme.ClientSecret)
	if len(scheme.Scopes) > 0 {
		form.Set("scope", strings.Join(scheme.Scopes, " "))
	}
	if scheme.Audience != "" {
		form.Set("audience", scheme.Audience)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, scheme.TokenURL,
		strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("oauth2: create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
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
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("oauth2: parse token response: %w", err)
	}
	if tokenResp.AccessToken == "" {
		return nil, fmt.Errorf("oauth2: token response missing access_token")
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
		Timeout:   30 * time.Second,
	}, nil
}
