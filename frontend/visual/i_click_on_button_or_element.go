package visual

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
)

func (s steps) iClickOn() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I click on {string}$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(label string) error {
				element, err := browser.GetElement(ctx.GetCurrentPage(), label)
				if err != nil {
					return err
				}

				return element.Click()
			}
		},
	}
}
