package navigation

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (n navigation) openANewBrowserTab() stepbuilder.TestStep {
	return stepbuilder.NewStepWithNoVariables(
		[]string{"the user opens a new browser tab"},
		func(ctx *scenario.Context) func() error {
			return func() error {
				ctx.InitBrowser(false)
				return nil
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "opens a new browser tab.",
			Variables:   nil,
			Example:     "Given the user opens a new browser tab",
			Category:    shared.Navigation,
		},
	)
}
