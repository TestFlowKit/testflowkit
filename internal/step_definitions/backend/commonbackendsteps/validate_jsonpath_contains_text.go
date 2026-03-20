package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

// validateJSONPathContains validates that a JSON path field contains specific text.
func (s steps) validateJSONPathContains() stepbuilder.Step {
	return s.newJSONPathContainsStep(
		`the response field "{string}" should contain "{string}"`,
		true,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path field contains the expected text (substring match).",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to the field to validate", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "Text that should be contained in the field value", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response field "user.email" should contain "@example.com"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}

// validateJSONPathNotContains validates that a JSON path field does not contain specific text.
func (s steps) validateJSONPathNotContains() stepbuilder.Step {
	return s.newJSONPathContainsStep(
		`the response field "{string}" should not contain "{string}"`,
		false,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path field does not contain the specified text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to the field to validate", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "Text that should not be present in the field value", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response field "user.email" should not contain "@internal"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}

func (s steps) newJSONPathContainsStep(
	sentence string, shouldContain bool, doc stepbuilder.DocParams,
) stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, jsonPath, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)
			text = scenario.ReplaceVariablesInString(scenarioCtx, text)

			responseBody := backend.GetResponseBody()

			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			actualValue, err := jsonhelpers.GetPathValueAsString(responseBody, jsonPath)
			if err != nil {
				return ctx, fmt.Errorf("failed to get value at path '%s': %w", jsonPath, err)
			}

			contains := strings.Contains(actualValue, text)
			if shouldContain && !contains {
				msg := "field '%s' value '%s' does not contain expected text '%s'"
				return ctx, fmt.Errorf(msg, jsonPath, actualValue, text)
			}

			if !shouldContain && contains {
				msg := "field '%s' value '%s' should not contain text '%s'"
				return ctx, fmt.Errorf(msg, jsonPath, actualValue, text)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
