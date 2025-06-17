package navigation

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (n navigation) openNewBrowserTab() core.TestStep {
	return core.NewStepWithNoVariables(
		[]string{"the user opens a new browser tab"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				ctx.InitBrowser(false)
				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "opens a new browser tab.",
			Variables:   nil,
			Example:     "Given the user opens a new browser tab",
			Category:    shared.Navigation,
		},
	)
}
