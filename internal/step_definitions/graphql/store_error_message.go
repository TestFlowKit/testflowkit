package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) storeGraphQLErrorMessage() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store the GraphQL error message at index {string} into {string} variable`},
		func(ctx context.Context, indexStr, variableName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			response := scenarioCtx.GetGraphQLResponse()

			if response == nil {
				return ctx, fmt.Errorf("no GraphQL response available - send a GraphQL request first")
			}

			if len(response.Errors) == 0 {
				return ctx, fmt.Errorf("GraphQL response does not contain any errors")
			}

			// Parse index
			var index int
			if indexStr == "first" || indexStr == "0" {
				index = 0
			} else {
				return ctx, fmt.Errorf("unsupported error index '%s' - use 'first' or '0'", indexStr)
			}

			if index >= len(response.Errors) {
				return ctx, fmt.Errorf("error index %d is out of range - response has %d error(s)", index, len(response.Errors))
			}

			errorMessage := response.Errors[index].Message

			// Store the error message in the scenario context
			scenarioCtx.SetVariable(variableName, errorMessage)

			logger.InfoFf("Stored GraphQL error message at index %d into variable '%s': %s", index, variableName, errorMessage)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a GraphQL error message into a variable for later use.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "index",
					Description: "Error index ('first' or '0' for the first error)",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "variableName",
					Description: "Name of the variable to store the error message",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And I store the GraphQL error message at index "first" into "errorMessage" variable`,
			Category: stepbuilder.GraphQL,
		},
	)
}
