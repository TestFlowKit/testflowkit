package table

import (
	"etoolse/frontend/common"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (s steps) iShouldSeeRowContainingTheFollowingElements() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I should see a row containing the following elements$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(table *godog.Table) error {
				data, err := assistdog.NewDefault().ParseSlice(table)
				if err != nil {
					return err
				}

				for _, rowDetails := range data {
					_, getRowErr := getTableRowByCellsContent(ctx.GetCurrentPage(), maps.Values(rowDetails))
					if getRowErr != nil {
						return getRowErr
					}
				}

				return nil
			}
		},
	}
}
