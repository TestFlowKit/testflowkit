package variables

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/logger"
)

func (steps) displayVariableValue() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{
			`I display the value of variable {string}`,
		},
		func(ctx context.Context, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			value, exists := scenarioCtx.GetVariable(varName)
			if !exists {
				return ctx, &apperrors.VariableNotFoundError{Name: varName}
			}

			logger.InfoFf("Variable '%s' value: %v", varName, value)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Displays the current value of a scenario variable in logs.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        stepbuilder.DocVarVarName,
					Description: "The name of the variable to display",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:    `Then I display the value of variable "user_id"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Variable, stepbuilder.Debug},
		},
	)
}
