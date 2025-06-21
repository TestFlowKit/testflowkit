package navigation

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (n navigation) switchToMostOpenedWindow() stepbuilder.TestStep {
	return stepbuilder.NewStepWithNoVariables(
		[]string{"^the user switches to the most recently window opened$"},
		func(ctx *scenario.Context) func() error {
			return func() error {
				pages := ctx.GetPages()

				const minPages = 2
				if len(pages) < minPages {
					return fmt.Errorf("no additional windows found to switch to (only %d window open)", len(pages))
				}

				// In Rod, the most recently opened page is typically the first in the pages list
				newPage := pages[0]
				ctx.SetCurrentPage(newPage)

				logger.Info(fmt.Sprintf("Switched to new window with URL: %s", newPage.GetInfo().URL))

				return nil
			}
		},
		func() stepbuilder.ValidationErrors {
			return stepbuilder.ValidationErrors{}
		},
		stepbuilder.StepDefDocParams{
			Description: "switches to the most recently opened browser window.",
			Variables:   []shared.StepVariable{},
			Example:     "When the user switches to the most recently window opened",
			Category:    shared.Navigation,
		},
	)
}
