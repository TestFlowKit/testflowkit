package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/api/validation"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

// validateJSONBodyContains checks if response body contains a specific value.
func (s steps) validateJSONBodyContains() stepbuilder.Step {
	return s.newJSONBodyContainsStep(
		`the response should contain "{string}"`,
		true,
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

// validateJSONBodyNotContains checks if response body does not contain a specific value.
func (s steps) validateJSONBodyNotContains() stepbuilder.Step {
	return s.newJSONBodyContainsStep(
		`the response should not contain "{string}"`,
		false,
		stepbuilder.DocParams{
			Description: "Validates that the response does not contain a specific text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "text", Description: "Text that should not be present in the response", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response should not contain "error"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}

func (s steps) newJSONBodyContainsStep(
	sentence string, shouldContain bool, doc stepbuilder.DocParams,
) stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{sentence},
		func(ctx context.Context, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			text = scenario.ReplaceVariablesInString(scenarioCtx, text)

			responseBody := backend.GetResponseBody()
			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			err := validation.ValidateBodyContains(responseBody, text)
			if shouldContain {
				return ctx, err
			}

			if err == nil {
				return ctx, fmt.Errorf("response body contains forbidden substring '%s'", text)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
