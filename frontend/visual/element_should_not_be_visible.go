package visual

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
	"fmt"
)

func (s steps) elementShouldNotBeVisible() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^{string} should not be visible$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(name string) error {
				element, err := browser.GetElement(ctx.GetCurrentPage(), name)
				if err != nil {
					return err
				}

				if element.IsVisible() {
					return fmt.Errorf("%s is visible", name)
				}

				return nil
			}
		},
	}
}
