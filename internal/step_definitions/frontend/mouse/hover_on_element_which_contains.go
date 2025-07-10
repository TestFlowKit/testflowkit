package mouse

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (s steps) hoverOnElementWhichContains() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the user hovers on {string} which contains {string}`},
		func(ctx context.Context, _, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			element, err := scenarioCtx.GetCurrentPageOnly().GetOneByTextContent(text)
			if err != nil {
				return ctx, err
			}
			err = element.Hover()
			return ctx, err
		},
		nil,
		stepbuilder.DocParams{
			Description: "hovers on an element which contains a specific text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to hover on.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user hovers on \"Submit button\" which contains \"Submit\"",
			Category: stepbuilder.Mouse,
		},
	)
}
