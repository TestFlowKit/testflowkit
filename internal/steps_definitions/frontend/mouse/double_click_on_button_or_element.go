package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) doubleClickOn() stepbuilder.Step {
	const docDescription = "The logical name of the button or element to double click on."

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user double clicks on {string}$`},
		func(ctx *scenario.Context) func(string) error {
			return func(label string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), label)
				if err != nil {
					return err
				}
				return element.DoubleClick()
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
			Description: "double clicks on a button or element.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: docDescription, Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user double clicks on \"File item\"",
			Category: stepbuilder.Mouse,
		},
	)
}
