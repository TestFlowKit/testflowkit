package navigation

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (n navigation) openANewPrivateBrowserTab() stepbuilder.TestStep {
	return stepbuilder.NewStepWithNoVariables(
		[]string{"the user opens a new private browser tab"},
		func(ctx *stepbuilder.TestSuiteContext) func() error {
			return func() error {
				ctx.InitBrowser(true)
				return nil
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "opens a new private browser tab.",
			Variables:   nil,
			Example:     "Given the user opens a new private browser tab",
			Category:    shared.Navigation,
		},
	)
}
