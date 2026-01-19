package mouse

import (
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/browser"
)

func (steps) hoverOnElement() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user hovers on {string}`},
		commonSimpleElementInteraction(func(element browser.Element) error {
			return element.Hover()
		}),
		func(label string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "performs a hover action on the element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of element to hover on.", Type: stepbuilder.VarTypeString},
			},
			Example:    "When the user hovers on \"Submit button\"",
			Categories: []stepbuilder.StepCategory{stepbuilder.Mouse},
		},
	)
}
