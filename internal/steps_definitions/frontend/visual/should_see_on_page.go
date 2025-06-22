package visual

import (
	"fmt"
	"strings"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) shouldSeeOnPage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user should see "{string}" on the page$`},
		func(ctx *scenario.Context) func(string) error {
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
		stepbuilder.DocParams{
			Description: "checks if a word is visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "word", Description: "The word to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the user should see \"Submit\" on the page",
			Category: stepbuilder.Visual,
		},
	)
}
