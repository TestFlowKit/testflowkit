package visual

import (
	"etoolse/frontend/common"
	"fmt"
)

func (s steps) iClickOnElementWhichContains() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I click on {string} which contains "{string}"$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(_, text string) error {
				xPath := fmt.Sprintf(`//*[contains(text(),"%s")]`, text)
				element, err := ctx.GetCurrentPage().GetOneByXPath(xPath)
				if err != nil {
					return fmt.Errorf("no element with text containing %s found", text)
				}
				return element.Click()
			}
		},
	}
}
