package form

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
	"fmt"
)

func (s steps) theFieldShouldContains() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^the {string} should be contain "{string}"`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(fieldId, text string) error {
				input, err := browser.GetInputElement(ctx.GetCurrentPage(), fieldId)
				if err != nil {
					return err
				}

				if input.TextContent() == text {
					return nil
				}

				return fmt.Errorf("field should be contains %s bue contains %s", text, input.TextContent())
			}
		},
	}
}
