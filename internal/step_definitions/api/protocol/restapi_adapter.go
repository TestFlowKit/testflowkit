package protocol

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/pkg/logger"
	"time"
)

type RESTAPIAdapter struct{}

func NewRESTAPIAdapter() *RESTAPIAdapter {
	return &RESTAPIAdapter{}
}

func (a *RESTAPIAdapter) PrepareRequest(ctx context.Context, endpointName string) (context.Context, error) {
	scenarioCtx := scenario.MustFromContext(ctx)
	cfg := scenarioCtx.GetConfig()

	endpoint, exists := cfg.Backend.Endpoints[endpointName]
	if !exists {
		return ctx, fmt.Errorf("endpoint '%s' not found in configuration", endpointName)
	}

	baseURL := cfg.GetBackendBaseURL()
	scenarioCtx.SetEndpoint(baseURL, endpoint)

	// Store this adapter as the protocol
	scenarioCtx.GetBackendContext().SetProtocol(a)

	return ctx, nil
}

func (a *RESTAPIAdapter) SendRequest(ctx context.Context) (context.Context, error) {
	scenarioCtx := scenario.MustFromContext(ctx)
	const defaultDuration = 2

	client := &http.Client{
		// TODO: make timeout configurable
		Timeout: time.Duration(defaultDuration) * time.Second,
	}

	req, err := a.createRequest(ctx, scenarioCtx)
	if err != nil {
		return ctx, err
	}

	startTime := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(startTime)
	if err != nil {
		return ctx, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx, fmt.Errorf("failed to read response body: %w", err)
	}

	// Store response in unified format
	scenarioCtx.SetResponse(resp.StatusCode, responseBody)

	logger.InfoFf("REST request completed - Status: %d, Duration: %v, Response size: %d bytes",
		resp.StatusCode, duration, len(responseBody))

	return ctx, nil
}

func (a *RESTAPIAdapter) createRequest(ctx context.Context, scenarioCtx *scenario.Context) (*http.Request, error) {
	endpoint := scenarioCtx.GetEndpoint()
	bodyReader := a.getBodyReader(scenarioCtx)

	finalURL := endpoint.GetFullURL()
	logger.InfoFf("Sending %s request to: %s", endpoint.Method, finalURL)

	req, err := http.NewRequestWithContext(ctx, endpoint.Method, finalURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	for key, value := range scenarioCtx.GetRequestHeaders() {
		req.Header.Set(key, value)
	}

	// Auto-detect and set Content-Type if not already set
	if req.Header.Get("Content-Type") == "" && scenarioCtx.GetRequestBody() != nil {
		body := scenarioCtx.GetRequestBody()
		contentType := a.detectContentType(body)
		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		}
	}

	return req, nil
}

func (a *RESTAPIAdapter) getBodyReader(scenarioCtx *scenario.Context) io.Reader {
	requestBody := scenarioCtx.GetRequestBody()
	if requestBody != nil {
		logger.InfoFf("Request body found: %d bytes", len(requestBody))
		return bytes.NewReader(requestBody)
	}
	logger.InfoFf("No request body set")
	return nil
}

func (a *RESTAPIAdapter) detectContentType(body []byte) string {
	if len(body) == 0 {
		return ""
	}

	// Check for JSON
	if body[0] == '{' || body[0] == '[' {
		return "application/json"
	}

	// Check for XML
	if strings.HasPrefix(string(body), "<?xml") {
		return "application/xml"
	}

	return "text/plain"
}

func (a *RESTAPIAdapter) GetResponseBody(ctx context.Context) ([]byte, error) {
	scenarioCtx := scenario.MustFromContext(ctx)
	backend := scenarioCtx.GetBackendContext()

	if !backend.HasResponse() {
		return nil, errors.New("no REST API response available")
	}

	return backend.GetResponseBody(), nil
}

func (a *RESTAPIAdapter) GetStatusCode(ctx context.Context) (int, error) {
	scenarioCtx := scenario.MustFromContext(ctx)
	backend := scenarioCtx.GetBackendContext()

	if !backend.HasResponse() {
		return 0, errors.New("no REST API response available")
	}

	return backend.GetStatusCode(), nil
}

func (a *RESTAPIAdapter) HasErrors(ctx context.Context) bool {
	scenarioCtx := scenario.MustFromContext(ctx)
	backend := scenarioCtx.GetBackendContext()

	if !backend.HasResponse() {
		return false
	}

	// REST API considers 4xx and 5xx status codes as errors
	const minErrorStatus = 400
	return backend.GetStatusCode() >= minErrorStatus
}

func (a *RESTAPIAdapter) GetProtocolName() string {
	return string(ProtocolRESTAPI)
}
