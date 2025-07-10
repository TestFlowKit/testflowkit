package navigation

import (
	"context"
	"errors"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) switchToOriginalWindow() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user switches back to the original window"},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			pages := scenarioCtx.GetPages()

			const minPages = 2
			if len(pages) < minPages {
				return ctx, errors.New("only one window is open, no original window to switch back to")
			}

			originalPage := pages[len(pages)-1]

			originalPage.Focus()

			originalPage.WaitLoading()

			if err := scenarioCtx.SetCurrentPage(originalPage); err != nil {
				return ctx, fmt.Errorf("failed to set current page: %w", err)
			}

			logger.Info(fmt.Sprintf("Switched back to original window with URL: %s", originalPage.GetInfo().URL))

			return ctx, nil
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
