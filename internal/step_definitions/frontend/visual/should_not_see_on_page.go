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
		[]string{`the user should not see {string} on the page`},
		func(ctx context.Context, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}
			elt, err := currentPage.GetOneByTextContent(text)
			if err != nil {
				// TODO: create specific error for not found element
				return ctx, nil
			}

			if elt.IsVisible() {
				return ctx, fmt.Errorf("the text \"%s\" is visible on the page but should not", strings.TrimSpace(text))
			}
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a specific text is not visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "text", Description: "The text that should not be visible on the page.", Type: stepbuilder.VarTypeString},
			},
			Example:    "Then the user should not see \"Error\" on the page",
			Categories: []stepbuilder.StepCategory{stepbuilder.Visual},
		},
	)
}
