package navigation

import (
	"errors"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (n navigation) theUserNavigateBack() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user navigate back"},
		func(ctx *scenario.Context) func() error {
			return func() error {
				if ctx.GetCurrentPage() == nil {
					return errors.New("no page opened")
				}
				ctx.GetCurrentPage().Back()
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "navigates back to the previous page.",
			Variables:   nil,
			Example:     "Given the user navigate back",
			Category:    stepbuilder.Navigation,
		},
	)
}
