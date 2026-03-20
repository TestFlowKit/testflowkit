package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

func (s steps) validateJSONPathPattern() stepbuilder.Step {
	return s.newJSONPathPatternStep(
		`the response field {string} should match pattern {string}`,
		true,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path field matches a regular expression pattern.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to the field to validate", Type: stepbuilder.VarTypeString},
				{Name: "pattern", Description: "Regular expression pattern to match against", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response field "user.email" should match pattern "^[a-z]+@example\\.com$"`,
			Categories: stepbuilder.Backend,
		},
	)
}

func (s steps) validateJSONPathNotPattern() stepbuilder.Step {
	return s.newJSONPathPatternStep(
		`the response field {string} should not match pattern {string}`,
		false,
		stepbuilder.DocParams{
			Description: "Validates that a specific JSON path field does not match a regular expression pattern.",
			Variables: []stepbuilder.DocVariable{
				{Name: "path", Description: "JSON path to the field to validate", Type: stepbuilder.VarTypeString},
				{Name: "pattern", Description: "Regular expression pattern that should not match", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response field "user.email" should not match pattern ".*@internal\\.com$"`,
			Categories: stepbuilder.Backend,
		},
	)
}

func (s steps) newJSONPathPatternStep(sentence string, shouldMatch bool, doc stepbuilder.DocParams) stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, jsonPath, pattern string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)
			pattern = scenario.ReplaceVariablesInString(scenarioCtx, pattern)

			responseBody := backend.GetResponseBody()

			if responseBody == nil {
				return ctx, errors.New("response body is empty")
			}

			actualValue, err := jsonhelpers.GetPathValueAsString(responseBody, jsonPath)
			if err != nil {
				return ctx, fmt.Errorf("failed to get value at path '%s': %w", jsonPath, err)
			}

			regex, err := regexp.Compile(pattern)
			if err != nil {
				return ctx, fmt.Errorf("invalid regex pattern '%s': %w", pattern, err)
			}

			matches := regex.MatchString(actualValue)
			if shouldMatch && !matches {
				return ctx, fmt.Errorf("field '%s' value '%s' does not match pattern '%s'", jsonPath, actualValue, pattern)
			}

			if !shouldMatch && matches {
				return ctx, fmt.Errorf("field '%s' value '%s' should not match pattern '%s'", jsonPath, actualValue, pattern)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
