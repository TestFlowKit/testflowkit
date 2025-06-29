package table

import (
	"context"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (steps) tableShouldContainsTheFollowingHeaders() stepbuilder.Step {
	example := `
	When the user should see a table with the following headers
	| Name | Age |
	`

	return stepbuilder.NewWithOneVariable(
		[]string{`the user should see a table with the following headers`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			data, err := assistdog.NewDefault().ParseMap(table)
			if err != nil {
				return ctx, err
			}

			currentPage := scenarioCtx.GetCurrentPageOnly()
			_, err = getTableHeaderByCellsContent(currentPage, maps.Values(data))
			return ctx, err
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
