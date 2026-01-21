package commonbackendsteps

import (
	"context"
	"errors"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONPathExists validates that a JSON path exists in the response.
func (steps) validateJSONPathExists() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response should have field {string}`},
		func(ctx context.Context, jsonPath string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to validate")
			}

			jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)

			responseBody := backend.GetResponseBody()

			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			if err := validation.ValidateJSONPathExists(responseBody, jsonPath); err != nil {
				return ctx, err
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path exists in the response.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to check", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response should have field "user.name"`,
			Categories: stepbuilder.Backend,
		},
	)
}
