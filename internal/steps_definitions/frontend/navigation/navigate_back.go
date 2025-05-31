package navigation

import (
	"errors"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (n navigation) theUserNavigateBack() core.TestStep {
	return core.NewStepWithNoVariables(
		[]string{"the user navigate back"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				if ctx.GetCurrentPage() == nil {
					return errors.New("no page opened")
				}
				ctx.GetCurrentPage().Back()
				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "navigates back to the previous page.",
			Variables:   nil,
			Example:     "Given the user navigate back",
			Category:    shared.Navigation,
		},
	)
}
