package form

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
)

func (s steps) iTypeXXXIntoInput() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I type "{string}" into the {string}`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(text, inputLabel string) error {
				input, err := browser.GetInputElement(ctx.GetCurrentPage(), inputLabel)
				if err != nil {
					return err
				}
				return input.Input(text)
			}
		},
	}
}
