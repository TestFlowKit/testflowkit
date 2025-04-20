package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) iRightClickOn() core.TestStep {
	return core.NewStepWithOneVariable(
		[]string{`^I right click on {string}$`},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(label string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), label)
				if err != nil {
					return err
				}
				return element.RightClick()
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
			Description: "Right clicks on a button or element.",
			Variables: []shared.StepVariable{
				{Name: "label", Description: "The label of element to right click on.", Type: shared.DocVarTypeString},
			},
			Example:  "When I right click on \"Submit button\"",
			Category: shared.Mouse,
		},
	)
}
