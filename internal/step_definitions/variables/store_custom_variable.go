package variables

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) storeCustomVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I store the value {string} into {string} variable`,
		},
		func(ctx context.Context, value, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			scenarioCtx.SetVariable(varName, value)

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a custom value into a scenario variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: stepbuilder.DocVarValue, Description: "The value to store in the variable", Type: stepbuilder.VarTypeString},
				{
					Name:        stepbuilder.DocVarVarName,
					Description: "The name of the variable to store the value in",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:    `When I store the value "John Doe" into "displayed_name" variable`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Variable},
		},
	)
}
