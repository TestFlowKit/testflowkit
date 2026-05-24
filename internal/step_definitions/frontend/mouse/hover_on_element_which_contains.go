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
			currentPage, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}
			element, err := currentPage.GetOneByTextContent(text)
			if err != nil {
				return ctx, err
			}
			err = element.Hover()
			return ctx, err
		},
		nil,
		elementWhichContainsDocParams(
			docDescElementHoverOn,
			"hovers on an element which contains a specific text.",
			`When the user hovers on "Submit button" which contains "Submit"`,
		),
	)
}
