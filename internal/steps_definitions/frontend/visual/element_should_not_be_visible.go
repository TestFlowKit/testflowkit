package visual

import (
	"etoolse/internal/browser"
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
	"fmt"
)

func (s steps) elementShouldNotBeVisible() core.TestStep {
	return core.NewStepWithOneVariable(
		[]string{`^{string} should not be visible$`},
		func(ctx *core.TestSuiteContext) func(string) error {
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
		func(name string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
	)
}
