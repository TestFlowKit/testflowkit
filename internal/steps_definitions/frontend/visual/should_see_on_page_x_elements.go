package visual

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) shouldSeeOnPageXElements() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user should see {number} {string} on the page$`},
		func(ctx *scenario.Context) func(int, string) error {
			return func(expectedCount int, elementName string) error {
				currentPage, pageName := ctx.GetCurrentPage()
				elementCount := browser.GetElementCount(currentPage, pageName, elementName)
				if elementCount != expectedCount {
					return fmt.Errorf("%d %s expected but %d %s found", expectedCount, elementName, elementCount, elementName)
				}
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a specific number of elements are visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "expectedCount", Description: "The expected number of elements.", Type: stepbuilder.VarTypeInt},
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the user should see 3 buttons on the page",
			Category: stepbuilder.Visual,
		},
	)
}
