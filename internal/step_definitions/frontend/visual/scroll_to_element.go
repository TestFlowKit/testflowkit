package visual

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) scrollToElement() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user scrolls to the {string} element`},
		func(ctx context.Context, elementName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)

			element, err := scenarioCtx.GetHTMLElementByLabel(elementName + "_element")

			if err != nil {
				return ctx, err
			}
			errScroll := element.ScrollIntoView()
			if errScroll != nil {
				return ctx, fmt.Errorf("failed to scroll to elementName '%s': %w", elementName, errScroll)
			}
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "scrolls to a specific element on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "elementName", Description: "The name of the element to scroll to.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user scrolls to the submit button element",
			Category: stepbuilder.Visual,
		},
	)
}
