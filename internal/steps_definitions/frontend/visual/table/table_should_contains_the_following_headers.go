package table

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (s steps) tableShouldContainsTheFollowingHeaders() stepbuilder.Step {
	example := `
	When the user should see a table with the following headers
	| Name | Age |
	`

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user should see a table with the following headers$`},
		func(ctx *scenario.Context) func(*godog.Table) error {
			return func(table *godog.Table) error {
				data, err := assistdog.NewDefault().ParseMap(table)
				if err != nil {
					return err
				}

				currentPage := ctx.GetCurrentPageOnly()
				_, err = getTableHeaderByCellsContent(currentPage, maps.Values(data))
				return err
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a table contains the following headers.",
			Variables: []stepbuilder.DocVariable{
				{Name: "table", Description: "The table containing the headers to check.", Type: stepbuilder.VarTypeTable},
			},
			Example:  example,
			Category: stepbuilder.Visual,
		},
	)
}
