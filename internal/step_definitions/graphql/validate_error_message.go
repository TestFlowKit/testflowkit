package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) validateGraphQLErrorMessage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the GraphQL response should contain error message {string}`},
		func(ctx context.Context, expectedMessage string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			response := scenarioCtx.GetGraphQLResponse()

			if response == nil {
				return ctx, fmt.Errorf("no GraphQL response available - send a GraphQL request first")
			}

			if len(response.Errors) == 0 {
				return ctx, fmt.Errorf("GraphQL response does not contain any errors")
			}

			for _, err := range response.Errors {
				if err.Message == expectedMessage {
					logger.InfoFf("Found expected error message: %s", expectedMessage)
					return ctx, nil
				}
			}

			errorMessages := make([]string, len(response.Errors))
			for i, err := range response.Errors {
				errorMessages[i] = err.Message
			}

			return ctx, fmt.Errorf("expected error message '%s' not found in GraphQL errors: %v", expectedMessage, errorMessages)
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the GraphQL response contains a specific error message.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "expectedMessage",
					Description: "The expected error message",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Then the GraphQL response should contain error message "User not found"`,
			Category: stepbuilder.GraphQL,
		},
	)
}
