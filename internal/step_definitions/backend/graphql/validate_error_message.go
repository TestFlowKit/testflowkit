package graphql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

// validateErrorMessage validates that GraphQL error messages contain specific text.
func (steps) validateErrorMessage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the GraphQL error message should contain "([^"]*)"`},
		func(ctx context.Context, expectedText string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no GraphQL response available - send a request first")
			}

			if !backend.HasGraphQLErrors() {
				return ctx, errors.New("no GraphQL errors found to validate")
			}

			expectedText = scenario.ReplaceVariablesInString(scenarioCtx, expectedText)
			errors := backend.GetGraphQLErrors()

			// Check if any error message contains the expected text
			for _, err := range errors {
				if strings.Contains(err.Message, expectedText) {
					logger.InfoFf("GraphQL error message validation passed: found '%s'", expectedText)
					return ctx, nil
				}
			}

			const msg = "expected GraphQL error message to contain '%s', but none of the %d error(s) matched"
			return ctx, fmt.Errorf(msg, expectedText, len(errors))
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that at least one GraphQL error message contains the expected text.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "text",
					Description: "Text that should be present in the error message",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Then the GraphQL error message should contain "User not found"`,
			Category: stepbuilder.GraphQL,
		},
	)
}
