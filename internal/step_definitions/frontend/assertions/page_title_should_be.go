package assertions

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (s steps) pageTitleShouldBe() stepbuilder.Step {
	return s.newPageTitleStep(
		`the page title should be {string}`,
		true,
		stepbuilder.DocParams{
			Description: "This assertion checks if the current page title matches the specified title exactly.",
			Variables: []stepbuilder.DocVariable{
				{Name: "expectedTitle", Description: "The expected page title to match.", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the page title should be "Welcome to TestFlowKit"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Assertions},
		},
	)
}

func (s steps) pageTitleShouldNotBe() stepbuilder.Step {
	return s.newPageTitleStep(
		`the page title should not be {string}`,
		false,
		stepbuilder.DocParams{
			Description: "This assertion checks if the current page title is different from the specified title.",
			Variables: []stepbuilder.DocVariable{
				{Name: "unexpectedTitle", Description: "The title that should not match the current page title.",
					Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the page title should not be "Error"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Assertions},
		},
	)
}

func (s steps) newPageTitleStep(sentence string, shouldEqual bool, doc stepbuilder.DocParams) stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{sentence},
		func(ctx context.Context, title string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}
			pageInfo := page.GetInfo()
			if shouldEqual && pageInfo.Title != title {
				return ctx, fmt.Errorf("expected page title to be '%s', but found '%s'", title, pageInfo.Title)
			}

			if !shouldEqual && pageInfo.Title == title {
				return ctx, fmt.Errorf("page title should not be '%s'", title)
			}
			return ctx, nil
		},
		nil,
		doc,
	)
}
