package table

import (
	"context"
	"errors"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (steps) shouldNotSeeRowContainingTheFollowingElements() stepbuilder.Step {
	example := `
	When the user should not see a row containing the following elements
	| Name | Age |
	| John | 30  |
	`

	return stepbuilder.NewWithOneVariable[*godog.Table](
		[]string{`^the user should not see a row containing the following elements$`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			data, err := assistdog.NewDefault().ParseSlice(table)
			if err != nil {
				return ctx, err
			}

			currentPage := scenarioCtx.GetCurrentPageOnly()
			for _, rowDetails := range data {
				_, err = getTableRowByCellsContent(currentPage, maps.Values(rowDetails))
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
			Example:  example,
			Category: stepbuilder.Visual,
		},
	)
}
