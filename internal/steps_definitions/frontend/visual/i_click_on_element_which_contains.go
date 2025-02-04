package visual

import (
	"etoolse/internal/steps_definitions/core"
	"fmt"
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
	)
}
