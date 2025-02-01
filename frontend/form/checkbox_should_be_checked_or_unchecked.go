package form

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"fmt"
	"reflect"
)

func (s steps) checkCheckboxStatus() common.FrontStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_checkbox", label)
	}
	definition := func(ctx *common.TestSuiteContext) func(string, string) error {
		return func(checkboxId, status string) error {
			input, err := browser.GetElement(ctx.GetCurrentPage(), formatVar(checkboxId))
			if err != nil {
				return err
			}
			checkValue, isBoolean := input.GetPropertyValue("checked", reflect.Bool).(bool)

			if isBoolean && checkValue && status == "checked" || !checkValue && status == "unchecked" {
				return nil
			}

			return fmt.Errorf("%s checkbox is not %s", checkboxId, status)
		}
	}

	return common.NewStepWithTwoVariables(
		[]string{`^the {string} checkbox should be (checked|unchecked)`},
		definition,
		nil,
	)
}
