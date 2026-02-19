package visual

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) shouldSeeOnPage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user should see {string} on the page`},
		func(ctx context.Context, text string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			currentPage, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}
			elt, errGetOne := currentPage.GetOneByTextContent(text)
			if errGetOne != nil {
				return ctx, errGetOne
			}

			if !elt.IsVisible() {
				return ctx, fmt.Errorf("%s is not visible", text)
			}

			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if a text is visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "text", Description: "The text to check.", Type: stepbuilder.VarTypeString},
			},
			Example:    "Then the user should see \"Submit\" on the page",
			Categories: []stepbuilder.StepCategory{stepbuilder.Visual},
		},
	)
}
