package navigation

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (n navigation) iRefreshPage() core.TestStep {
	return core.NewStepWithoutVariables(
		[]string{"I refresh the page"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				ctx.GetCurrentPage().Refresh()
				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "refreshes the current page.",
			Variables:   nil,
			Example:     "When I refresh the page",
			Category:    shared.Navigation,
		},
	)
}
