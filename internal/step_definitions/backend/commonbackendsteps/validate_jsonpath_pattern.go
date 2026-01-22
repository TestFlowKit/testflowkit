package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) validateJSONPathPattern() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the response field {string} should match pattern {string}`},
		func(ctx context.Context, jsonPath, pattern string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to validate")
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

			if !regex.MatchString(actualValue) {
				return ctx, fmt.Errorf("field '%s' value '%s' does not match pattern '%s'", jsonPath, actualValue, pattern)
			}

			return ctx, nil
		},
		nil,
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
