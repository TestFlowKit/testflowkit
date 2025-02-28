package visual

import (
	"etoolse/internal/browser"
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
	"etoolse/shared"
)

func (s steps) iClickOn() core.TestStep {
	return core.NewStepWithOneVariable(
		[]string{`^I click on {string}$`},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(label string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), label)
				if err != nil {
					return err
				}
				return element.Click()
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
			Description: "clicks on a button or element.",
			Variables: []shared.StepVariable{
				{Name: "label", Description: "The label of the button or element to click on.", Type: shared.DocVarTypeString},
			},
			Example:  "When I click on \"Submit button\"",
			Category: shared.Visual,
		},
	)
}
