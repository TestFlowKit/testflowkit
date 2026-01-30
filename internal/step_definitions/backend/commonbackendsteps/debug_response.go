package commonbackendsteps

import (
	"context"
	"errors"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) debugAPIResponse() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`I debug the API response`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to debug")
			}

			logger.InfoFf("=== API RESPONSE DEBUG INFO ===")

			response := backend.GetResponse()

			statusCode := backend.GetStatusCode()
			logger.InfoFf("Status Code: %d", statusCode)

			// Headers
			if len(response.Headers) > 0 {
				logger.InfoFf("Headers:")
				for key, value := range response.Headers {
					logger.InfoFf("  %s: %s", key, value)
				}
			} else {
				logger.InfoFf("No headers in response")
			}

			// Body
			body := backend.GetResponseBody()
			if len(body) > 0 {
				logger.InfoFf("Body: %s (%d bytes)", string(body), len(body))
			} else {
				logger.InfoFf("No body in response")
			}

			logger.InfoFf("=== END DEBUG INFO ===")

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Debug helper to show the current API response details including status code, headers, and body.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `Then I debug the API response`,
			Categories:  stepbuilder.Backend,
		},
	)
}
