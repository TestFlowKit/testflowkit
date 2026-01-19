package variables

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) storeElementContentIntoVariable() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{
			`I store the content of {string} into {string} variable`,
		},
		func(ctx context.Context, elementName, varName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			element, err := scenarioCtx.GetHTMLElementByLabel(elementName)
			if err != nil {
				return ctx, fmt.Errorf("failed to find element '%s': %w", elementName, err)
			}

			content := element.TextContent()

			scenarioCtx.SetVariable(varName, content)
			logger.InfoFf("Stored content '%s' from element '%s' into variable '%s'", content, elementName, varName)

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Stores the text content of an HTML element into a scenario variable.",
			Variables: []stepbuilder.DocVariable{
				{Name: "elementName", Description: "The logical name of the HTML element", Type: stepbuilder.VarTypeString},
				{Name: "varName", Description: "The name of the variable to store the content in", Type: stepbuilder.VarTypeString},
			},
			Example:    `When I store the content of "user_name_label" into "displayed_name" variable`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Variable},
		},
	)
}
