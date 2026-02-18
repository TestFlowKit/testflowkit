package visual

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (steps) shouldSeeDetailsOnPage() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "details")
	}
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

		elementContainer, errElContainer := scenarioCtx.GetHTMLElementByLabel(formatLabel(elementName))
		if errElContainer != nil {
			return ctx, errElContainer
		}

		eltTextContent := stringutils.Inline(elementContainer.TextContent())
		var errMsgs []string
		for name, value := range parsedData {
			if !strings.Contains(eltTextContent, value) {
				errMsgs = append(errMsgs, fmt.Sprintf("%s %s not found", elementName, name))
				continue
			}
		}

		if len(errMsgs) > 0 {
			return ctx, errors.New(strings.Join(errMsgs, "\n"))
		}

		return ctx, nil
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`the user should see {string} details on the page`},
		definition,
		nil,
		stepbuilder.DocParams{
			Description: "checks if the details are visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "table", Description: "The table containing the details to check.", Type: stepbuilder.VarTypeTable},
			},
			Example:    "When the user should see \"User\" details on the page\n| Name | John |\n| Age | 30 |",
			Categories: []stepbuilder.StepCategory{stepbuilder.Visual},
		},
	)
}
