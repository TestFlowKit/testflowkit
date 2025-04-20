package mouse

import (
	"fmt"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) iRightClickOnElementWhichContains() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I right click on {string} which contains "{string}"$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(_ string, text string) error {
				xPath := fmt.Sprintf(`//*[contains(text(),"%s")]`, text)
				element, err := ctx.GetCurrentPage().GetOneByXPath(xPath)
				if err != nil {
					return fmt.Errorf("no element with text containing %s found", text)
				}
				return element.RightClick()
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "Right clicks on an element which contains a specific text.",
			Variables: []shared.StepVariable{
				{Name: "label", Description: "The label of the element to right click on.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: shared.DocVarTypeString},
			},
			Example:  "When I right click on \"Submit button\" which contains \"Submit\"",
			Category: shared.Mouse,
		},
	)
}
