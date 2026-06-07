package logger

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

var sensitiveKeys = []string{
	"password",
	"pass",
	"pwd",
	"token",
	"access_token",
	"refresh_token",
	"secret",
	"api_key",
	"apikey",
	"authorization",
	"auth",
	"ssn",
	"cookie",
}

func isSensitiveKey(k string) bool {
	lk := strings.ToLower(k)
	for _, s := range sensitiveKeys {
		if lk == s || strings.HasSuffix(lk, "_"+s) || strings.HasPrefix(lk, s+"_") || strings.Contains(lk, s) {
			return true
		}
	}
	return false
}

// IsSensitiveKey reports whether a variable/header/query key looks sensitive.
func IsSensitiveKey(k string) bool {
	return isSensitiveKey(k)
}

// MaskHeaders returns a shallow copy of headers with sensitive header values redacted.
func MaskHeaders(h http.Header) http.Header {
	out := http.Header{}
	for k, vals := range h {
		if isSensitiveKey(k) || strings.EqualFold(k, "Authorization") ||
			strings.EqualFold(k, "Cookie") || strings.EqualFold(k, "Set-Cookie") {
			out[k] = []string{"[REDACTED]"}
			continue
		}
		// copy values
		out[k] = append([]string(nil), vals...)
	}
	return out
}

// MaskURL returns a string representation of the URL with sensitive query params redacted.
func MaskURL(u *url.URL) string {
	if u == nil {
		return ""
	}
	q := u.Query()
	for k := range q {
		if isSensitiveKey(k) {
			q.Set(k, "[REDACTED]")
		}
	}
	masked := *u
	masked.RawQuery = q.Encode()
	return masked.String()
}

// MaskBody attempts to redact sensitive JSON fields. If the body is not JSON
// or parsing fails, the original body is returned unchanged.
func MaskBody(contentType string, body []byte) []byte {
	base := strings.ToLower(contentType)
	if !strings.Contains(base, "json") {
		return body
	}

	var v interface{}
	if err := json.Unmarshal(body, &v); err != nil {
		return body
	}

	maskRecursive(v)
	out, err := json.Marshal(v)
	if err != nil {
		return body
	}
	return out
}

func maskRecursive(v any) {
	switch t := v.(type) {
	case map[string]any:
		for k, val := range t {
			if isSensitiveKey(k) {
				t[k] = "[REDACTED]"
				continue
			}
			maskRecursive(val)
		}
	case []any:
		for i := range t {
			maskRecursive(t[i])
		}
	}
}

// HeadersToString returns a deterministic, sorted string representation of headers.
func HeadersToString(h http.Header) string {
	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		b.WriteString(k)
		b.WriteString(": ")
		b.WriteString(strings.Join(h[k], ", "))
		b.WriteString("\n")
	}
	return b.String()
}
