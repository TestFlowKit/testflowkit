package table

import (
	"context"
	"maps"
	"slices"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
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

			parsedData, err := scenario.ReplaceVariablesInMap(scenarioCtx, data)
			if err != nil {
				return ctx, err
			}

			currentPage := scenarioCtx.GetCurrentPageOnly()
			_, err = getTableHeaderByCellsContent(currentPage, slices.Sorted(maps.Values(parsedData)))
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
