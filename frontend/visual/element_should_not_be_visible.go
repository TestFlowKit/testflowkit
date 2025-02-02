package visual

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"etoolse/internal/config/testsConfig"
	"fmt"
)

func (s steps) elementShouldNotBeVisible() common.TestStep {
	return common.NewStepWithOneVariable(
		[]string{`^{string} should not be visible$`},
		func(ctx *common.TestSuiteContext) func(string) error {
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
		func(name string) common.ValidationErrors {
			vc := common.ValidationErrors{}
			if !testsConfig.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
	)
}
