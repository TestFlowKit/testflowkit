package table

import (
	"etoolse/internal/steps_definitions/core"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (s steps) tableShouldContainsTheFollowingHeaders() core.TestStep {
	return core.NewStepWithOneVariable(
		[]string{`^I should see a table with the following headers$`},
		func(ctx *core.TestSuiteContext) func(*godog.Table) error {
			return func(table *godog.Table) error {
				data, err := assistdog.NewDefault().ParseMap(table)
				if err != nil {
					return err
				}

				_, err = getTableHeaderByCellsContent(ctx.GetCurrentPage(), maps.Values(data))
				return err
			}
		},
		nil,
	)
}
