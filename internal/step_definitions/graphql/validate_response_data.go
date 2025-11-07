package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"

	"github.com/tidwall/gjson"
)

func (steps) validateGraphQLResponse() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the GraphQL response should contain data at path {string}`},
		func(ctx context.Context, jsonPath string) (context.Context, error) {
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

			logger.InfoFf("GraphQL response validation passed for path: %s", jsonPath)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the GraphQL response contains data at the specified JSON path.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "jsonPath",
					Description: "JSON path to validate in the response data",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Then the GraphQL response should contain data at path "user.name"`,
			Category: stepbuilder.GraphQL,
		},
	)
}
