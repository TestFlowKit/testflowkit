package commonbackendsteps

import (
	"context"
	"errors"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONBodyEquals validates that the entire response body matches expected JSON.
func (steps) validateJSONBodyEquals() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response body should be:`},
		func(ctx context.Context, expectedJSON string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to validate")
			}

			expectedJSON = scenario.ReplaceVariablesInString(scenarioCtx, expectedJSON)

			responseBody := backend.GetResponseBody()
			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			err := validation.ValidateJSONBodyEquals(responseBody, expectedJSON)
			if err != nil {
				return ctx, err
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the response body matches the expected JSON exactly.",
			Variables: []stepbuilder.DocVariable{
				{Name: "json", Description: "Expected JSON body", Type: stepbuilder.VarTypeString},
			},
			Example: `Then the response body should be:
"""
{"status":"success","data":{"id":1}}
"""`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}
