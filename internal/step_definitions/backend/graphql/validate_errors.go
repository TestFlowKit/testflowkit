package graphql

import (
	"context"
	"errors"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) validateHaveErrors() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`the GraphQL response should have errors`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no GraphQL response available - send a request first")
			}

			if !backend.HasGraphQLErrors() {
				return ctx, errors.New("expected GraphQL errors but found none")
			}

			logger.InfoFf("GraphQL response validation passed: errors present")
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

func (steps) validateNoErrors() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`the GraphQL response should not have errors`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, errors.New("no GraphQL response available - send a request first")
			}

			if backend.HasGraphQLErrors() {
				errors := backend.GetGraphQLErrors()
				return ctx, fmt.Errorf("expected no GraphQL errors but found %d: %v", len(errors), errors)
			}

			logger.InfoFf("GraphQL response validation passed: no errors")
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the GraphQL response contains no errors.",
			Example:     `Then the GraphQL response should not have errors`,
			Categories:  []stepbuilder.StepCategory{stepbuilder.GraphQL},
		},
	)
}
