package navigation

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (n navigation) iNavigateBack() core.TestStep {
	return core.NewStepWithNoVariables(
		[]string{"I navigate back"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				ctx.GetCurrentPage().Back()
				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "navigates back to the previous page.",
			Variables:   nil,
			Example:     "Given I navigate back",
			Category:    shared.Navigation,
		},
	)
}
