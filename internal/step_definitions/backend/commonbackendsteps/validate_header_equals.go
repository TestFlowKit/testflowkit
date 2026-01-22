package commonbackendsteps

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/helpers"
)

func (steps) validateResponseHeaderEquals() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the response header {string} should equal {string}`},
		func(ctx context.Context, headerName, expectedValue string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no response available to validate")
			}

			headerName = scenario.ReplaceVariablesInString(scenarioCtx, headerName)
			expectedValue = scenario.ReplaceVariablesInString(scenarioCtx, expectedValue)

			response := backend.GetResponse()
			if response == nil || response.Headers == nil {
				return ctx, errors.New("response headers are not available")
			}

			actualValue, found := helpers.GetHeaderCaseInsensitive(response.Headers, headerName)
			if !found {
				return ctx, fmt.Errorf("response header '%s' not found", headerName)
			}

			if actualValue != expectedValue {
				return ctx, fmt.Errorf("header '%s' has value '%s', expected '%s'", headerName, actualValue, expectedValue)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that a specific response header has the expected value.",
			Variables: []stepbuilder.DocVariable{
				{Name: "header", Description: "Name of the header to validate", Type: stepbuilder.VarTypeString},
				{Name: "value", Description: "Expected value of the header", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the response header "Content-Type" should equal "application/json"`,
			Categories: stepbuilder.Backend,
		},
	)
}
