package navigation

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) refreshPage() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`the user refreshes the page`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}
			currentPage.Refresh()
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "refreshes the current page.",
			Example:     "When the user refreshes the page",
			Categories:  []stepbuilder.StepCategory{stepbuilder.Navigation},
		},
	)
}
