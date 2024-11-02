package visual

import (
	"cucumber/frontend/common"
	"fmt"
)

var iMustSeeOnPageAnElementWithText = common.FrontStep{
	Sentences: []string{`^I must see on page a (link|button|element) with text "{string}"$`},
	Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
		return func(elementLabel, text string) error {
			cases := map[string]string{
				"link":    "a",
				"button":  "button",
				"element": "*",
			}

			xPath := fmt.Sprintf("//%s[contains(text(),\"%s\")]", cases[elementLabel], text)
			element, err := ctx.GetCurrentPage().GetOneByXPath(xPath)
			cErr := fmt.Errorf("no %s is visible with text \"%s\"", elementLabel, text)
			if err != nil {
				return cErr
			}

			visible := element.IsVisible()
			if !visible {
				return cErr
			}

			return nil
		}
	},
}
