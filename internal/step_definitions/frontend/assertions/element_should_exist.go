package assertions

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) elementShouldExist() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the {string} should exist`},
		func(ctx context.Context, name string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			_, err := scenarioCtx.GetHTMLElementByLabel(name)
			return ctx, err
		},
		func(name string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "verifies that an element exists on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the submit button should exist",
			Category: stepbuilder.Visual,
		},
	)
}
