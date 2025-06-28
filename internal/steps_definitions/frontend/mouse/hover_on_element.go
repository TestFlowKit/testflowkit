package mouse

import (
	"testflowkit/internal/browser/common"
	"testflowkit/internal/config"

	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) hoverOnElement() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user hovers on {string}$`},
		commonSimpleElementInteraction(func(element common.Element) error {
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
			Example:  "When the user hovers on \"Submit button\"",
			Category: stepbuilder.Mouse,
		},
	)
}
