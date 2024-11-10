package visual

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
	"fmt"
)

func (s steps) elementShouldBeVisible() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^{string} (input|button|element) should be visible$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(name, elementType string) error {
				element, err := browser.GetElementByType(ctx.GetCurrentPage(), name, elementType)
				if err != nil {
					return err
				}

				if !element.IsVisible() {
					return fmt.Errorf("%s %s is not visible", name, elementType)
				}

				return nil
			}
		},
	}
}
