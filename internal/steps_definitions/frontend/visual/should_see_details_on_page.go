package visual

import (
	"errors"
	"fmt"
	"strings"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (s steps) shouldSeeDetailsOnPage() stepbuilder.TestStep {
	definition := func(ctx *stepbuilder.TestSuiteContext) func(string, *godog.Table) error {
		return func(elementName string, table *godog.Table) error {
			data, parseErr := assistdog.NewDefault().ParseMap(table)
			if parseErr != nil {
				return errors.New("details malformed please go to the doc")
			}

			var errMsgs []string
			for name, value := range data {
				elt, err := ctx.GetCurrentPage().GetOneByTextContent(value)
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

	return stepbuilder.NewStepWithTwoVariables(
		[]string{`^the user should see "{string}" details on the page$`},
		definition,
		nil,
		stepbuilder.StepDefDocParams{
			Description: "checks if the details are visible on the page.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: shared.DocVarTypeString},
				{Name: "table", Description: "The table containing the details to check.", Type: shared.DocVarTypeTable},
			},
			Example:  "When the user should see \"User\" details on the page\n| Name | John |\n| Age | 30 |",
			Category: shared.Visual,
		},
	)
}
