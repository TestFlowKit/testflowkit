package variables

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) variableShouldContains() stepbuilder.Step {
	contentDesc := "The string that should be present in the variable"
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`the variable {string} should contain {string}`,
		},
		func(ctx context.Context, varName, content string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			variable, exists := scenarioCtx.GetVariable(varName)
			if !exists {
				return ctx, fmt.Errorf("variable '%s' not found", varName)
			}

			variableValue, ok := variable.(string)
			if !ok {
				return ctx, fmt.Errorf("variable '%s' is not a string", varName)
			}

			if !strings.Contains(variableValue, content) {
				return ctx, fmt.Errorf("variable '%s' does not contain '%s'", varName, content)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that a variable contains a specific string.",
			Variables: []stepbuilder.DocVariable{
				{Name: "varName", Description: "The name of the variable to check", Type: stepbuilder.VarTypeString},
				{Name: "content", Description: contentDesc, Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the variable "user_name_label" should contain "John Doe"`,
			Category: stepbuilder.Variable,
		},
	)
}
