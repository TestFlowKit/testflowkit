package table

import (
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

func (s steps) shouldSeeRowContainingTheFollowingElements() core.TestStep {
	example := `
	When the user should see a row containing the following elements
	| Name | Age |
	| John | 30  |
	`
	return core.NewStepWithOneVariable[*godog.Table](
		[]string{`^the user should see a row containing the following elements$`},
		func(ctx *core.TestSuiteContext) func(*godog.Table) error {
			return func(table *godog.Table) error {
				data, err := assistdog.NewDefault().ParseSlice(table)
				if err != nil {
					return err
				}

				for _, rowDetails := range data {
					_, getRowErr := getTableRowByCellsContent(ctx.GetCurrentPage(), maps.Values(rowDetails))
					if getRowErr != nil {
						return getRowErr
					}
				}

				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "checks if a row containing the following elements is visible in the table.",
			Variables: []shared.StepVariable{
				{Name: "table", Description: "The table containing the elements to check.", Type: shared.DocVarTypeTable},
			},
			Example:  example,
			Category: shared.Visual,
		},
	)
}
