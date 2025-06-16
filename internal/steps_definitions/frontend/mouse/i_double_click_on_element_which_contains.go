package mouse

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) iDoubleClickOnElementWhichContains() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I double click on {string} which contains "{string}"$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(_ string, text string) error {
				element, err := ctx.GetCurrentPage().GetOneByTextContent(text)
				if err != nil {
					return fmt.Errorf("no element with text containing %s found", text)
				}
				return element.DoubleClick()
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "Double clicks on an element which contains a specific text.",
			Variables: []shared.StepVariable{
				{Name: "label", Description: "The label of the element to double click on.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: shared.DocVarTypeString},
			},
			Example:  "When I double click on \"Submit button\" which contains \"Submit\"",
			Category: shared.Mouse,
		},
	)
}
