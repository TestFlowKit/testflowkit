package mouse

import (
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/browser"
)

func (steps) rightClickOn() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user right clicks on {string}`},
		commonSimpleElementInteraction(func(element browser.Element) error {
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
				{Name: stepbuilder.DocVarName, Description: docDescElementRightClick, Type: stepbuilder.VarTypeString},
			},
			Example:    "When the user right clicks on \"Submit button\"",
			Categories: []stepbuilder.StepCategory{stepbuilder.Mouse, stepbuilder.Frontend},
		},
	)
}
