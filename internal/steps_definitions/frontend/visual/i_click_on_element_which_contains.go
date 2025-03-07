package visual

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) iClickOnElementWhichContains() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I click on {string} which contains "{string}"$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(_ string, text string) error {
				xPath := fmt.Sprintf(`//*[contains(text(),"%s")]`, text)
				element, err := ctx.GetCurrentPage().GetOneByXPath(xPath)
				if err != nil {
					return fmt.Errorf("no element with text containing %s found", text)
				}
				return element.Click()
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "clicks on an element which contains a specific text.",
			Variables: []shared.StepVariable{
				{Name: "label", Description: "The label of the element to click on.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: shared.DocVarTypeString},
			},
			Example:  "When I click on \"Submit button\" which contains \"Submit\"",
			Category: shared.Visual,
		},
	)
}
