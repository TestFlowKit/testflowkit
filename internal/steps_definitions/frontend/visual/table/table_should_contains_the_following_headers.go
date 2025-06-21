package table

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (s steps) tableShouldContainsTheFollowingHeaders() stepbuilder.TestStep {
	example := `
	When the user should see a table with the following headers
	| Name | Age |
	`

	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user should see a table with the following headers$`},
		func(ctx *scenario.Context) func(*godog.Table) error {
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
		stepbuilder.StepDefDocParams{
			Description: "checks if a table contains the following headers.",
			Variables: []shared.StepVariable{
				{Name: "table", Description: "The table containing the headers to check.", Type: shared.DocVarTypeTable},
			},
			Example:  example,
			Category: shared.Visual,
		},
	)
}
