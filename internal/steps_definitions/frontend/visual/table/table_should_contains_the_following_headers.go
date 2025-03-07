package table

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (s steps) tableShouldContainsTheFollowingHeaders() core.TestStep {
	example := `
	When I should see a table with the following headers
	| Name | Age |
	`

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
		core.StepDefDocParams{
			Description: "checks if a table contains the following headers.",
			Variables: []shared.StepVariable{
				{Name: "table", Description: "The table containing the headers to check.", Type: shared.DocVarTypeTable},
			},
			Example:  example,
			Category: shared.Visual,
		},
	)
}
