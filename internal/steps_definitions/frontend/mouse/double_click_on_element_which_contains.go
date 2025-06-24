package mouse

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) doubleClickOnElementWhichContains() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user double clicks on {string} which contains "{string}"$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(_, text string) error {
				element, err := ctx.GetCurrentPageOnly().GetOneByTextContent(text)
				if err != nil {
					return err
				}

				return element.DoubleClick()
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Double clicks on an element which contains a specific text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to double click on.", Type: stepbuilder.VarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user double clicks on \"Submit button\" which contains \"Submit\"",
			Category: stepbuilder.Mouse,
		},
	)
}
