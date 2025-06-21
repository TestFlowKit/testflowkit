package mouse

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) doubleClickOnElementWhichContains() stepbuilder.TestStep {
	return stepbuilder.NewStepWithTwoVariables(
		[]string{`^the user double clicks on {string} which contains "{string}"$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(_ string, text string) error {
				element, err := ctx.GetCurrentPage().GetOneByTextContent(text)
				if err != nil {
					return fmt.Errorf("no element with text containing %s found", text)
				}
				return element.DoubleClick()
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "Double clicks on an element which contains a specific text.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the element to double click on.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: shared.DocVarTypeString},
			},
			Example:  "When the user double clicks on \"Submit button\" which contains \"Submit\"",
			Category: shared.Mouse,
		},
	)
}
