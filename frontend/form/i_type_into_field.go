package form

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
)

func (s steps) iTypeXXXIntoInput() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I type "{string}" into the {string}`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(text, inputLabel string) error {
				input, err := browser.GetElement(ctx.GetCurrentPage(), inputLabel)
				if err != nil {
					return err
				}
				return input.Input(text)
			}
		},
	}
}
