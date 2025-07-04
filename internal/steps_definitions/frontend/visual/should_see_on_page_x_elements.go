package visual

import (
	"context"
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) shouldSeeOnPageXElements() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the user should see {number} {string} on the page`},
		func(ctx context.Context, expectedCount int, elementName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage, pageName := scenarioCtx.GetCurrentPage()
			elementCount := browser.GetElementCount(currentPage, pageName, elementName)
			if elementCount != expectedCount {
				return ctx, fmt.Errorf("%d %s expected but %d %s found", expectedCount, elementName, elementCount, elementName)
			}
			return ctx, nil
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
