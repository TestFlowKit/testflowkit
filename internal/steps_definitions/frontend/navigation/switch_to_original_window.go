package navigation

import (
	"errors"
	"fmt"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (n navigation) switchToOriginalWindow() stepbuilder.TestStep {
	return stepbuilder.NewStepWithNoVariables(
		[]string{"^the user switches back to the original window$"},
		func(ctx *stepbuilder.TestSuiteContext) func() error {
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
		func() stepbuilder.ValidationErrors {
			return stepbuilder.ValidationErrors{}
		},
		stepbuilder.StepDefDocParams{
			Description: "switches back to the original browser window (usually the first window).",
			Variables:   []shared.StepVariable{},
			Example:     "When the user switches back to the original window",
			Category:    shared.Navigation,
		},
	)
}
