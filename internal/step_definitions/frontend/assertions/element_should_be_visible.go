package assertions

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) elementShouldBeVisible() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the {string} should be visible`},
		func(ctx context.Context, elementName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			element, err := scenarioCtx.GetHTMLElementByLabel(elementName)
			if err != nil {
				return ctx, err
			}

			if !element.IsVisible() {
				return ctx, fmt.Errorf("%s is not visible", elementName)
			}

			return ctx, nil
		},
		func(name string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "This assertion checks if the element is present in the DOM and displayed on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element.", Type: stepbuilder.VarTypeString},
			},
			Example:    "Then the submit button should be visible",
			Categories: []stepbuilder.StepCategory{stepbuilder.Assertions},
		},
	)
}
