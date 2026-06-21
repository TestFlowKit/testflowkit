package commonbackendsteps

import (
	"context"
	"slices"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/formatter"
	"testflowkit/pkg/logger"
)

func (steps) debugAPIResponse() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`I debug the API response`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			logger.Debug("=== API RESPONSE DEBUG INFO ===")

			response := backend.GetResponse()

			statusCode := backend.GetStatusCode()
			logger.Debugf("Status Code: %d", statusCode)

			// Headers
			if len(response.Headers) > 0 {
				logger.Debugf("Headers:")
				for key, value := range response.Headers {
					logger.Debugf("  %s: %s", key, value)
				}
			} else {
				logger.Debugf("No headers in response")
			}

			// Body
			body := backend.GetResponseBody()
			if len(body) > 0 {
				contentType := response.Headers["Content-Type"]
				formatted := formatter.Format(contentType, body, formatter.DefaultMaxBodySize)
				logger.Debugf("Body (%d bytes):", len(body))
				logger.Debug(formatted)
			} else {
				logger.Debug("No body in response")
			}

			logger.Debug("=== END DEBUG INFO ===")

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Debug helper to show the current API response details including status code, headers, and body.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `Then I debug the API response`,
			Categories:  slices.Concat([]stepbuilder.StepCategory{stepbuilder.Debug}, stepbuilder.Backend),
		},
	)
}
