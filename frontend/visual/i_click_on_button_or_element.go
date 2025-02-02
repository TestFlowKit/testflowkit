package visual

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"etoolse/internal/config/testsConfig"
)

func (s steps) iClickOn() common.TestStep {
	return common.NewStepWithOneVariable(
		[]string{`^I click on {string}$`},
		func(ctx *common.TestSuiteContext) func(string) error {
			return func(label string) error {
				element, err := browser.GetElement(ctx.GetCurrentPage(), label)
				if err != nil {
					return err
				}
				return element.Click()
			}
		},
		func(label string) common.ValidationErrors {
			vc := common.ValidationErrors{}
			if !testsConfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
	)
}
