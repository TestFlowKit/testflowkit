package graphql

import (
	"context"
	"errors"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

// storeGraphQLError stores the entire GraphQL error array as a variable.
func (steps) storeGraphQLError() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I store the GraphQL error as "([^"]*)"`},
		func(ctx context.Context, variableName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no GraphQL response available - send a request first")
			}

			if !backend.HasGraphQLErrors() {
				return ctx, errors.New("no GraphQL errors found to store")
			}

			errors := backend.GetGraphQLErrors()

			// Store the entire error array
			backend.SetGraphQLVariable(variableName, errors)
			scenarioCtx.SetVariable(variableName, errors)

			logger.InfoFf("Stored GraphQL errors as '%s': %d error(s)", variableName, len(errors))
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores the entire GraphQL error array in a variable for later use.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "variableName",
					Description: "Name of the variable to store the errors",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:    `And I store the GraphQL error as "lastError"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.GraphQL},
		},
	)
}

// storeGraphQLErrorMessage stores the first GraphQL error message as a variable.
func (steps) storeGraphQLErrorMessage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I store the GraphQL error message as "([^"]*)"`},
		func(ctx context.Context, variableName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no GraphQL response available - send a request first")
			}

			if !backend.HasGraphQLErrors() {
				return ctx, errors.New("no GraphQL errors found to store")
			}

			gqlErrors := backend.GetGraphQLErrors()
			if len(gqlErrors) == 0 {
				return ctx, errors.New("error array is empty")
			}

			// Store the first error message
			errorMessage := gqlErrors[0].Message
			backend.SetGraphQLVariable(variableName, errorMessage)
			scenarioCtx.SetVariable(variableName, errorMessage)

			logger.InfoFf("Stored GraphQL error message as '%s': %s", variableName, errorMessage)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores the first GraphQL error message in a variable for later use.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "variableName",
					Description: "Name of the variable to store the error message",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:    `And I store the GraphQL error message as "errorMsg"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.GraphQL},
		},
	)
}
