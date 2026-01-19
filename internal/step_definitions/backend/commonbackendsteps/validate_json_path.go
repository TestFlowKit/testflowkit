package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
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
