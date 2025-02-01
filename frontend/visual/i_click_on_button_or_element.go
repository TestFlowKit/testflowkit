package visual

import (
	"etoolse/config"
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
)

func (s steps) iClickOn() common.FrontStep {
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
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
	)
}
