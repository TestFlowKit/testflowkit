package form

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"fmt"
)

func (s steps) theFieldShouldContains() common.FrontStep {
	return common.NewStepWithTwoVariables(
		[]string{`^the {string} should be contain "{string}"`},
		func(ctx *common.TestSuiteContext) func(string, string) error {
			return func(fieldId, text string) error {
				input, err := browser.GetElement(ctx.GetCurrentPage(), fieldId)
				if err != nil {
					return err
				}

				if input.TextContent() == text {
					return nil
				}

				return fmt.Errorf("field should be contains %s but contains %s", text, input.TextContent())
			}
		},
		nil,
	)
}
