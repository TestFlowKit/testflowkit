package visual

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
)

var iClickOnButtonOrElement = common.FrontStep{
	Sentences: []string{`^I click on {string}$`},
	Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
		return func(label string) error {
			element := browser.GetElement(ctx.GetCurrentPage(), label)
			return element.Click()
		}
	},
}
