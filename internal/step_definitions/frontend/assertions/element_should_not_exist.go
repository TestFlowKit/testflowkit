package assertions

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) elementShouldNotExist() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the {string} should not exist`},
		func(ctx context.Context, name string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			_, err := scenarioCtx.GetHTMLElementByLabel(name)
			if err == nil {
				return ctx, fmt.Errorf("%s should not exist", name)
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
			Description: "verifies that an element does not exist on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the submit button should not exist",
			Category: stepbuilder.Visual,
		},
	)
}
