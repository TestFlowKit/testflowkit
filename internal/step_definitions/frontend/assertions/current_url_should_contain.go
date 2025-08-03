package assertions

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) currentURLShouldContain() stepbuilder.Step {
	const description = "The URL substring that should be contained in the current URL."
	return stepbuilder.NewWithOneVariable(
		[]string{`the current URL should contain {string}`},
		func(ctx context.Context, expectedURLPart string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page := scenarioCtx.GetCurrentPageOnly()

			pageInfo := page.GetInfo()
			if !strings.Contains(pageInfo.URL, expectedURLPart) {
				return ctx, fmt.Errorf("current URL '%s' does not contain '%s'", pageInfo.URL, expectedURLPart)
			}

			return ctx, nil
		},
		func(_ string) stepbuilder.ValidationErrors {
			return stepbuilder.ValidationErrors{}
		},
		stepbuilder.DocParams{
			Description: "This assertion checks if the current page URL contains the specified substring.",
			Variables: []stepbuilder.DocVariable{
				{Name: "expectedURLPart", Description: description, Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the current URL should contain "dashboard"`,
			Category: stepbuilder.Assertions,
		},
	)
}
