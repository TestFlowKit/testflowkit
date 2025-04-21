package mouse

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) iHoverOnElementWhichContains() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I hover on {string} which contains {string}$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(_ string, text string) error {
				xPath := fmt.Sprintf(`//*[contains(text(),"%s")]`, text)
				element, err := ctx.GetCurrentPage().GetOneByXPath(xPath)
				if err != nil {
					return fmt.Errorf("no element with text containing %s found", text)
				}
				return element.Hover()
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "Hover on an element which contains a specific text.",
			Variables: []shared.StepVariable{
				{Name: "label", Description: "The name of the element to hover on.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: shared.DocVarTypeString},
			},
			Example:  "When I hover on \"Submit button\" which contains \"Submit\"",
			Category: shared.Mouse,
		},
	)
}
