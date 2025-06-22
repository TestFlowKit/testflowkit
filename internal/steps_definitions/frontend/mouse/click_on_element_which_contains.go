package mouse

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) clickOnElementWhichContains() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user clicks on {string} which contains "{string}"$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(_ string, text string) error {
				element, err := ctx.GetCurrentPage().GetOneByTextContent(text)
				if err != nil {
					return fmt.Errorf("no element with text containing %s found", text)
				}
				return element.Click()
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
