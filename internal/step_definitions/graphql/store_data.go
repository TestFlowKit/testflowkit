package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"

	"github.com/tidwall/gjson"
)

func (steps) storeGraphQLData() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store the GraphQL data at path {string} into {string} variable`},
		func(ctx context.Context, jsonPath, variableName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			response := scenarioCtx.GetGraphQLResponse()

			if response == nil {
				return ctx, fmt.Errorf("no GraphQL response available - send a GraphQL request first")
			}

			if response.Data == nil {
				return ctx, fmt.Errorf("GraphQL response contains no data")
			}

			value := gjson.GetBytes(response.Data, jsonPath)
			if !value.Exists() {
				return ctx, fmt.Errorf("path '%s' not found in GraphQL response data", jsonPath)
			}

			// Store the value in the scenario context (TestFlowKit variable system)
			scenarioCtx.SetVariable(variableName, value.Value())

			logger.InfoFf("Stored GraphQL data from path '%s' into variable '%s': %v", jsonPath, variableName, value.Value())
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores data from the GraphQL response into a variable for later use.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "jsonPath",
					Description: "JSON path to extract data from the response",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "variableName",
					Description: "Name of the variable to store the extracted data",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And I store the GraphQL data at path "user.id" into "userId" variable`,
			Category: stepbuilder.GraphQL,
		},
	)
}
