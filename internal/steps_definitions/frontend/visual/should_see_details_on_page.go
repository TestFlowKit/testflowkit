package visual

import (
	"errors"
	"fmt"
	"strings"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (s steps) shouldSeeDetailsOnPage() stepbuilder.Step {
	definition := func(ctx *scenario.Context) func(string, *godog.Table) error {
		return func(elementName string, table *godog.Table) error {
			data, parseErr := assistdog.NewDefault().ParseMap(table)
			if parseErr != nil {
				return errors.New("details malformed please go to the doc")
			}

			currentPage := ctx.GetCurrentPageOnly()
			var errMsgs []string
			for name, value := range data {
				elt, err := currentPage.GetOneByTextContent(value)
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

	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user should see "{string}" details on the page$`},
		definition,
		nil,
		stepbuilder.DocParams{
			Description: "checks if the details are visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "table", Description: "The table containing the details to check.", Type: stepbuilder.VarTypeTable},
			},
			Example:  "When the user should see \"User\" details on the page\n| Name | John |\n| Age | 30 |",
			Category: stepbuilder.Visual,
		},
	)
}
