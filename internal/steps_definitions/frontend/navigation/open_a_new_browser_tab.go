package navigation

import (
	"context"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) openANewBrowserTab() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user opens a new browser tab"},
		func(scenarioCtx *scenario.Context) func(context.Context) (context.Context, error) {
			return func(ctx context.Context) (context.Context, error) {
				scenarioCtx.InitBrowser(false)
				return ctx, nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "opens a new browser tab.",
			Variables:   nil,
			Example:     "Given the user opens a new browser tab",
			Category:    stepbuilder.Navigation,
		},
	)
}
