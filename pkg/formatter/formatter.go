// Package formatter provides utilities to pretty-print HTTP request and
// response bodies for debug output. It is intentionally kept in the pkg layer
// so it can be reused by future report-generation features without introducing
// import cycles.
package formatter

import (
	"fmt"
	"mime"
	"strings"
)

const DefaultMaxBodySize int64 = 1 << 20 // 1 MB

// content-type constants used for dispatch.
const (
	ctApplicationJSON = "application/json"
	ctTextJSON        = "text/json"
	ctApplicationXML  = "application/xml"
	ctTextXML         = "text/xml"
)

// binaryPrefixes lists MIME type prefixes and exact values that indicate
// non-text binary content that must not be pretty-printed.
var binaryPrefixes = []string{
	"image/",
	"audio/",
	"video/",
	"application/octet-stream",
	"application/pdf",
	"application/zip",
	"application/gzip",
	"application/x-tar",
	"application/x-bzip2",
}

// IsBinary reports whether contentType represents binary data that should not
// be printed as text.
func IsBinary(contentType string) bool {
	base := baseMediaType(contentType)
	for _, prefix := range binaryPrefixes {
		if strings.HasPrefix(base, prefix) || base == prefix {
			return true
		}
	}
	return false
}

// NeedsFormatting reports whether contentType is a JSON or XML type that
// benefits from indented pretty-printing.
func NeedsFormatting(contentType string) bool {
	base := baseMediaType(contentType)
	switch base {
	case ctApplicationJSON, ctTextJSON, ctApplicationXML, ctTextXML:
		return true
	}
	return false
}

// Format returns a human-readable string representation of body suitable for
// debug log output.
//
//   - If len(body) > maxSize (when maxSize > 0) the body is not formatted and a
//     truncation notice is returned instead.
//   - Binary content types are represented as a "[binary content, N bytes]"
//     placeholder.
//   - JSON and XML bodies are pretty-printed with 2-space indentation. If
//     formatting fails the raw string is returned so that malformed payloads
//     still appear in the log.
//   - All other content types are returned as a raw string.
func Format(contentType string, body []byte, maxSize int64) string {
	if len(body) == 0 {
		return "(empty body)"
	}

	if maxSize > 0 && int64(len(body)) > maxSize {
		return fmt.Sprintf("[body not formatted: %d bytes exceeds limit of %d bytes]", len(body), maxSize)
	}

	if IsBinary(contentType) {
		return fmt.Sprintf("[binary content, %d bytes]", len(body))
	}

	base := baseMediaType(contentType)
	switch base {
	case ctApplicationJSON, ctTextJSON:
		formatted, err := FormatJSON(body)
		if err != nil {
			// Malformed JSON — return raw so the log is still useful.
			return string(body)
		}
		return string(formatted)

	case ctApplicationXML, ctTextXML:
		formatted, err := FormatXML(body)
		if err != nil {
			return string(body)
		}
		return string(formatted)

	default:
		return string(body)
	}
}

// baseMediaType strips MIME parameters (e.g. "; charset=utf-8") and returns
// the lower-cased base media type only.
func baseMediaType(contentType string) string {
	if contentType == "" {
		return ""
	}
	base, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		// mime.ParseMediaType is strict; fall back to trimming manually.
		if idx := strings.IndexByte(contentType, ';'); idx != -1 {
			return strings.ToLower(strings.TrimSpace(contentType[:idx]))
		}
		return strings.ToLower(strings.TrimSpace(contentType))
	}
	return base
}
