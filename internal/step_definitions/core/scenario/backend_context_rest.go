package scenario

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"testflowkit/internal/config"
)

func (bc *BackendContext) SetEndpoint(endpoint *EndpointEnricher) {
	bc.Rest.Endpoint = endpoint
}

func (bc *BackendContext) GetEndpoint() *EndpointEnricher {
	return bc.Rest.Endpoint
}

func (bc *BackendContext) SetRESTRequestBody(body []byte) {
	bc.Rest.RequestBody = body
}

func (bc *BackendContext) GetRESTRequestBody() []byte {
	return bc.Rest.RequestBody
}

func (bc *BackendContext) AddRESTPathParam(param, value string) {
	if bc.Rest.Endpoint == nil {
		bc.Rest.Endpoint = &EndpointEnricher{
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		}
	}
	bc.Rest.Endpoint.PathParams[param] = value
}

func (bc *BackendContext) AddRESTQueryParam(key, value string) {
	if bc.Rest.Endpoint == nil {
		bc.Rest.Endpoint = &EndpointEnricher{
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		}
	}
	bc.Rest.Endpoint.QueryParams[key] = value
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
