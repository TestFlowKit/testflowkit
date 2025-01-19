package form

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"fmt"
	"reflect"
)

func (s steps) checkCheckboxStatus() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^the {string} checkbox should be (checked|unchecked)`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(checkboxId, status string) error {
				input, err := browser.GetElement(ctx.GetCurrentPage(), fmt.Sprintf("%s_checkbox", checkboxId))
				if err != nil {
					return err
				}
				checkValue, isBoolean := input.GetPropertyValue("checked", reflect.Bool).(bool)

				if isBoolean && checkValue && status == "checked" || !checkValue && status == "unchecked" {
					return nil
				}

				return fmt.Errorf("%s checkbox is not %s", checkboxId, status)
			}
		},
	}
}
