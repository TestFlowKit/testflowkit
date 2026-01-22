package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/helpers"
)

func (steps) validateResponseHeaderPattern() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the response header {string} should match pattern {string}`},
		func(ctx context.Context, headerName, pattern string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to validate")
			}

			headerName = scenario.ReplaceVariablesInString(scenarioCtx, headerName)
			pattern = scenario.ReplaceVariablesInString(scenarioCtx, pattern)

			response := backend.GetResponse()
			if response == nil || response.Headers == nil {
				return ctx, errors.New("response headers are not available")
			}

			// Case-insensitive header lookup (HTTP standard)
			actualValue, found := helpers.GetHeaderCaseInsensitive(response.Headers, headerName)
			if !found {
				return ctx, fmt.Errorf("response header '%s' not found", headerName)
			}

			// Compile and match the regex pattern
			regex, err := regexp.Compile(pattern)
			if err != nil {
				return ctx, fmt.Errorf("invalid regex pattern '%s': %w", pattern, err)
			}

			if !regex.MatchString(actualValue) {
				return ctx, fmt.Errorf("header '%s' value '%s' does not match pattern '%s'", headerName, actualValue, pattern)
			}

			return ctx, nil
		},
		nil,
		func() stepbuilder.DocParams {
			dsc := "Validates that a specific response header matches a regular expression pattern."

			return stepbuilder.DocParams{
				Description: dsc,
				Variables: []stepbuilder.DocVariable{
					{Name: "header", Description: "Name of the header to validate", Type: stepbuilder.VarTypeString},
					{Name: "pattern", Description: "Regular expression pattern to match against", Type: stepbuilder.VarTypeString},
				},
				Example:    `Then the response header "Content-Type" should match pattern "application/(json|xml)"`,
				Categories: stepbuilder.Backend,
			}
		}(),
	)
}
