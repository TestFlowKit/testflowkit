package assertions

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) elementShouldNotBeVisible() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the {string} should not be visible`},
		func(ctx context.Context, name string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			element, err := scenarioCtx.GetHTMLElementByLabel(name)
			if err != nil {
				return ctx, err
			}

			if element.IsVisible() {
				return ctx, fmt.Errorf("%s should not be visible", name)
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
			Description: "verifies that an element is not visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then \"Submit button\" should not be visible",
			Category: stepbuilder.Visual,
		},
	)
}
