package navigation

import (
	"context"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) theUserNavigateBack() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`the user navigates back`},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage := scenarioCtx.GetCurrentPageOnly()
			currentPage.Back()
			return ctx, nil

		},
		nil,
		stepbuilder.DocParams{
			Description: "navigates back in the browser history.",
			Variables:   nil,
			Example:     "When the user navigates back",
			Category:    stepbuilder.Navigation,
		},
	)
}
