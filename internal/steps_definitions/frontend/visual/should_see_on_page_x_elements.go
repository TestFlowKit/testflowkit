package visual

import (
	"context"
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) shouldSeeOnPageXElements() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the user should see {number} {string} on the page`},
		func(ctx context.Context, count int, elementName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			elements, err := scenarioCtx.GetCurrentPageOnly().GetAllBySelector(elementName)
			if err != nil {
				return ctx, err
			}
			if len(elements) != count {
				return ctx, fmt.Errorf("expected %d %s elements, but found %d", count, elementName, len(elements))
			}
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if the specified number of elements are present on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "count", Description: "The expected number of elements.", Type: stepbuilder.VarTypeInt},
				{Name: "elementName", Description: "The name of the element to count.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the user should see 3 buttons on the page",
			Category: stepbuilder.Visual,
		},
	)
}
