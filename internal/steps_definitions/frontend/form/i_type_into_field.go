package form

import (
	"etoolse/internal/browser"
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
)

func (s steps) iTypeXXXIntoInput() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I type "{string}" into the {string}`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(text, inputLabel string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), inputLabel)
				if err != nil {
					return err
				}
				return input.Input(text)
			}
		},
		func(_, inputLabel string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsElementDefined(inputLabel) {
				vc.AddMissingElement(inputLabel)
			}

			return vc
		},
	)
}
