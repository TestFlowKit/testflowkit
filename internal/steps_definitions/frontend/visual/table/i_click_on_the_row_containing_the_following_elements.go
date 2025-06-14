package table

import (
	"errors"
	"slices"
	"strings"
	"testflowkit/internal/browser/common"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/pkg/logger"
	"testflowkit/shared"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"golang.org/x/exp/maps"
)

// TODO: click on cell instead of row
func (s steps) iClickOnTheRowContainingTheFollowingElements() core.TestStep {
	const example = `
	When I click on the row containing the following elements
	| Name | John |
	| Age | 30  |
	`
	return core.NewStepWithOneVariable(
		[]string{`^I click on the row containing the following elements$`},
		func(ctx *core.TestSuiteContext) func(*godog.Table) error {
			return func(table *godog.Table) error {
				data, parseErr := assistdog.NewDefault().ParseMap(table)
				if parseErr != nil {
					return parseErr
				}

				values := maps.Values(data)

				row, err := getTableRowByCellsContent(ctx.GetCurrentPage(), values)
				if err != nil {
					return err
				}

				clickErr := row.Click()
				if clickErr != nil {
					return clickErr
				}
				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "clicks on the row containing the following elements.",
			Variables: []shared.StepVariable{
				{Name: "Map", Description: "The map containing the row content", Type: shared.DocVarTypeTable},
			},
			Example:  example,
			Category: shared.Visual,
		},
	)
}
