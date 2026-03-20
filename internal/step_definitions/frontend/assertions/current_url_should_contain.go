package assertions

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (s steps) currentURLShouldContain() stepbuilder.Step {
	return s.newCurrentURLContainsStep(
		`the current URL should contain {string}`,
		true,
		stepbuilder.DocParams{
			Description: "This assertion checks if the current page URL contains the specified substring.",
			Variables: []stepbuilder.DocVariable{{
				Name:        "expectedURLPart",
				Description: "The URL substring that should be contained in the current URL.",
				Type:        stepbuilder.VarTypeString,
			}},
			Example:    `Then the current URL should contain "dashboard"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Assertions},
		},
	)
}

func (s steps) currentURLShouldNotContain() stepbuilder.Step {
	return s.newCurrentURLContainsStep(
		`the current URL should not contain {string}`,
		false,
		stepbuilder.DocParams{
			Description: "This assertion checks if the current page URL does not contain the specified substring.",
			Variables: []stepbuilder.DocVariable{{
				Name:        "unexpectedURLPart",
				Description: "The URL substring that should not be contained in the current URL.",
				Type:        stepbuilder.VarTypeString,
			}},
			Example:    `Then the current URL should not contain "error"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Assertions},
		},
	)
}

func (s steps) newCurrentURLContainsStep(
	sentence string, shouldContain bool, doc stepbuilder.DocParams,
) stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{sentence},
		func(ctx context.Context, urlPart string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			page, pageErr := scenarioCtx.GetCurrentPageOnly()
			if pageErr != nil {
				return ctx, pageErr
			}

			pageInfo := page.GetInfo()
			contains := strings.Contains(pageInfo.URL, urlPart)
			if shouldContain && !contains {
				return ctx, fmt.Errorf("current URL '%s' does not contain '%s'", pageInfo.URL, urlPart)
			}

			if !shouldContain && contains {
				return ctx, fmt.Errorf("current URL '%s' should not contain '%s'", pageInfo.URL, urlPart)
			}
			return ctx, nil
		},
		nil,
		doc,
	)
}
