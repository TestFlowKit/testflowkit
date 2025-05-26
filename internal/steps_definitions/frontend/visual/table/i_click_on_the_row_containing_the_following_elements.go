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

				logger.Info("try to get table ...")
				trs, _ := ctx.GetCurrentPage().GetAllBySelector("tr")

				idx := slices.IndexFunc(trs, func(elt common.Element) bool {
					textContent := elt.TextContent()
					for _, value := range values {
						if !strings.Contains(textContent, value) {
							return false
						}
					}
					return true
				})

				if idx == -1 {
					return errors.New("row not found with the following values: " + strings.Join(values, ", "))
				}

				trs[idx].Click()
				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "clicks on the row containing the following elements.",
			Variables: []shared.StepVariable{
				{Name: "Map", Description: "The map containing the element to click on.", Type: shared.DocVarTypeTable},
			},
			Example:  example,
			Category: shared.Visual,
		},
	)
}
