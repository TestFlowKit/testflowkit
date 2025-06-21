package visual

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) shouldSeeOnPageXElements() stepbuilder.TestStep {
	return stepbuilder.NewStepWithTwoVariables(
		[]string{`^the user should see {number} {string} on the page$`},
		func(ctx *scenario.Context) func(int, string) error {
			return func(expectedCount int, elementName string) error {
				elementCount := browser.GetElementCount(ctx.GetCurrentPage(), elementName)
				if elementCount != expectedCount {
					return fmt.Errorf("%d %s expected but %d %s found", expectedCount, elementName, elementCount, elementName)
				}
				return nil
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "checks if a specific number of elements are visible on the page.",
			Variables: []shared.StepVariable{
				{Name: "expectedCount", Description: "The expected number of elements.", Type: shared.DocVarTypeInt},
				{Name: "name", Description: "The logical name of the element to check.", Type: shared.DocVarTypeString},
			},
			Example:  "Then the user should see 3 buttons on the page",
			Category: shared.Visual,
		},
	)
}
