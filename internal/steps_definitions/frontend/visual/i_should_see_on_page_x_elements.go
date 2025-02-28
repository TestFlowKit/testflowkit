package visual

import (
	"etoolse/internal/browser"
	"etoolse/internal/steps_definitions/core"
	"etoolse/shared"
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
		core.StepDefDocParams{
			Description: "checks if a specific number of elements are visible on the page.",
			Variables: []shared.StepVariable{
				{Name: "expectedCount", Description: "The expected number of elements.", Type: shared.DocVarTypeInt},
				{Name: "elementName", Description: "The name of the element to check.", Type: shared.DocVarTypeString},
			},
			Example:  "Then I should see 3 buttons on the page",
			Category: shared.Visual,
		},
	)
}
