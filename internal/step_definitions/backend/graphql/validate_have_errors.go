package graphql

import (
	"context"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/logger"
)

func (steps) validateHaveErrors() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`the GraphQL response should have errors`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoGraphQLResponse
			}

			if !backend.HasGraphQLErrors() {
				return ctx, apperrors.ErrNoGraphQLErrors
			}

			logger.Infof("GraphQL response validation passed: errors present")
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the GraphQL response contains at least one error.",
			Example:     `Then the GraphQL response should have errors`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.GraphQL},
		},
	)
}
