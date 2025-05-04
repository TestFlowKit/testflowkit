package navigation

import (
	"errors"
	"fmt"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (n navigation) iSwitchToOriginalWindow() core.TestStep {
	return core.NewStepWithNoVariables(
		[]string{"^I switch back to the original window$"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				pages := ctx.GetPages()

				const minPages = 2
				if len(pages) < minPages {
					return errors.New("only one window is open, no original window to switch back to")
				}

				originalPage := pages[len(pages)-1]

				originalPage.Focus()

				originalPage.WaitLoading()

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
