package visual

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) shouldSeeDetailsOnPage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`the user should see "{string}" details on the page`},
		func(ctx context.Context, details string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			elt, err := scenarioCtx.GetCurrentPageOnly().GetOneBySelector("body")
			if err != nil {
				return ctx, err
			}
			if !strings.Contains(elt.TextContent(), details) {
				return ctx, fmt.Errorf("%s details should be visible", details)
			}
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "checks if details are visible on the page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "details", Description: "The details to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then the user should see \"User Profile\" details on the page",
			Category: stepbuilder.Visual,
		},
	)
}
