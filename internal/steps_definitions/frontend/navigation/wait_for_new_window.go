package navigation

import (
	"context"
	"errors"
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
	"time"
)

func (steps) switchToNewOpenedWindow() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"^the user switches to the newly opened window$"},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			initialPageCount := len(scenarioCtx.GetPages())
			logger.Info(fmt.Sprintf("Waiting for new window. Current window count: %d", initialPageCount))

			// TODO: refactor in order to use the Wait function
			startTime := time.Now()
			const duration = 6 * time.Minute
			for {
				if time.Since(startTime) > duration {
					return ctx, errors.New("no new window detected")
				}

				currentPageCount := len(scenarioCtx.GetPages())
				if currentPageCount > initialPageCount {
					logger.Info(fmt.Sprintf("New window detected! Page count increased from %d to %d",
						initialPageCount, currentPageCount))

					pages := scenarioCtx.GetPages()
					// In Rod, the most recently opened page is typically the first in the pages list
					newPage := pages[0]
					if err := scenarioCtx.SetCurrentPage(newPage); err != nil {
						return ctx, fmt.Errorf("failed to set current page: %w", err)
					}

					return ctx, nil
				}

				const milliseconds = 100
				time.Sleep(milliseconds * time.Millisecond)
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "switches to the newly opened browser window.",
			Example:     "When the user switches to the newly opened window",
			Category:    stepbuilder.Navigation,
		},
	)
}
