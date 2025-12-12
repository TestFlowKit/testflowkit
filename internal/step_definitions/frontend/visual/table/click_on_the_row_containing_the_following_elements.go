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

// TODO: click on cell instead of row.
func (steps) clickOnTheRowContainingTheFollowingElements() stepbuilder.Step {
	const example = `
	When the user clicks on the row containing the following elements
	| Name | Age |
	| John | 30  |
	`
	return stepbuilder.NewWithOneVariable(
		[]string{`the user clicks on the row containing the following elements`},
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

			page, errPage := scenarioCtx.GetCurrentPageOnly()
			if errPage != nil {
				return ctx, errPage
			}

			for _, rowDetails := range parsedData {
				values := slices.Sorted(maps.Values(rowDetails))
				element, getRowErr := getTableRowByCellsContent(page, values)
				if getRowErr != nil {
					return ctx, getRowErr
				}

				clickErr := element.Click()
				if clickErr != nil {
					return ctx, clickErr
				}
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "clicks on the row containing the following elements.",
			Variables: []stepbuilder.DocVariable{
				{Name: "table", Description: "The table containing the elements to click on.", Type: stepbuilder.VarTypeTable},
			},
			Example:  example,
			Category: stepbuilder.Visual,
		},
	)
}
