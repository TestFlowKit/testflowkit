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
			page := scenarioCtx.GetCurrentPageOnly()

			pageInfo := page.GetInfo()
			if pageInfo.Title != expectedTitle {
				return ctx, fmt.Errorf("page title is '%s', expected '%s'", pageInfo.Title, expectedTitle)
			}

			return ctx, nil
		},
		func(_ string) stepbuilder.ValidationErrors {
			return stepbuilder.ValidationErrors{}
		},
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
