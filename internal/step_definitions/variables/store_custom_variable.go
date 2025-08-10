package variables

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) storeCustomVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I store the value {string} into {string} variable`,
		},
		func(ctx context.Context, value, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			scenarioCtx.SetVariable(varName, value)
			logger.InfoFf("Stored value '%s' into variable '%s'", value, varName)

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores a custom value into a scenario variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "value", Description: "The value to store in the variable", Type: stepbuilder.VarTypeString},
				{Name: "varName", Description: "The name of the variable to store the value in", Type: stepbuilder.VarTypeString},
			},
			Example:  `When I store the value "John Doe" into "displayed_name" variable`,
			Category: stepbuilder.Variable,
		},
	)
}
