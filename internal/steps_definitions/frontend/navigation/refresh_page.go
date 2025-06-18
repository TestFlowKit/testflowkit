package navigation

import (
	"errors"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (n navigation) refreshPage() stepbuilder.TestStep {
	return stepbuilder.NewStepWithNoVariables(
		[]string{"the user refresh the page"},
		func(ctx *stepbuilder.TestSuiteContext) func() error {
			return func() error {
				if ctx.GetCurrentPage() == nil {
					return errors.New("no page opened")
				}
				ctx.GetCurrentPage().Refresh()
				return nil
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "refreshes the current page.",
			Variables:   nil,
			Example:     "When the user refreshes the page",
			Category:    shared.Navigation,
		},
	)
}
