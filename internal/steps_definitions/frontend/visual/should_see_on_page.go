package visual

import (
	"fmt"
	"strings"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) shouldSeeOnPage() stepbuilder.TestStep {
	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user should see "{string}" on the page$`},
		func(ctx *stepbuilder.TestSuiteContext) func(string) error {
			return func(word string) error {
				elt, err := ctx.GetCurrentPage().GetOneBySelector("body")
				if err != nil {
					return err
				}
				if !strings.Contains(elt.TextContent(), word) {
					return fmt.Errorf("%s should be visible", word)
				}
				return nil
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "checks if a word is visible on the page.",
			Variables: []shared.StepVariable{
				{Name: "word", Description: "The word to check.", Type: shared.DocVarTypeString},
			},
			Example:  "Then the user should see \"Submit\" on the page",
			Category: shared.Visual,
		},
	)
}
