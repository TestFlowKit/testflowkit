package mouse

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) hoverOnElementWhichContains() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user hovers on {string} which contains {string}$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(_, text string) error {
				element, err := ctx.GetCurrentPageOnly().GetOneByTextContent(text)
				if err != nil {
					return err
				}

				return element.Hover()
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Hover on an element which contains a specific text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to hover on.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user hovers on \"Submit button\" which contains \"Submit\"",
			Category: stepbuilder.Mouse,
		},
	)
}
