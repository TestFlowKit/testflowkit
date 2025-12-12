package navigation

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) switchToMostOpenedWindow() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user switches to the most recently window opened"},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			pages := scenarioCtx.GetPages()
			const minPages = 2
			if len(pages) < minPages {
				return ctx, fmt.Errorf("no additional windows found to switch to (only %d window open)", len(pages))
			}
			newPage := pages[0]
			if err := scenarioCtx.SetCurrentPage(newPage); err != nil {
				return ctx, fmt.Errorf("failed to set current page: %w", err)
			}
			logger.Info("Switched to new window with URL: " + newPage.GetInfo().URL)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "switches to the most recently opened browser window.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     "When the user switches to the most recently window opened",
			Category:    stepbuilder.Navigation,
		},
	)
}
