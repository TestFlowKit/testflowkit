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

			cfg := scenarioCtx.GetConfig()
			client := &http.Client{
				Timeout: time.Duration(cfg.Settings.DefaultTimeout) * time.Millisecond,
			}

			var bodyReader io.Reader
			requestBody := scenarioCtx.GetRequestBody()
			if requestBody != nil {
				bodyReader = bytes.NewReader(requestBody)
			}

			req, prepReqErr := prepareRequest(ctx, scenarioCtx, bodyReader)
			if prepReqErr != nil {
				return ctx, prepReqErr
			}

			startTime := time.Now()
			resp, reqErr := client.Do(req)
			duration := time.Since(startTime)
			if reqErr != nil {
				return ctx, fmt.Errorf("failed to send request: %w", reqErr)
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

func prepareRequest(ctx context.Context, scenarioCtx *scenario.Context, bodyReader io.Reader, ) (*http.Request, error) {

	endpoint := scenarioCtx.GetEndpoint()
	finalURL := endpoint.GetFullURL()
	req, err := http.NewRequestWithContext(ctx, endpoint.Method, finalURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	for key, value := range scenarioCtx.GetRequestHeaders() {
		req.Header.Set(key, value)
	}


	const contentTypeHeader = "Content-Type"
	if req.Header.Get(contentTypeHeader) == "" && scenarioCtx.GetRequestBody() != nil {
		body := scenarioCtx.GetRequestBody()
		contentType := getContentType(body)
		if contentType != "" {
			req.Header.Set(contentTypeHeader, contentType)
		}
	}
	return req, nil
}

func getContentType(body []byte) string {
	if len(body) == 0 {
		return "text/plain"
	}
	if body[0] == '{' || body[0] == '[' {
		return "application/json"
	}

	if strings.HasPrefix(string(body), "<?xml") {
		return "application/xml"
	}

	return "text/plain"
}
