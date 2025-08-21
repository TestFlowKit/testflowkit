package restapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) sendRequest() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`I send the request`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			const defaultDuration = 2

			client := &http.Client{
				// TODO: add timeout configureable
				// Timeout: time.Duration(cfg.Frontend.DefaultTimeout) * time.Millisecond,
				Timeout: time.Duration(defaultDuration) * time.Second,
			}

			req, err := createRequest(ctx)
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

			scenarioCtx.SetResponse(resp.StatusCode, responseBody)

			logger.InfoFf("Request completed - Status: %d, Duration: %v, Response size: %d bytes",
				resp.StatusCode, duration, len(responseBody))

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sends the prepared HTTP request and stores the response.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When I send the request`,
			Category:    stepbuilder.RESTAPI,
		},
	)
}

func createRequest(ctx context.Context) (*http.Request, error) {
	scenarioCtx := scenario.MustFromContext(ctx)

	endpoint := scenarioCtx.GetEndpoint()
	bodyReader := getBodyReader(scenarioCtx)

	finalURL := endpoint.GetFullURL()
	logger.InfoFf("Sending %s request to: %s", endpoint.Method, finalURL)

	req, err := http.NewRequestWithContext(ctx, endpoint.Method, finalURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	for key, value := range scenarioCtx.GetRequestHeaders() {
		req.Header.Set(key, value)
	}

	if req.Header.Get("Content-Type") == "" && scenarioCtx.GetRequestBody() != nil {
		body := scenarioCtx.GetRequestBody()
		contentType := getContentType(body)
		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		}
	}
	return req, nil
}

func getBodyReader(scenarioCtx *scenario.Context) io.Reader {
	var bodyReader io.Reader
	requestBody := scenarioCtx.GetRequestBody()
	if requestBody != nil {
		bodyReader = bytes.NewReader(requestBody)
		logger.InfoFf("Request body found: %d bytes", len(requestBody))
	} else {
		logger.InfoFf("No request body set")
	}
	return bodyReader
}

func getContentType(body []byte) string {
	if body[0] == '{' || body[0] == '[' {
		return "application/json"
	}

	if strings.HasPrefix(string(body), "<?xml") {
		return "application/xml"
	}

	return "text/plain"
}
