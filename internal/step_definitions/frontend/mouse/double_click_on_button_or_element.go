package mouse

import (
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/browser"
)

func (steps) doubleClickOn() stepbuilder.Step {
	const docDescription = "The logical name of element to double click on."

	return stepbuilder.NewWithOneVariable(
		[]string{`the user double clicks on {string}`},
		commonSimpleElementInteraction(func(element browser.Element) error {
			return element.DoubleClick()
		}),
		func(label string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "performs a double click action on the element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: docDescription, Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user double clicks on \"File item\"",
			Category: stepbuilder.Mouse,
		},
	)
}
