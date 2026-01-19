package table

import (
	"context"
	"errors"
	"maps"
	"slices"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (steps) shouldNotSeeRowContainingTheFollowingElements() stepbuilder.Step {
	example := `
	When the user should not see a row containing the following elements
	| Name | Age |
	| John | 30  |
	`

	return stepbuilder.NewWithOneVariable(
		[]string{`the user should not see a row containing the following elements`},
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

			currentPage, errPage := scenarioCtx.GetCurrentPageOnly()
			if errPage != nil {
				return ctx, errPage
			}

			for _, rowDetails := range parsedData {
				values := slices.Sorted(maps.Values(rowDetails))

				_, err = getTableRowByCellsContent(currentPage, values)
				if err == nil {
					return ctx, errors.New("row containing the specified elements was found but should not be visible")
				}
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a row containing the following elements is not visible in the table.",
			Variables: []stepbuilder.DocVariable{
				{Name: "table", Description: "The table containing the elements to check.", Type: stepbuilder.VarTypeTable},
			},
			Example:    example,
			Categories: []stepbuilder.StepCategory{stepbuilder.Visual},
		},
	)
}
