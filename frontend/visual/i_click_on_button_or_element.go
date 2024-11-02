package visual

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
)

var iClickOnButtonOrElement = common.FrontStep{
	Sentences: []string{`^I click on "{string}" {string}$`},
	Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
		return func(label string) error {
			button := browser.GetElement(ctx.GetCurrentPage(), label)
			return button.Click()
		}
	},
}
