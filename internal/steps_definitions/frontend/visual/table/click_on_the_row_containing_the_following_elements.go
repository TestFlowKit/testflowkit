package table

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testflowkit/internal/browser/common"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

// TODO: click on cell instead of row
func (steps) clickOnTheRowContainingTheFollowingElements() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user clicks on the row containing the following elements`},
		func(ctx context.Context, table *godog.Table) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			data, parseErr := assistdog.NewDefault().ParseMap(table)
			if parseErr != nil {
				return ctx, errors.New("table malformed please go to the doc")
			}

			currentPage := scenarioCtx.GetCurrentPageOnly()
			errMsgs := checkIfElementsAreVisible(data, currentPage)
			if len(errMsgs) > 0 {
				return ctx, errors.New(strings.Join(errMsgs, "\n"))
			}

			// Click on the first element found
			for _, value := range data {
				elt, err := currentPage.GetOneByTextContent(value)
				if err == nil && elt.IsVisible() {
					err = elt.Click()
					if err != nil {
						return ctx, err
					}
					break
				}
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "clicks on a table row that contains the specified elements.",
			Variables: []stepbuilder.DocVariable{
				{Name: "table", Description: "The table containing the elements to click on.", Type: stepbuilder.VarTypeTable},
			},
			Example:  "When the user clicks on the row containing the following elements\n| Name | John |\n| Age | 30 |",
			Category: stepbuilder.Visual,
		},
	)
}

func checkIfElementsAreVisible(data map[string]string, currentPage common.Page) []string {
	var errMsgs []string
	for name, value := range data {
		elt, err := currentPage.GetOneByTextContent(value)
		if err != nil {
			errMsgs = append(errMsgs, fmt.Sprintf("%s %s not found", name, value))
			continue
		}

		if !elt.IsVisible() {
			errMsgs = append(errMsgs, fmt.Sprintf("%s %s is found but is not visible", name, value))
		}
	}
	return errMsgs
}
