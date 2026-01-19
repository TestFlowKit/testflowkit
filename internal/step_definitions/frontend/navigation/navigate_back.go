package navigation

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) navigateBack() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`the user navigates back`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}
			currentPage.Back()
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "navigates back to the previous page in browser history.",
			Example:     "When the user navigates back",
			Categories:  []stepbuilder.StepCategory{stepbuilder.Navigation}},
	)
}
