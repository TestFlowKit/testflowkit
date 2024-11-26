package table

import (
	"cucumber/frontend/common"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

// TODO: click on cell instead of row
func (s steps) iClickOnTheRowContainingTheFollowingElements() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I click on the row containing the following elements$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
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
	}
}
