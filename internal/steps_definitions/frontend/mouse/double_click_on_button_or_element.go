package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) doubleClickOn() core.TestStep {
	const docDescription = "The logical name of the button or element to double click on."

	return core.NewStepWithOneVariable(
		[]string{`^the user double clicks on {string}$`},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(label string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), label)
				if err != nil {
					return err
				}
				return element.DoubleClick()
			}
		},
		func(label string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		core.StepDefDocParams{
			Description: "double clicks on a button or element.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: docDescription, Type: shared.DocVarTypeString},
			},
			Example:  "When the user double clicks on \"File item\"",
			Category: shared.Mouse,
		},
	)
}
