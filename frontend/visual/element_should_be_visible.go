package visual

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"fmt"
)

func (s steps) elementShouldBeVisible() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^{string} should be visible$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(name string) error {
				element, err := browser.GetElement(ctx.GetCurrentPage(), name)
				if err != nil {
					return err
				}

				if !element.IsVisible() {
					return fmt.Errorf("%s is not visible", name)
				}

				return nil
			}
		},
	}
}
