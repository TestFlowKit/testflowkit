package restapi

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) debugRequest() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`I debug the current request`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			logger.InfoFf("=== REQUEST DEBUG INFO ===")

			endpoint := scenarioCtx.GetEndpoint()
			if endpoint.Path != "" {
				logger.InfoFf("Endpoint: %s", endpoint.GetFullURL())
				logger.InfoFf("Method: %s", endpoint.Method)
			} else {
				logger.InfoFf("No endpoint configured")
			}

			headers := scenarioCtx.GetRequestHeaders()
			if len(headers) > 0 {
				logger.InfoFf("Headers: %v", headers)
			} else {
				logger.InfoFf("No headers set")
			}

			body := scenarioCtx.GetRESTRequestBody()
			if body != nil {
				logger.InfoFf("Body: %s (%d bytes)", string(body), len(body))
			} else {
				logger.InfoFf("No body set")
			}

			logger.InfoFf("=== END DEBUG INFO ===")

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Debug helper to show the current request configuration.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When I debug the current request`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.RESTAPI}},
	)
}
