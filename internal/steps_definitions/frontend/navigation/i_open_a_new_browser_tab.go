package navigation

import (
	"etoolse/internal/steps_definitions/core"
	"etoolse/shared"
)

func (n navigation) iOpenNewBrowserTab() core.TestStep {
	return core.NewStepWithoutVariables(
		[]string{"I open a new browser tab"},
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
			Example:     "Given I open a new browser tab",
			Category:    shared.Navigation,
		},
	)
}
