package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

// validateJSONPathContains validates that a JSON path field contains specific text.
func (steps) validateJSONPathContains() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the response field "{string}" should contain "{string}"`},
		func(ctx context.Context, jsonPath, expectedText string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to validate")
			}

			jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)
			expectedText = scenario.ReplaceVariablesInString(scenarioCtx, expectedText)

			// Get the appropriate response body based on protocol

			responseBody := backend.GetResponseBody()

			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			// Extract the value at the JSON path
			actualValue, err := jsonhelpers.GetPathValueAsString(responseBody, jsonPath)
			if err != nil {
				return ctx, fmt.Errorf("failed to get value at path '%s': %w", jsonPath, err)
			}

			// Check if the value contains the expected text
			if !strings.Contains(actualValue, expectedText) {
				msg := "field '%s' value '%s' does not contain expected text '%s'"
				return ctx, fmt.Errorf(msg, jsonPath, actualValue, expectedText)
			}

			return ctx, nil
		},
		nil,
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
