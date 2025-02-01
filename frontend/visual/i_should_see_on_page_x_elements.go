package visual

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"fmt"
)

func (s steps) iShouldSeeOnPageXElements() common.FrontStep {
	return common.NewStepWithTwoVariables(
		[]string{`^I should see {number} {string} on the page$`},
		func(ctx *common.TestSuiteContext) func(int, string) error {
			return func(expectedCount int, elementName string) error {
				elementCount := browser.GetElementCount(ctx.GetCurrentPage(), elementName)
				if elementCount != expectedCount {
					return fmt.Errorf("%d %s expected but %d %s found", expectedCount, elementName, elementCount, elementName)
				}
				return nil
			}
		},
		nil,
	)
}
