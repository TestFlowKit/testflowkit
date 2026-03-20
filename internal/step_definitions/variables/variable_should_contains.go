package variables

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
)

func (s steps) variableShouldContains() stepbuilder.Step {
	return s.newVariableContainsStep(
		`the variable {string} should contain {string}`,
		true,
		stepbuilder.DocParams{
			Description: "Verifies that a variable contains a specific string.",
			Variables: []stepbuilder.DocVariable{
				{Name: "varName", Description: "The name of the variable to check", Type: stepbuilder.VarTypeString},
				{Name: "content", Description: "The string that should be present in the variable",
					Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the variable "user_name_label" should contain "John Doe"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Variable},
		},
	)
}

func (s steps) variableShouldNotContains() stepbuilder.Step {
	return s.newVariableContainsStep(
		`the variable {string} should not contain {string}`,
		false,
		stepbuilder.DocParams{
			Description: "Verifies that a variable does not contain a specific string.",
			Variables: []stepbuilder.DocVariable{
				{Name: "varName", Description: "The name of the variable to check", Type: stepbuilder.VarTypeString},
				{Name: "content", Description: "The string that should not be present in the variable",
					Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the variable "error_message" should not contain "panic"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Variable},
		},
	)
}

func (s steps) newVariableContainsStep(
	sentence string, shouldContain bool, doc stepbuilder.DocParams,
) stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, varName, content string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			variable, exists := scenarioCtx.GetVariable(varName)
			if !exists {
				return ctx, &apperrors.VariableNotFoundError{Name: varName}
			}

			variableValue, ok := variable.(string)
			if !ok {
				return ctx, &apperrors.VariableNotStringError{Name: varName}
			}

			contains := strings.Contains(variableValue, content)
			if shouldContain && !contains {
				return ctx, fmt.Errorf("variable '%s' does not contain '%s'", varName, content)
			}

			if !shouldContain && contains {
				return ctx, fmt.Errorf("variable '%s' should not contain '%s'", varName, content)
			}

			return ctx, nil
		},
		nil,
		doc,
	)
}
