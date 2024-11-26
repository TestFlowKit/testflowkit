package table

import (
	"cucumber/frontend/common"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (s steps) tableShouldContainsTheFollowingHeaders() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I should see a table with the following headers$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(table *godog.Table) error {
				data, err := assistdog.NewDefault().ParseMap(table)
				if err != nil {
					return err
				}

				_, err = getTableHeaderByCellsContent(ctx.GetCurrentPage(), maps.Values(data))
				return err
			}
		},
	}
}
