package mouse

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (s steps) rightClickOnElementWhichContains() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the user right clicks on {string} which contains "{string}"`},
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
			err = element.RightClick()
			return ctx, err
		},
		nil,
		elementWhichContainsDocParams(
			docDescElementRightClickOn,
			"right clicks on an element which contains a specific text.",
			`When the user right clicks on "Submit button" which contains "Submit"`,
		),
	)
}
