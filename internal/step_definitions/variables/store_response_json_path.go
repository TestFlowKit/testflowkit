package variables

import (
	"context"
	"errors"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/helpers"
	"testflowkit/pkg/logger"
)

func (steps) storeJSONPathIntoVariable() stepbuilder.Step {
	finalDescription := "The JSON path to extract the value from (e.g., 'data.user.id', 'items[0].name')"
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I store the JSON path {string} from the response into {string} variable`,
		},
		func(ctx context.Context, jsonPath, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			response := scenarioCtx.GetResponse()
			if response == nil {
				return ctx, errors.New("no response available. Please send a request first")
			}

			value, err := helpers.GetJSONPathValue(response.Body, jsonPath)
			if err != nil {
				return ctx, fmt.Errorf("failed to extract JSON path '%s': %w", jsonPath, err)
			}

			scenarioCtx.SetVariable(varName, value)
			logger.InfoFf("Stored JSON path '%s' value '%v' into variable '%s'", jsonPath, value, varName)

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a value from a JSON response using a JSON path into a scenario variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "jsonPath", Description: finalDescription, Type: stepbuilder.VarTypeString},
				{Name: "varName", Description: "The name of the variable to store the value in", Type: stepbuilder.VarTypeString},
			},
			Example:  `When I store the response JSON path "data.user.id" from the response into "user_id" variable`,
			Category: stepbuilder.RESTAPI,
		},
	)
}
