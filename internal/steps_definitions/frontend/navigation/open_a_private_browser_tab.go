package navigation

import (
	"context"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) openANewPrivateBrowserTab() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user opens a new private browser tab"},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			scenarioCtx.InitBrowser(true)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "opens a new private browser tab.",
			Variables:   nil,
			Example:     "Given the user opens a new private browser tab",
			Category:    stepbuilder.Navigation,
		},
	)
}
