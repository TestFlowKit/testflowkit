package graphql

import (
	"context"
	"encoding/json"
	"fmt"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"

	"github.com/tidwall/gjson"
)

func (steps) storeGraphQLArrayData() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`I store the GraphQL array at path {string} into {string} variable`},
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

			if !value.IsArray() {
				return ctx, fmt.Errorf("path '%s' does not contain an array", jsonPath)
			}

			var arr []interface{}
			if err := json.Unmarshal([]byte(value.Raw), &arr); err != nil {
				return ctx, fmt.Errorf("failed to parse array at path '%s': %w", jsonPath, err)
			}

			// Store the array in the scenario context (TestFlowKit variable system)
			scenarioCtx.SetVariable(variableName, arr)

			logger.InfoFf("Stored GraphQL array from path '%s' into variable '%s' (length: %d)", jsonPath, variableName, len(arr))
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores an array from the GraphQL response into a variable for later use.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "jsonPath",
					Description: "JSON path to extract array data from the response",
					Type:        stepbuilder.VarTypeString,
				},
				{
					Name:        "variableName",
					Description: "Name of the variable to store the extracted array",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `And I store the GraphQL array at path "user.tags" into "userTags" variable`,
			Category: stepbuilder.GraphQL,
		},
	)
}
