package navigation

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (n navigation) iSwitchToOriginalWindow() core.TestStep {
	return core.NewStepNoVariable(
		[]string{"^I switch back to the original window$"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				pages := ctx.GetPages()
				
				if len(pages) < 2 {
					return fmt.Errorf("only one window is open, no original window to switch back to")
				}
				
				// In most cases, the original window is the first one in the list of pages
				originalPage := pages[0]
				
				// Focus on the original page
				originalPage.Focus()
				
				// Wait for the page to load completely
				originalPage.WaitLoading()
				
				// CRITICAL: Update the current page in the context
				ctx.SetCurrentPage(originalPage)
				
				logger.Info(fmt.Sprintf("Switched back to original window with URL: %s", originalPage.GetInfo().URL))
				
				return nil
			}
		},
		func() core.ValidationErrors {
			return core.ValidationErrors{}
		},
		core.StepDefDocParams{
			Description: "switches back to the original browser window (usually the first window).",
			Variables:   []shared.StepVariable{},
			Example:     "When I switch back to the original window",
			Category:    shared.Navigation,
		},
	)
} 