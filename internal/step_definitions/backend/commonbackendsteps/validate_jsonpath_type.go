package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/helpers"
	"testflowkit/pkg/apperrors"
)

// validateJSONPathType validates that a JSON path field has a specific type.
func (s steps) validateJSONPathType() stepbuilder.Step {
	return s.newJSONPathTypeStep(
		`the response field {string} should have type {string}`,
		true,
		func() stepbuilder.DocParams {
			tdsc := "Expected type (string, number, integer, boolean, object, array, null)"
			vars := []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to the field to validate", Type: stepbuilder.VarTypeString},
				{Name: "type", Description: tdsc, Type: stepbuilder.VarTypeString},
			}
			return stepbuilder.DocParams{
				Description: "Validates that a specific JSON path field has the expected type.",
				Variables:   vars,
				Example:     `Then the response field "user.age" should have type "integer"`,
				Categories:  stepbuilder.Backend,
			}
		}(),
	)
}

// validateJSONPathNotType validates that a JSON path field does not have a specific type.
func (s steps) validateJSONPathNotType() stepbuilder.Step {
	return s.newJSONPathTypeStep(
		`the response field {string} should not have type {string}`,
		false,
		func() stepbuilder.DocParams {
			tdsc := "Type that must not match (string, number, integer, boolean, object, array, null)"
			vars := []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to the field to validate", Type: stepbuilder.VarTypeString},
				{Name: "type", Description: tdsc, Type: stepbuilder.VarTypeString},
			}
			return stepbuilder.DocParams{
				Description: "Validates that a specific JSON path field does not have the specified type.",
				Variables:   vars,
				Example:     `Then the response field "user.id" should not have type "string"`,
				Categories:  stepbuilder.Backend,
			}
		}(),
	)
}

func (s steps) newJSONPathTypeStep(sentence string, shouldMatchType bool, doc stepbuilder.DocParams) stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, jsonPath, expectedType string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)
			expectedType = scenario.ReplaceVariablesInString(scenarioCtx, expectedType)

			responseBody := backend.GetResponseBody()

			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			// Extract the value at the JSON path
			value, err := jsonhelpers.GetPathValue(responseBody, jsonPath)
			if err != nil {
				return ctx, fmt.Errorf("failed to get value at path '%s': %w", jsonPath, err)
			}

			// Get the actual type
			actualType := helpers.GetJSONType(value)

			// Normalize expected type (allow aliases)
			normalizedExpectedType := helpers.NormalizeType(expectedType)

			if shouldMatchType && actualType != normalizedExpectedType {
				return ctx, fmt.Errorf("field '%s' has type '%s', expected '%s'", jsonPath, actualType, normalizedExpectedType)
			}

			if !shouldMatchType && actualType == normalizedExpectedType {
				return ctx, fmt.Errorf("field '%s' has forbidden type '%s'", jsonPath, actualType)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
