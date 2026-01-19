package variables

import (
	"context"
	"errors"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/helpers"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/variables"
)

func (steps) storeValueIntoGlobalVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I store the value {string} into global variable {string}`,
		},
		func(ctx context.Context, value, varName string) (context.Context, error) {
			variables.SetGlobalVariable(varName, value)
			logger.InfoFf("Stored value '%s' into global variable '%s'", value, varName)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a custom value into a global variable accessible by all scenarios.",
			Variables: []stepbuilder.DocVariable{
				{Name: "value", Description: "The value to store", Type: stepbuilder.VarTypeString},
				{Name: "varName", Description: "The name of the global variable", Type: stepbuilder.VarTypeString},
			},
			Categories: []stepbuilder.StepCategory{stepbuilder.Variable},
			Example:    `When I store the value "admin_token" into global variable "AUTH_TOKEN"`,
		},
	)
}

func (steps) storeJSONPathIntoGlobalVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I save the response path {string} as global variable {string}`,
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

			variables.SetGlobalVariable(varName, value)
			logger.InfoFf("Stored JSON path '%s' value '%v' into global variable '%s'", jsonPath, value, varName)

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a value from a JSON response using a JSON path into a global variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "jsonPath", Description: "The JSON path to extract", Type: stepbuilder.VarTypeString},
				{Name: "varName", Description: "The name of the global variable", Type: stepbuilder.VarTypeString},
			},
			Categories: []stepbuilder.StepCategory{stepbuilder.Variable},
			Example:    `Then I save the response path "token" as global variable "AUTH_TOKEN"`,
		},
	)
}
