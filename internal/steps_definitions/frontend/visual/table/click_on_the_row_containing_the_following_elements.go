package table

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

// TODO: click on cell instead of row
func (s steps) clickOnTheRowContainingTheFollowingElements() stepbuilder.TestStep {
	const example = `
	When the user clicks on the row containing the following elements
	| Name | Age |
	| John | 30  |
	`
	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user clicks on the row containing the following elements$`},
		func(ctx *scenario.Context) func(*godog.Table) error {
			return func(table *godog.Table) error {
				data, parseErr := assistdog.NewDefault().ParseSlice(table)
				if parseErr != nil {
					return parseErr
				}

				for _, rowDetails := range data {
					element, gerRowErr := getTableRowByCellsContent(ctx.GetCurrentPage(), maps.Values(rowDetails))
					if gerRowErr != nil {
						return gerRowErr
					}

					clickErr := element.Click()
					if clickErr != nil {
						return clickErr
					}
				}

				return nil
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "clicks on the row containing the following elements.",
			Variables: []shared.StepVariable{
				{Name: "table", Description: "The table containing the elements to click on.", Type: shared.DocVarTypeTable},
			},
			Example:  example,
			Category: shared.Visual,
		},
	)
}
