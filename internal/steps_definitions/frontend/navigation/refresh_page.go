package navigation

import (
	"errors"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (n navigation) refreshPage() core.TestStep {
	return core.NewStepWithNoVariables(
		[]string{"the user refresh the page"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				if ctx.GetCurrentPage() == nil {
					return errors.New("no page opened")
				}
				ctx.GetCurrentPage().Refresh()
				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "refreshes the current page.",
			Variables:   nil,
			Example:     "When the user refreshes the page",
			Category:    shared.Navigation,
		},
	)
}
