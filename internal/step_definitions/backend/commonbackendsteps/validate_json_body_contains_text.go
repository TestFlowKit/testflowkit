package commonbackendsteps

import (
	"context"
	"errors"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONBodyContains checks if response body contains a specific value.
func (steps) validateJSONBodyContains() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the response should contain "{string}"`},
		func(ctx context.Context, expectedText string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to validate")
			}

			expectedText = scenario.ReplaceVariablesInString(scenarioCtx, expectedText)

			// Get the appropriate response body based on protocol
			responseBody := backend.GetResponseBody()
			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			err := validation.ValidateBodyContains(responseBody, expectedText)
			return ctx, err
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the response contains a specific text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "text", Description: "Text that should be present in the response", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response should contain "success"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}
