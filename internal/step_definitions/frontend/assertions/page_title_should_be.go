package assertions

import (
	"context"
	"fmt"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) pageTitleShouldBe() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the page title should be {string}`},
		func(ctx context.Context, expectedTitle string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}
			pageInfo := page.GetInfo()
			if pageInfo.Title != expectedTitle {
				return ctx, fmt.Errorf("expected page title to be '%s', but found '%s'", expectedTitle, pageInfo.Title)
			}
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "This assertion checks if the current page title matches the specified title exactly.",
			Variables: []stepbuilder.DocVariable{
				{Name: "expectedTitle", Description: "The expected page title to match.", Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the page title should be "Welcome to TestFlowKit"`,
			Category: stepbuilder.Assertions,
		},
	)
}
