package scenario

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"testflowkit/internal/config"
)

func (bc *BackendContext) SetEndpoint(endpoint *EndpointEnricher) {
	bc.Endpoint = endpoint
}

func (bc *BackendContext) GetEndpoint() *EndpointEnricher {
	return bc.Endpoint
}

// SetRequestBody sets the REST request body.
func (bc *BackendContext) SetRequestBody(body []byte) {
	bc.RequestBody = body
}

func (bc *BackendContext) GetRequestBody() []byte {
	return bc.RequestBody
}

func (bc *BackendContext) AddPathParam(param, value string) {
	if bc.Endpoint == nil {
		bc.Endpoint = &EndpointEnricher{
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		}
	}
	bc.Endpoint.PathParams[param] = value
}

func (bc *BackendContext) AddQueryParam(key, value string) {
	if bc.Endpoint == nil {
		bc.Endpoint = &EndpointEnricher{
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		}
	}
	bc.Endpoint.QueryParams[key] = value
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
			placeholder := fmt.Sprintf("{%s}", name)
			fullURL = strings.ReplaceAll(fullURL, placeholder, value)
		}
	}

	if len(e.QueryParams) > 0 {
		values := url.Values{}
		for key, value := range e.QueryParams {
			values.Set(key, value)
		}
		encoded := values.Encode()
		if encoded != "" {
			fullURL += "?" + encoded
		}
	}

	return fullURL
}

func (e *EndpointEnricher) getSimpleURL() (string, error) {
	path := e.Path
	if path == "" {
		return "", errors.New("endpoint path is empty")
	}

	// If it's already an absolute URL, return as-is
	if strings.Contains(path, "://") {
		return path, nil
	}

	// If no base URL, return path as-is
	if e.ConfiguredBaseURL == "" {
		return path, nil
	}

	// Join base URL and path, handling slashes
	baseURL := strings.TrimSuffix(e.ConfiguredBaseURL, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return baseURL + path, nil
}
