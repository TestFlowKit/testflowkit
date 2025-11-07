package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) validateGraphQLErrors() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`the GraphQL response should not contain errors`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			response := scenarioCtx.GetGraphQLResponse()

			if response == nil {
				return ctx, fmt.Errorf("no GraphQL response available - send a GraphQL request first")
			}

			if len(response.Errors) > 0 {
				errorMessages := make([]string, len(response.Errors))
				for i, err := range response.Errors {
					errorMessages[i] = err.Message
				}
				return ctx, fmt.Errorf("GraphQL response contains errors: %v", errorMessages)
			}

			logger.InfoFf("GraphQL response contains no errors")
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the GraphQL response does not contain any errors.",
			Example:     `Then the GraphQL response should not contain errors`,
			Category:    stepbuilder.GraphQL,
		},
	)
}
