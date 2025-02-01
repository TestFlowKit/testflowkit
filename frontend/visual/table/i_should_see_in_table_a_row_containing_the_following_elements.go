package table

import (
	"etoolse/frontend/common"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (s steps) iShouldSeeRowContainingTheFollowingElements() common.FrontStep {
	return common.NewStepWithOneVariable[*godog.Table](
		[]string{`^I should see a row containing the following elements$`},
		func(ctx *common.TestSuiteContext) func(*godog.Table) error {
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
		nil,
	)
}
