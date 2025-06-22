package navigation

import (
	"errors"
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
	"time"
)

func (n navigation) switchToNewOpenedWindow() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"^the user switches to the newly opened window$"},
		func(ctx *scenario.Context) func() error {
			return func() error {
				initialPageCount := len(ctx.GetPages())
				logger.Info(fmt.Sprintf("Waiting for new window. Current window count: %d", initialPageCount))

				// TODO: refactor in order to use the Wait function
				startTime := time.Now()
				const duration = 6 * time.Minute
				for {
					if time.Since(startTime) > duration {
						return errors.New("no new window detected")
					}

					currentPageCount := len(ctx.GetPages())
					if currentPageCount > initialPageCount {
						logger.Info(fmt.Sprintf("New window detected! Page count increased from %d to %d",
							initialPageCount, currentPageCount))

						pages := ctx.GetPages()
						// In Rod, the most recently opened page is typically the first in the pages list
						newPage := pages[0]
						ctx.SetCurrentPage(newPage)

						return nil
					}

					const milliseconds = 100
					time.Sleep(milliseconds * time.Millisecond)
				}
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
