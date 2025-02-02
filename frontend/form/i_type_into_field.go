package form

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"etoolse/internal/config/testsConfig"
)

func (s steps) iTypeXXXIntoInput() common.TestStep {
	return common.NewStepWithTwoVariables(
		[]string{`^I type "{string}" into the {string}`},
		func(ctx *common.TestSuiteContext) func(string, string) error {
			return func(text, inputLabel string) error {
				input, err := browser.GetElement(ctx.GetCurrentPage(), inputLabel)
				if err != nil {
					return err
				}
				return input.Input(text)
			}
		},
		func(_, inputLabel string) common.ValidationErrors {
			vc := common.ValidationErrors{}
			if !testsConfig.IsElementDefined(inputLabel) {
				vc.AddMissingElement(inputLabel)
			}

			return vc
		},
	)
}
