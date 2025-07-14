package visual

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (steps) shouldSeeDetailsOnPage() stepbuilder.Step {
	definition := func(ctx context.Context, elementName string, table *godog.Table) (context.Context, error) {
		scenarioCtx := scenario.MustFromContext(ctx)
		data, parseErr := assistdog.NewDefault().ParseMap(table)
		if parseErr != nil {
			return ctx, errors.New("details malformed please go to the doc")
		}

		parsedData, parseErr := scenario.ReplaceVariablesInMap(scenarioCtx, data)
		if parseErr != nil {
			return ctx, fmt.Errorf("failed to parse variables in table data: %w", parseErr)
		}

		currentPage := scenarioCtx.GetCurrentPageOnly()
		var errMsgs []string
		for name, value := range parsedData {
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
			return ctx, errors.New(strings.Join(errMsgs, "\n"))
		}

		return ctx, nil
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`the user should see "{string}" details on the page`},
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
