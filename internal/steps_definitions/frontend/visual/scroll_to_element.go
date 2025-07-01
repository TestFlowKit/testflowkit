package visual

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) scrollToElement() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user scrolls to the {string} element`},
		func(ctx context.Context, elementName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			elt, err := scenarioCtx.GetHTMLElementByLabel(fmt.Sprintf("%s_element", elementName))
			if err != nil {
				return ctx, err
			}

			scrollErr := elt.ScrollIntoView()
			if scrollErr != nil {
				return ctx, scrollErr
			}
			return ctx, nil
		},
		func(name string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			variable := fmt.Sprintf("%s_element", name)
			if !config.IsElementDefined(variable) {
				vc.AddMissingElement(variable)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "scrolls to the specified element.",
			Variables: []stepbuilder.DocVariable{
				{Name: "elementName", Description: "The name of the element to scroll to.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user scrolls to the submit button element",
			Category: stepbuilder.Visual,
		},
	)
}
