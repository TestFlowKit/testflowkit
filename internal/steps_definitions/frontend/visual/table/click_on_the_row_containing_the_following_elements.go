package table

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

// TODO: click on cell instead of row
func (s steps) clickOnTheRowContainingTheFollowingElements() stepbuilder.Step {
	const example = `
	When the user clicks on the row containing the following elements
	| Name | Age |
	| John | 30  |
	`
	return stepbuilder.NewWithOneVariable[*godog.Table](
		[]string{`^the user clicks on the row containing the following elements$`},
		func(ctx *scenario.Context) func(*godog.Table) error {
			return func(table *godog.Table) error {
				data, err := assistdog.NewDefault().ParseSlice(table)
				if err != nil {
					return err
				}

				for _, rowDetails := range data {
					element, getRowErr := getTableRowByCellsContent(ctx.GetCurrentPageOnly(), maps.Values(rowDetails))
					if getRowErr != nil {
						return getRowErr
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
		stepbuilder.DocParams{
			Description: "clicks on the row containing the following elements.",
			Variables: []stepbuilder.DocVariable{
				{Name: "table", Description: "The table containing the elements to click on.", Type: stepbuilder.VarTypeTable},
			},
			Example:  example,
			Category: stepbuilder.Visual,
		},
	)
}
