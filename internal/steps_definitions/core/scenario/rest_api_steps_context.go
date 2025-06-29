package scenario

import (
	"fmt"
	"net/url"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

type RESTAPIContext struct {
	Endpoint       *EndpointEnricher
	RequestHeaders map[string]string
	RequestBody    []byte
	Response       *HTTPResponse
}

type EndpointEnricher struct {
	config.Endpoint
	QueryParams       map[string]string
	PathParams        map[string]string
	ConfiguredBaseURL string
}

func (e *EndpointEnricher) GetFullURL() string {
	fullURL, err := e.getSimpleURL()
	if err != nil {
		return ""
	}

	if len(e.PathParams) > 0 {
		for name, value := range e.PathParams {
			fullURL = strings.ReplaceAll(fullURL, fmt.Sprintf("{%s}", name), value)
		}
	}

	if len(e.QueryParams) > 0 {
		params := make([]string, 0, len(e.QueryParams))
		for key, value := range e.QueryParams {
			params = append(params, fmt.Sprintf("%s=%s", key, value))
		}
		fullURL += "?" + strings.Join(params, "&")
	}

	return fullURL
}

func (e *EndpointEnricher) getSimpleURL() (string, error) {
	parsedURL, err := url.Parse(e.Endpoint.Path)
	if err != nil {
		return "", err
	}

	if parsedURL.Scheme != "" {
		logger.InfoFf("Parsed URL: %s", parsedURL.String())
		return parsedURL.String(), nil
	}

	fullURL, err := url.JoinPath(e.ConfiguredBaseURL, parsedURL.Path)
	logger.InfoFf("Full URL joined: %s", fullURL)
	if err != nil {
		return "", err
	}

	unescapedURL, _ := url.PathUnescape(fullURL)
	return unescapedURL, nil
}
