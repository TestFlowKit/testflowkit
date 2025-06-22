package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) hoverOnElement() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user hovers on {string}$`},
		func(ctx *scenario.Context) func(label string) error {
			return func(label string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), label)
				if err != nil {
					return err
				}
				return element.Hover()
			}
		},
		func(label string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "Hover on a element.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to hover on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user hovers on \"Submit button\"",
			Category: stepbuilder.Mouse,
		},
	)
}
