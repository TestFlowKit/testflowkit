package graphql

import (
	"context"
	"errors"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
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
				return ctx, apperrors.ErrNoGraphQLResponse
			}

			if !backend.HasGraphQLErrors() {
				return ctx, errors.New("no GraphQL errors found to store")
			}

			gqlErrors := backend.GetGraphQLErrors()

			// Store the entire error array
			backend.SetGraphQLVariable(variableName, gqlErrors)
			scenarioCtx.SetVariable(variableName, gqlErrors)

			logger.InfoFf("Stored GraphQL errors as '%s': %d error(s)", variableName, len(gqlErrors))
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
