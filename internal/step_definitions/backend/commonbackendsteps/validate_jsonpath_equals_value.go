package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

// validateJSONPathValue validates that a JSON path has a specific value.
func (s steps) validateJSONPathValue() stepbuilder.Step {
	return s.newJSONPathValueStep(
		`the response field "{string}" should be "{string}"`,
		true,
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

// validateJSONPathValueNot validates that a JSON path does not have a specific value.
func (s steps) validateJSONPathValueNot() stepbuilder.Step {
	return s.newJSONPathValueStep(
		`the response field "{string}" should not be "{string}"`,
		false,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path does not have the specified value.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to validate", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "Value that should not be at the path", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response field "user.role" should not be "admin"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}

func (s steps) newJSONPathValueStep(sentence string, shouldEqual bool, doc stepbuilder.DocParams) stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, jsonPath, value string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)
			value = scenario.ReplaceVariablesInString(scenarioCtx, value)

			responseBody := backend.GetResponseBody()

			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			actualValue, err := jsonhelpers.GetPathValueAsString(responseBody, jsonPath)
			if err != nil {
				return ctx, fmt.Errorf("failed to get value at path '%s': %w", jsonPath, err)
			}

			if shouldEqual && actualValue != value {
				return ctx, fmt.Errorf("expected value '%s' at path '%s', but got '%s'", value, jsonPath, actualValue)
			}

			if !shouldEqual && actualValue == value {
				return ctx, fmt.Errorf("field '%s' has forbidden value '%s'", jsonPath, actualValue)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
