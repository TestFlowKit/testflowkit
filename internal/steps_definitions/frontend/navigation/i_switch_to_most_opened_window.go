package navigation

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (n navigation) iSwitchToMostRecentlyOpenedWindow() core.TestStep {
	return core.NewStepWithNoVariables(
		[]string{"^I switch to the most recently window opened$"},
		func(ctx *core.TestSuiteContext) func() error {
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
		func() core.ValidationErrors {
			return core.ValidationErrors{}
		},
		core.StepDefDocParams{
			Description: "switches to the most recently opened browser window.",
			Variables:   []shared.StepVariable{},
			Example:     "When I switch to the most recently window opened",
			Category:    shared.Navigation,
		},
	)
}
