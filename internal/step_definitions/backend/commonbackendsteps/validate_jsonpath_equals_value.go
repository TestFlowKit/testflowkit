package commonbackendsteps

import (
	"context"
	"errors"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONPathValue validates that a JSON path has a specific value.
func (steps) validateJSONPathValue() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the response field "{string}" should be "{string}"`},
		func(ctx context.Context, jsonPath, expectedValue string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to validate")
			}

			jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)
			expectedValue = scenario.ReplaceVariablesInString(scenarioCtx, expectedValue)

			responseBody := backend.GetResponseBody()

			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			if err := validation.ValidateJSONPathValue(responseBody, jsonPath, expectedValue); err != nil {
				return ctx, err
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path has the expected value.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to validate", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "Expected value at the path", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response field "user.name" should be "John Doe"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}
