package visual

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
	"fmt"
)

var iMustSeeOnPageXElements = common.FrontStep{
	Sentences: []string{`^I must see on page {number} {string}$`},
	Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
		return func(expectedCount int, elementName string) error {
			elementCount := browser.GetElementCount(ctx.GetCurrentPage(), elementName)
			if elementCount != expectedCount {
				return fmt.Errorf("%d %s expected but %d %s found", expectedCount, elementName, elementCount, elementName)
			}
			return nil
		}
	},
}
