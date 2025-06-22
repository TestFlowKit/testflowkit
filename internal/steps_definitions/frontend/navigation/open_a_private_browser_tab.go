package navigation

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (n navigation) openANewPrivateBrowserTab() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user opens a new private browser tab"},
		func(ctx *scenario.Context) func() error {
			return func() error {
				ctx.InitBrowser(true)
				return nil
			}
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
