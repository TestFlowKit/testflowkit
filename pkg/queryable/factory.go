package queryable

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// Format identifies the response body format used by a query engine.
type Format string

const (
	FormatAuto Format = "auto"
	FormatJSON Format = "json"
	FormatXML  Format = "xml"
)

const (
	xPathOnJSONMessage = "XPath syntax detected on JSON content; please use GJSON syntax"
	gjsonOnXMLMessage  = "JSONPath/GJSON syntax detected on XML content; please use XPath"
)

// NewEngine creates a query engine for the provided response body and format.
func NewEngine(body []byte, format Format) (Queryable, error) {
	trimmedBody := bytes.TrimSpace(body)
	if len(trimmedBody) == 0 {
		return nil, errors.New("response body is empty")
	}

	switch normalizeFormat(format) {
	case FormatXML:
		return newXMLEngine(trimmedBody)
	case FormatJSON:
		return newJSONEngine(trimmedBody), nil
	case FormatAuto:
		if trimmedBody[0] == '<' {
			return newXMLEngine(trimmedBody)
		}
		return newJSONEngine(trimmedBody), nil
	default:
		return nil, fmt.Errorf("unsupported response format %q", format)
	}
}

// DetectFormatFromContentType maps an HTTP Content-Type header to a query format.
func DetectFormatFromContentType(contentType string) Format {
	normalizedContentType := strings.ToLower(strings.TrimSpace(contentType))
	if normalizedContentType == "" {
		return FormatAuto
	}

	switch {
	case strings.Contains(normalizedContentType, "xml"):
		return FormatXML
	case strings.Contains(normalizedContentType, "json"):
		return FormatJSON
	default:
		return FormatAuto
	}
}

func normalizeFormat(format Format) Format {
	normalizedFormat := strings.ToLower(strings.TrimSpace(string(format)))
	if normalizedFormat == "" {
		return FormatAuto
	}

	return Format(normalizedFormat)
}

func validatePathSyntax(path string, format Format) error {
	trimmedPath := strings.TrimSpace(path)
	if trimmedPath == "" {
		return errors.New("path cannot be empty")
	}

	switch format {
	case FormatXML:
		if strings.Contains(trimmedPath, "#(") || strings.Contains(trimmedPath, ".#") {
			return errors.New(gjsonOnXMLMessage)
		}
	case FormatJSON:
		if strings.HasPrefix(trimmedPath, "/") {
			return errors.New(xPathOnJSONMessage)
		}
	case FormatAuto:
		return nil
	}

	return nil
}
