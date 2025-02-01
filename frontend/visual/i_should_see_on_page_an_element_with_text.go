package visual

import (
	"etoolse/frontend/common"
	"fmt"
)

func (s steps) iShouldSeeElementWhichContains() common.FrontStep {
	return common.NewStepWithTwoVariables(
		[]string{`^I should see a (link|button|element) which contains "{string}"$`},
		func(ctx *common.TestSuiteContext) func(string, string) error {
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
		nil,
	)
}
