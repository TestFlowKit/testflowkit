package table

import (
	"etoolse/internal/steps_definitions/core"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

// TODO: click on cell instead of row
func (s steps) iClickOnTheRowContainingTheFollowingElements() core.TestStep {
	return core.NewStepWithOneVariable(
		[]string{`^I click on the row containing the following elements$`},
		func(ctx *core.TestSuiteContext) func(*godog.Table) error {
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
	)
}
