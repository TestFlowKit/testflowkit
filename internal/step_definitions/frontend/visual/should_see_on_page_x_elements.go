package visual

import (
	"context"
	"fmt"
	"strconv"
	"testflowkit/internal/browser"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) shouldSeeOnPageXElements() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the user should see {number} {string} elementson the page`},
		func(ctx context.Context, expectedCount, elementName string) (context.Context, error) {
			expectedCountInt, err := strconv.Atoi(expectedCount)
			if err != nil {
				return ctx, fmt.Errorf("invalid expected count: %s", expectedCount)
			}

			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage, pageName := scenarioCtx.GetCurrentPage()
			elementCount := browser.GetElementCount(currentPage, pageName, elementName)
			if elementCount != expectedCountInt {
				return ctx, fmt.Errorf("%d %s expected but %d %s found", expectedCountInt, elementName, elementCount, elementName)
			}
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a specific number of elements are visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "expectedCount", Description: "The expected number of elements.", Type: stepbuilder.VarTypeInt},
				{Name: "elementName", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the user should see 3 buttons elements on the page",
			Category: stepbuilder.Visual,
		},
	)
}
