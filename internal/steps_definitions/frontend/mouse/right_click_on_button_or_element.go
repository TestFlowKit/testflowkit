package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) rightClickOn() stepbuilder.TestStep {
	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user right clicks on {string}$`},
		func(ctx *stepbuilder.TestSuiteContext) func(string) error {
			return func(label string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), label)
				if err != nil {
					return err
				}
				return element.RightClick()
			}
		},
		func(label string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "Right clicks on a button or element.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the element to right click on.", Type: shared.DocVarTypeString},
			},
			Example:  "When the user right clicks on \"Submit button\"",
			Category: shared.Mouse,
		},
	)
}
