package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) iDoubleClickOn() core.TestStep {
	const docDescription = "The label of the button or element to double click on."

	return core.NewStepWithOneVariable(
		[]string{`^I double click on {string}$`},
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
				{Name: "label", Description: docDescription, Type: shared.DocVarTypeString},
			},
			Example:  "When I double click on \"File item\"",
			Category: shared.Mouse,
		},
	)
}
