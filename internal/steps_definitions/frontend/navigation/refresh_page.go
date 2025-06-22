package navigation

import (
	"errors"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (n navigation) refreshPage() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{"the user refresh the page"},
		func(ctx *scenario.Context) func() error {
			return func() error {
				if ctx.GetCurrentPage() == nil {
					return errors.New("no page opened")
				}
				ctx.GetCurrentPage().Refresh()
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "refreshes the current page.",
			Variables:   nil,
			Example:     "When the user refreshes the page",
			Category:    stepbuilder.Navigation,
		},
	)
}
