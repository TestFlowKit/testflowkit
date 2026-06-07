// Package httpauth provides AuthTransport, a composable http.RoundTripper that:
//   - resolves auth credentials via the provider layer
//   - injects them into outgoing requests (header / query / cookie)
//   - supports per-scheme http/https proxy
//   - optionally retries once on a 401 response (opt-in, disabled by default)
package httpauth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"testflowkit/internal/config"
	"testflowkit/internal/security"
	"testflowkit/internal/security/providers"
	"testflowkit/pkg/formatter"
)

// AuthTransport wraps a base http.RoundTripper and adds auth injection.
// It is intended to be the innermost user-facing transport; TLS, proxy, and
// timeout concerns are handled at the base transport level.
type AuthTransport struct {
	// Base is the underlying transport used to actually send requests.
	// Defaults to http.DefaultTransport when nil.
	Base http.RoundTripper

	// Resolved is the effective security context for this request, computed
	// by the resolver during request preparation.
	Resolved security.ResolvedSecurity
}

// RoundTrip implements http.RoundTripper.
func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Resolved.Disabled || t.Resolved.Scheme.Type == "" {
		return t.base().RoundTrip(req)
	}

	token, err := t.getToken(req.Context())
	if err != nil {
		return nil, fmt.Errorf("httpauth: get token: %w", err)
	}

	injected := req.Clone(req.Context())
	injectCredential(injected, token)

	resp, err := t.base().RoundTrip(injected)
	if err != nil {
		return nil, err
	}
	return t.retryOnUnauthorized(req, resp)
}

func (t *AuthTransport) retryOnUnauthorized(req *http.Request, resp *http.Response) (*http.Response, error) {
	if resp.StatusCode != http.StatusUnauthorized || !t.Resolved.Scheme.RetryOn401 {
		return resp, nil
	}

	DrainAndClose(resp.Body)

	token, err := t.fetchFreshToken(req.Context())
	if err != nil {
		return nil, fmt.Errorf("httpauth: re-auth after 401: %w", err)
	}

	retryReq := req.Clone(req.Context())
	bodyErr := restoreRetryBody(req, retryReq)
	if bodyErr != nil {
		return nil, bodyErr
	}

	injectCredential(retryReq, token)
	return t.base().RoundTrip(retryReq)
}

func restoreRetryBody(original *http.Request, retryReq *http.Request) error {
	if original.Body == nil || original.GetBody == nil {
		return nil
	}

	body, err := original.GetBody()
	if err != nil {
		return fmt.Errorf("httpauth: restore request body for retry: %w", err)
	}
	retryReq.Body = body
	return nil
}

func (t *AuthTransport) getToken(ctx context.Context) (*providers.TokenResult, error) {
	return t.fetchFreshToken(ctx)
}

func (t *AuthTransport) fetchFreshToken(ctx context.Context) (*providers.TokenResult, error) {
	result, err := providers.Authenticate(ctx, t.Resolved.Scheme)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// injectCredential attaches the token to req according to the placement strategy.
func injectCredential(req *http.Request, result *providers.TokenResult) {
	switch result.Placement {
	case config.APIKeyPlacementQuery:
		q := req.URL.Query()
		q.Set(result.QueryParam, result.AccessToken)
		req.URL.RawQuery = q.Encode()
	case config.APIKeyPlacementCookie:
		req.AddCookie(&http.Cookie{
			Name:     result.HeaderName,
			Value:    result.AccessToken,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		})
	case config.APIKeyPlacementHeader, "":
		// header placement (default for bearer, basic, apikey→header)
		headerName := result.HeaderName
		if headerName == "" {
			headerName = "Authorization"
		}
		headerValue := result.HeaderValue
		if headerValue == "" {
			headerValue = result.AccessToken
		}
		req.Header.Set(headerName, headerValue)
	default:
		// Defensive default for any unknown placement.
		req.Header.Set("Authorization", result.AccessToken)
	}
}

func (t *AuthTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

// NewBaseTransport creates an http.Transport that preserves the standard
// library defaults and optionally overrides the proxy URL per scheme. Use this
// as the Base for AuthTransport.
func NewBaseTransport(proxyURL string) (*http.Transport, error) {
	base, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return nil, fmt.Errorf("httpauth: unexpected default transport type %T", http.DefaultTransport)
	}

	tr := base.Clone()
	if proxyURL == "" {
		return tr, nil
	}

	parsed, err := url.Parse(proxyURL)
	if err != nil {
		return nil, fmt.Errorf("httpauth: invalid proxy_url '%s': %w", proxyURL, err)
	}
	tr.Proxy = http.ProxyURL(parsed)
	return tr, nil
}

// NewClient builds an *http.Client with AuthTransport as its transport.
// timeout controls the per-request deadline (0 means no timeout).
func NewClient(
	timeout time.Duration,
	resolved security.ResolvedSecurity,
) (*http.Client, error) {
	proxyURL := ""
	if !resolved.Disabled {
		proxyURL = resolved.Scheme.ProxyURL
	}

	base, err := NewBaseTransport(proxyURL)
	if err != nil {
		return nil, err
	}

	transport := &AuthTransport{
		Base:     base,
		Resolved: resolved,
	}

	// Wrap with DebugTransport when debug mode is enabled in config.
	// DebugTransport sits outside AuthTransport: DebugTransport -> AuthTransport -> Base
	var finalTransport http.RoundTripper = transport
	if cfg, errCfg := config.Get(); errCfg == nil && cfg.IsDebugEnabled() {
		dt := &DebugTransport{
			Base:        transport,
			MaxBodySize: cfg.GetDebugMaxBodySize(formatter.DefaultMaxBodySize),
		}
		finalTransport = dt
	}

	return &http.Client{
		Transport: finalTransport,
		Timeout:   timeout,
	}, nil
}

// DrainAndClose discards any remaining body bytes so the underlying TCP
// connection can be reused.
func DrainAndClose(body io.ReadCloser) {
	if body == nil {
		return
	}
	_, _ = io.Copy(io.Discard, body)
	_ = body.Close()
}
