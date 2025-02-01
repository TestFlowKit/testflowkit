package visual

import (
	"etoolse/config"
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"fmt"
)

func (s steps) elementShouldBeVisible() common.FrontStep {
	return common.NewStepWithOneVariable(
		[]string{`^{string} should be visible$`},
		func(ctx *common.TestSuiteContext) func(string) error {
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
		func(name string) common.ValidationErrors {
			vc := common.ValidationErrors{}
			if !config.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
	)
}
