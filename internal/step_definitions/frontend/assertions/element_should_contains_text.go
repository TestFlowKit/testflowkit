package assertions

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) elementShouldContainsText() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the {string} should contain the text {string}`},
		func(ctx context.Context, name, expectedText string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			element, err := scenarioCtx.GetHTMLElementByLabel(name)
			if err != nil {
				return ctx, err
			}

			if !element.IsVisible() {
				return ctx, fmt.Errorf("%s is not visible", name)
			}

			if !strings.Contains(element.TextContent(), expectedText) {
				return ctx, fmt.Errorf("%s does not contain text '%s'", name, expectedText)
			}

			return ctx, nil
		},
		func(name, _ string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "This assertion checks if the element's visible text includes the specified substring.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "expectedText", Description: "The text that should be contained.", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the welcome card should contain the text "Hello John"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Assertions},
		},
	)
}
