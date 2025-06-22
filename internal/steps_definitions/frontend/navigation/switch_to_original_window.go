package navigation

import (
	"errors"
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (n navigation) switchToOriginalWindow() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"^the user switches back to the original window$"},
		func(ctx *scenario.Context) func() error {
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
		stepbuilder.DocParams{
			Description: "switches back to the original browser window (usually the first window).",
			Variables:   []stepbuilder.DocVariable{},
			Example:     "When the user switches back to the original window",
			Category:    stepbuilder.Navigation,
		},
	)
}
