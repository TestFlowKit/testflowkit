package navigation

import (
	"errors"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (n navigation) theUserNavigateBack() stepbuilder.TestStep {
	return stepbuilder.NewStepWithNoVariables(
		[]string{"the user navigate back"},
		func(ctx *stepbuilder.TestSuiteContext) func() error {
			return func() error {
				if ctx.GetCurrentPage() == nil {
					return errors.New("no page opened")
				}
				ctx.GetCurrentPage().Back()
				return nil
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "navigates back to the previous page.",
			Variables:   nil,
			Example:     "Given the user navigate back",
			Category:    shared.Navigation,
		},
	)
}
