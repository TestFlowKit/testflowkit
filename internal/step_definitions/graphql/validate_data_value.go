package graphql

import (
	"context"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"

	"github.com/tidwall/gjson"
)

func (steps) validateGraphQLDataValue() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the GraphQL response data at path {string} should be {string}`},
		func(ctx context.Context, jsonPath, expectedValue string) (context.Context, error) {
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

			actualValue := value.String()
			if actualValue != expectedValue {
				return ctx, fmt.Errorf("expected value '%s' at path '%s', but got '%s'", expectedValue, jsonPath, actualValue)
			}

			logger.InfoFf("GraphQL response data validation passed: %s = %s", jsonPath, expectedValue)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Validates that the GraphQL response data at a specific path matches the expected value.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "jsonPath",
					Description: "JSON path to validate in the response data",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "expectedValue",
					Description: "The expected value at the specified path",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Then the GraphQL response data at path "user.name" should be "John Doe"`,
			Category: stepbuilder.GraphQL,
		},
	)
}
