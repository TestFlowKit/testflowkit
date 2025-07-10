package visual

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) shouldNotSeeOnPage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user should not see "{string}" on the page`},
		func(ctx context.Context, word string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			elt, err := scenarioCtx.GetCurrentPageOnly().GetOneBySelector("body")
			if err != nil {
				return ctx, err
			}
			if strings.Contains(elt.TextContent(), word) {
				return ctx, fmt.Errorf("%s should not be visible", word)
			}
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a word is not visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "word", Description: "The word to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the user should not see \"Submit\" on the page",
			Category: stepbuilder.Visual,
		},
	)
}
