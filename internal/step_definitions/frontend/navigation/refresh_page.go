package navigation

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) refreshPage() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user refresh the page"},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage := scenarioCtx.GetCurrentPageOnly()
			currentPage.Refresh()
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "refreshes the current page.",
			Variables:   nil,
			Example:     "When the user refreshes the page",
			Category:    stepbuilder.Navigation,
		},
	)
}
