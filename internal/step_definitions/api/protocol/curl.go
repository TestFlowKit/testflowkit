package protocol

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"testflowkit/internal/config"
	"testflowkit/internal/security"
	"testflowkit/internal/security/providers"
)

func buildCurlCommand(method string, requestURL string, headers map[string]string, body []byte) string {
	parts := []string{"curl", "-X", method}

	for _, header := range formatHeaders(headers) {
		parts = append(parts, "-H", quoteForShell(header))
	}

	if len(body) > 0 {
		parts = append(parts, "--data-raw", quoteForShell(string(body)))
	}

	parts = append(parts, quoteForShell(requestURL))
	return strings.Join(parts, " ")
}

func formatHeaders(headers map[string]string) []string {
	if len(headers) == 0 {
		return nil
	}

	keys := make([]string, 0, len(headers))
	for key := range headers {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	formatted := make([]string, 0, len(keys))
	for _, key := range keys {
		formatted = append(formatted, key+": "+headers[key])
	}

	return formatted
}

func withStaticSecurity(
	headers map[string]string,
	requestURL string,
	resolved security.ResolvedSecurity,
) (map[string]string, string) {
	outHeaders := cloneHeaders(headers)
	if resolved.Disabled || resolved.Scheme.Type == "" {
		return outHeaders, requestURL
	}

	scheme := resolved.Scheme
	switch scheme.Type {
	case config.SecurityTypeBearer:
		outHeaders["Authorization"] = "Bearer " + scheme.Token
	case config.SecurityTypeBasic:
		creds := scheme.Username + ":" + scheme.Password
		outHeaders["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(creds))
	case config.SecurityTypeAPIKey:
		return applyAPIKeySecurity(outHeaders, requestURL, scheme)
	case config.SecurityTypeOAuth2,
		config.SecurityTypeOIDC,
		config.SecurityTypeCertificate,
		config.SecurityTypeNone:
		return outHeaders, requestURL
	}

	return outHeaders, requestURL
}

func applyAPIKeySecurity(
	headers map[string]string,
	requestURL string,
	scheme config.SecurityScheme,
) (map[string]string, string) {
	placement := scheme.Placement
	if placement == "" {
		placement = config.APIKeyPlacementHeader
	}

	switch placement {
	case config.APIKeyPlacementHeader:
		headerName := scheme.HeaderName
		if headerName == "" {
			headerName = "Authorization"
		}
		headers[headerName] = scheme.Key
		return headers, requestURL
	case config.APIKeyPlacementQuery:
		parsedURL, err := url.Parse(requestURL)
		if err != nil {
			return headers, requestURL
		}

		queryParam := scheme.QueryParam
		if queryParam == "" {
			queryParam = providers.DefaultAPIKeyQueryParam
		}

		query := parsedURL.Query()
		query.Set(queryParam, scheme.Key)
		parsedURL.RawQuery = query.Encode()
		return headers, parsedURL.String()
	case config.APIKeyPlacementCookie:
		cookieName := scheme.HeaderName
		if cookieName == "" {
			cookieName = "Authorization"
		}

		cookieValue := cookieName + "=" + scheme.Key
		if existingCookie := headers["Cookie"]; existingCookie != "" {
			headers["Cookie"] = existingCookie + "; " + cookieValue
		} else {
			headers["Cookie"] = cookieValue
		}
		return headers, requestURL
	default:
		return headers, requestURL
	}
}

func cloneHeaders(headers map[string]string) map[string]string {
	if len(headers) == 0 {
		return make(map[string]string)
	}

	cloned := make(map[string]string, len(headers))
	for key, value := range headers {
		cloned[key] = value
	}
	return cloned
}

func quoteForShell(value string) string {
	return strconv.Quote(value)
}

func defaultRESTMethod(method string) string {
	if method == "" {
		return http.MethodGet
	}
	return method
}
