package mouse

import (
	"context"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) userClicksOnElement() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "element")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user clicks the {string} element$`},
		func(scenarioCtx *scenario.Context) func(context.Context, string) (context.Context, error) {
			return func(ctx context.Context, name string) (context.Context, error) {
				return clickCommonHandler(formatLabel).handler()(scenarioCtx)(ctx, name)
			}
		},
		clickCommonHandler(formatLabel).validation(),
		stepbuilder.DocParams{
			Description: "performs a click action on the element identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of element to click on.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user clicks the \"Main Logo\" element",
			Category: stepbuilder.Mouse,
		},
	)
}
