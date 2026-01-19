package navigation

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) openANewBrowserTab() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user opens a new browser tab"},
		func(ctx context.Context) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			scenarioCtx.InitBrowser(false)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "opens a new browser tab.",
			Variables:   nil,
			Example:     "Given the user opens a new browser tab",
			Categories:  []stepbuilder.StepCategory{stepbuilder.Navigation}},
	)
}
