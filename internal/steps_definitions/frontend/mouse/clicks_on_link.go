package mouse

import (
	"context"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) userClicksOnLink() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "link")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user clicks the {string} link$`},
		func(scenarioCtx *scenario.Context) func(context.Context, string) (context.Context, error) {
			return func(ctx context.Context, name string) (context.Context, error) {
				return clickCommonHandler(formatLabel).handler()(scenarioCtx)(ctx, name)
			}
		},
		clickCommonHandler(formatLabel).validation(),
		stepbuilder.DocParams{
			Description: "performs a click action on the link identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of link to click on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user clicks the \"Forgot Password\" link",
			Category: stepbuilder.Mouse,
		},
	)
}
