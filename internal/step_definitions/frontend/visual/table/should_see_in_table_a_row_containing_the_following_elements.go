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

func (steps) shouldSeeRowContainingTheFollowingElements() stepbuilder.Step {
	example := `
	When the user should see a row containing the following elements
	| Name | Age |
	| John | 30  |
	`
	return stepbuilder.NewWithOneVariable(
		[]string{`the user should see a row containing the following elements`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			data, err := assistdog.NewDefault().ParseSlice(table)
			if err != nil {
				return ctx, err
			}

			parsedData, err := scenario.ReplaceVariablesInArray(scenarioCtx, data)
			if err != nil {
				return ctx, err
			}

			currentPage := scenarioCtx.GetCurrentPageOnly()
			for _, rowDetails := range parsedData {
				values := slices.Sorted(maps.Values(rowDetails))
				_, getRowErr := getTableRowByCellsContent(currentPage, values)
				if getRowErr != nil {
					return ctx, getRowErr
				}
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a row containing the following elements is visible in the table.",
			Variables: []stepbuilder.DocVariable{
				{Name: "table", Description: "The table containing the elements to check.", Type: stepbuilder.VarTypeTable},
			},
			Example:  example,
			Category: stepbuilder.Visual,
		},
	)
}
