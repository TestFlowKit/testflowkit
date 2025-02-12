package form

import (
	"etoolse/internal/browser"
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
	"fmt"
)

func (s steps) theFieldShouldContains() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^the {string} should be contain "{string}"`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(fieldId, text string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), fieldId)
				if err != nil {
					return err
				}

				if input.TextContent() == text {
					return nil
				}

				return fmt.Errorf("field should be contains %s but contains %s", text, input.TextContent())
			}
		},
		func(fieldId, _ string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsElementDefined(fieldId) {
				vc.AddMissingElement(fieldId)
			}

			return vc
		},
	)
}
