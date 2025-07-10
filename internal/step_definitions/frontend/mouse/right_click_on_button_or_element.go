package mouse

import (
	"testflowkit/internal/browser/common"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) rightClickOn() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user right clicks on {string}`},
		commonSimpleElementInteraction(func(element common.Element) error {
			return element.RightClick()
		}),
		func(label string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "performs a right click action on the element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of element to right click on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user right clicks on \"Submit button\"",
			Category: stepbuilder.Mouse,
		},
	)
}
