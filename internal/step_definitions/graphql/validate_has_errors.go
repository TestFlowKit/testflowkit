package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) validateGraphQLHasErrors() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`the GraphQL response should contain errors`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			response := scenarioCtx.GetGraphQLResponse()

			if response == nil {
				return ctx, fmt.Errorf("no GraphQL response available - send a GraphQL request first")
			}

			if len(response.Errors) == 0 {
				return ctx, fmt.Errorf("GraphQL response does not contain any errors")
			}

			logger.InfoFf("GraphQL response contains %d error(s)", len(response.Errors))
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the GraphQL response contains errors.",
			Example:     `Then the GraphQL response should contain errors`,
			Category:    stepbuilder.GraphQL,
		},
	)
}
