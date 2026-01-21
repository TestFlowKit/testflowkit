package variables

import (
	"context"
	"testflowkit/internal/step_definitions/core/stepbuilder"
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
