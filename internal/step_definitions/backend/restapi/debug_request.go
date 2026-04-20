package restapi

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/formatter"
	"testflowkit/pkg/logger"
)

func (steps) debugRequest() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"I debug the API request", "I debug the current request"},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			logger.Debug("=== REQUEST DEBUG INFO ===")

			endpoint := scenarioCtx.GetEndpoint()
			if endpoint.Path != "" {
				logger.DebugFf("Endpoint: %s", endpoint.GetFullURL())
				logger.DebugFf("Method: %s", endpoint.Method)
			} else {
				logger.DebugFf("No endpoint configured")
			}

			headers := scenarioCtx.GetRequestHeaders()
			if len(headers) > 0 {
				logger.DebugFf("Headers: %v", headers)
			} else {
				logger.DebugFf("No headers set")
			}

			body := scenarioCtx.GetRESTRequestBody()
			if body != nil {
				contentType := scenarioCtx.GetRequestHeaders()["Content-Type"]
				formatted := formatter.Format(contentType, body, formatter.DefaultMaxBodySize)
				logger.DebugFf("Body (%d bytes):", len(body))
				logger.Debug(formatted)
			} else {
				logger.Debug("No body set")
			}

			logger.Debug("=== END DEBUG INFO ===")

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Debug helper to show the current request configuration.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When I debug the API request`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.RESTAPI}},
	)
}
