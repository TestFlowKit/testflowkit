package visual

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) shouldSeeOnPage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user should see "{string}" on the page`},
		func(ctx context.Context, word string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}

			elt, err := currentPage.GetOneByTextContent(word)
			if err != nil {
				return ctx, err
			}

			if !elt.IsVisible() {
				return ctx, fmt.Errorf("%s should be visible", word)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a word is visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "word", Description: "The word to check.", Type: stepbuilder.VarTypeString},
			},
			Example:    "Then the user should see \"Submit\" on the page",
			Categories: []stepbuilder.StepCategory{stepbuilder.Visual},
		},
	)
}
