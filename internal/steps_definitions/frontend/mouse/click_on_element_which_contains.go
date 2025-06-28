package mouse

import (
	"context"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) clickOnElementWhichContains() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user clicks on {string} which contains "{string}"$`},
		func(scenarioCtx *scenario.Context) func(context.Context, string, string) (context.Context, error) {
			return func(ctx context.Context, _, text string) (context.Context, error) {
				element, err := scenarioCtx.GetCurrentPageOnly().GetOneByTextContent(text)
				if err != nil {
					return ctx, err
				}
				err = element.Click()
				return ctx, err
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "clicks on an element which contains a specific text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to click on.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user clicks on \"Submit button\" which contains \"Submit\"",
			Category: stepbuilder.Mouse,
		},
	)
}
