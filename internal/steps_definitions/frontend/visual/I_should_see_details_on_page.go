package visual

import (
	"errors"
	"etoolse/internal/steps_definitions/core"
	"etoolse/shared"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (s steps) iShouldSeeDetailsOnPage() core.TestStep {
	definition := func(ctx *core.TestSuiteContext) func(string, *godog.Table) error {
		return func(elementName string, table *godog.Table) error {
			data, parseErr := assistdog.NewDefault().ParseMap(table)
			if parseErr != nil {
				return errors.New("details malformed please go to the doc")
			}

			var errMsgs []string
			for name, value := range data {
				xPath := fmt.Sprintf("//*[contains(text(),\"%s\")]", value)
				elt, err := ctx.GetCurrentPage().GetOneByXPath(xPath)
				if err != nil {
					errMsgs = append(errMsgs, fmt.Sprintf("%s %s not found", elementName, name))
					continue
				}

				if !elt.IsVisible() {
					errMsgs = append(errMsgs, fmt.Sprintf("%s %s is found but is no visible", elementName, name))
				}
			}

			if len(errMsgs) > 0 {
				return errors.New(strings.Join(errMsgs, "\n"))
			}

			return nil
		}
	}

	return core.NewStepWithTwoVariables(
		[]string{`^I should see "{string}" details on the page$`},
		definition,
		nil,
		core.StepDefDocParams{
			Description: "checks if the details are visible on the page.",
			Variables: []shared.StepVariable{
				{Name: "elementName", Description: "The name of the element to check.", Type: shared.DocVarTypeString},
				{Name: "table", Description: "The table containing the details to check.", Type: shared.DocVarTypeTable},
			},
			Example:  "When I should see \"User\" details on the page\n| Name | John |\n| Age | 30 |",
			Category: shared.Visual,
		},
	)
}
