package scenario

import (
	"errors"
	"fmt"
	"maps"
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
	raw, err := e.getSimpleURL()
	if err != nil {
		return ""
	}

	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}

	for name, value := range e.PathParams {
		u.Path = strings.ReplaceAll(
			u.Path,
			fmt.Sprintf("{%s}", name),
			value,
		)
	}

	q := u.Query()
	for key, value := range e.QueryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	return strings.TrimRight(u.String(), "/")
}

func (e *EndpointEnricher) getSimpleURL() (string, error) {
	if strings.TrimSpace(e.Path) == "" {
		return "", errors.New("endpoint path is empty")
	}

	rawPath := strings.TrimSpace(e.Path)

	// More robust than strings.Contains(path, "://")
	if u, err := url.Parse(rawPath); err == nil && u.IsAbs() {
		return u.String(), nil
	}

	if e.ConfiguredBaseURL == "" {
		return rawPath, nil
	}

	base, err := url.Parse(e.ConfiguredBaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL %q: %w", e.ConfiguredBaseURL, err)
	}

	base.Path = strings.TrimSuffix(base.Path, "/") + "/" + strings.TrimPrefix(rawPath, "/")
	return base.String(), nil
}

func (e *EndpointEnricher) SetPathParams(params map[string]string) {
	if e.PathParams == nil {
		e.PathParams = make(map[string]string)
	}
	maps.Copy(e.PathParams, params)
}

func (e *EndpointEnricher) SetQueryParams(params map[string]string) {
	if e.QueryParams == nil {
		e.QueryParams = make(map[string]string)
	}
	maps.Copy(e.QueryParams, params)
}
