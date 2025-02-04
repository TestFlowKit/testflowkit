package visual

import (
	"etoolse/internal/browser"
	"etoolse/internal/steps_definitions/core"
	"fmt"
)

func (s steps) iShouldSeeOnPageXElements() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I should see {number} {string} on the page$`},
		func(ctx *core.TestSuiteContext) func(int, string) error {
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
