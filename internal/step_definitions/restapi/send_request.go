package restapi

import (
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
			endpoint := scenarioCtx.GetEndpoint()
			logger.InfoFf("Sending %s request to: %s", endpoint.Method, endpoint.GetFullURL())
			cfg := scenarioCtx.GetConfig()
			client := &http.Client{
				Timeout: time.Duration(cfg.Settings.DefaultTimeout) * time.Millisecond,
			}

			finalURL := endpoint.GetFullURL()

			logger.InfoFf("Final URL: %s", finalURL)

			var bodyReader io.Reader
			if scenarioCtx.GetRequestBody() != nil {
				bodyReader = strings.NewReader(string(scenarioCtx.GetRequestBody()))
			}

			req, err := http.NewRequestWithContext(ctx, endpoint.Method, finalURL, bodyReader)
			if err != nil {
				return ctx, fmt.Errorf("failed to create HTTP request: %w", err)
			}

			for key, value := range scenarioCtx.GetRequestHeaders() {
				req.Header.Set(key, value)
			}

			startTime := time.Now()
			logger.InfoFf("Sending request to: %s", req.URL.String())
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
