package graphql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

// validateErrorMessage validates that GraphQL error messages contain specific text.
func (s steps) validateErrorMessage() stepbuilder.Step {
	return s.newValidateErrorMessageStep(
		`the GraphQL error message should contain "([^"]*)"`,
		true,
		stepbuilder.DocParams{
			Description: "Validates that at least one GraphQL error message contains the expected text.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "text",
					Description: "Text that should be present in the error message",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:    `Then the GraphQL error message should contain "User not found"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.GraphQL},
		},
	)
}

// validateErrorMessageNot validates that GraphQL error messages do not contain specific text.
func (s steps) validateErrorMessageNot() stepbuilder.Step {
	return s.newValidateErrorMessageStep(
		`the GraphQL error message should not contain "([^"]*)"`,
		false,
		stepbuilder.DocParams{
			Description: "Validates that none of the GraphQL error messages contain the specified text.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "text",
					Description: "Text that should not be present in the error message",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:    `Then the GraphQL error message should not contain "stack trace"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.GraphQL},
		},
	)
}

func (s steps) newValidateErrorMessageStep(
	sentence string, shouldContain bool, doc stepbuilder.DocParams,
) stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{sentence},
		func(ctx context.Context, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoGraphQLResponse
			}

			if !backend.HasGraphQLErrors() {
				return ctx, errors.New("no GraphQL errors found to validate")
			}

			text = scenario.ReplaceVariablesInString(scenarioCtx, text)
			gqlErrors := backend.GetGraphQLErrors()

			hasMatch := false
			for _, err := range gqlErrors {
				if strings.Contains(err.Message, text) {
					hasMatch = true
					break
				}
			}

			if shouldContain && !hasMatch {
				const msg = "expected GraphQL error message to contain '%s', but none of the %d error(s) matched"
				return ctx, fmt.Errorf(msg, text, len(gqlErrors))
			}

			if !shouldContain && hasMatch {
				return ctx, fmt.Errorf("GraphQL error message contains forbidden text '%s'", text)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
