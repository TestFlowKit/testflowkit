package navigation

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (n navigation) iOpenNewPrivateBrowserTab() core.TestStep {
	return core.NewStepWithoutVariables(
		[]string{"I open a new private browser tab"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				ctx.InitBrowser(true)
				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "opens a new private browser tab.",
			Variables:   nil,
			Example:     "Given I open a new private browser tab",
			Category:    shared.Navigation,
		},
	)
}
