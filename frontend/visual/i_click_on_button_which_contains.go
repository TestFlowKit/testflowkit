package visual

import (
	"cucumber/frontend/common"
	"fmt"
)

var iClickOnElementWhichContains = common.FrontStep{
	Sentences: []string{`^I click on element which contains "{string}"$`},
	Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
		return func(text string) error {
			xPath := fmt.Sprintf("//*[contains(text(),%s)]", text)
			element, err := ctx.GetCurrentPage().GetOneByXPath(xPath)
			if err != nil {
				return fmt.Errorf("no element with text containing %s found", text)
			}
			return element.Click()
		}
	},
}
