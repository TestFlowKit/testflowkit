package navigation

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (n navigation) openANewBrowserTab() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user opens a new browser tab"},
		func(ctx *scenario.Context) func() error {
			return func() error {
				ctx.InitBrowser(false)
				return nil
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
