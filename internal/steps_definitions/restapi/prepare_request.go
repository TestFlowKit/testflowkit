package restapi

import (
	"context"
	"fmt"

	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) prepareRequest() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I prepare a request for the {string} endpoint`},
		func(ctx context.Context, endpointName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			cfg := scenarioCtx.GetConfig()

			_, endpoint, err := cfg.GetAPIEndpoint(endpointName)
			if err != nil {
				return ctx, fmt.Errorf("failed to resolve endpoint '%s': %w", endpointName, err)
			}

			scenarioCtx.SetEndpoint(cfg.GetBackendBaseURL(), endpoint)

			logger.InfoFf("Request prepared for endpoint '%s' (%s)", endpointName, endpoint.Description)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Prepares an HTTP request for a configured endpoint with automatic method and URL resolution.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "endpointName",
					Description: "The logical endpoint name as defined in the configuration (e.g., 'get_user', 'create_product')",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Given I prepare a request for the "get_user" endpoint`,
			Category: stepbuilder.RESTAPI,
		},
	)
}
