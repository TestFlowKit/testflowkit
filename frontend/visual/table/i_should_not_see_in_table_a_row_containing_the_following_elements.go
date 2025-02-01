package table

import (
	"errors"
	"etoolse/frontend/common"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (s steps) iShouldNotSeeRowContainingTheFollowingElements() common.FrontStep {
	return common.NewStepWithOneVariable[*godog.Table](
		[]string{`^I should not see a row containing the following elements$`},
		func(ctx *common.TestSuiteContext) func(*godog.Table) error {
			return func(table *godog.Table) error {
				data, err := assistdog.NewDefault().ParseSlice(table)
				if err != nil {
					return err
				}

				for _, rowDetails := range data {
					_, err = getTableRowByCellsContent(ctx.GetCurrentPage(), maps.Values(rowDetails))
					if err == nil {
						// TODO: better log
						return errors.New("row found")
					}
				}
				return nil
			}
		},
		nil,
	)
}
