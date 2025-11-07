package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) storeGraphQLError() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store the GraphQL error at index {string} into {string} variable`},
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

			errorInfo := map[string]interface{}{
				"message":    response.Errors[index].Message,
				"locations":  response.Errors[index].Locations,
				"path":       response.Errors[index].Path,
				"extensions": response.Errors[index].Extensions,
			}

			// Store the error information in the scenario context
			scenarioCtx.SetVariable(variableName, errorInfo)

			logger.InfoFf("Stored GraphQL error at index %d into variable '%s': %s", index, variableName, response.Errors[index].Message)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores GraphQL error information into a variable for later use.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "index",
					Description: "Error index ('first' or '0' for the first error)",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "variableName",
					Description: "Name of the variable to store the error information",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And I store the GraphQL error at index "first" into "errorInfo" variable`,
			Category: stepbuilder.GraphQL,
		},
	)
}
