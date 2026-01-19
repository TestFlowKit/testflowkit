package mouse

import (
	"context"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (steps) userClicksOnButton() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "button")
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`the user clicks the {string} button`},
		func(ctx context.Context, name string) (context.Context, error) {
			return clickCommonHandler(formatLabel).handler()(ctx, name)
		},
		clickCommonHandler(formatLabel).validation(),
		stepbuilder.DocParams{
			Description: "performs a click action on the button identified by its logical name",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of button to click on.", Type: stepbuilder.VarTypeString},
			},
			Example:    "When the user clicks the \"Submit Order\" button",
			Categories: []stepbuilder.StepCategory{stepbuilder.Mouse},
		},
	)
}
