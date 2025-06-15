package assertions

import (
	"fmt"
	"reflect"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) radioButtonShouldBeSelectedOrNot() core.TestStep {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "radio_button")
	}

	definition := func(ctx *core.TestSuiteContext) func(string, string) error {
		return func(radioId, status string) error {
			input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatLabel(radioId))
			if err != nil {
				return err
			}
			checkValue, isBoolean := input.GetPropertyValue("checked", reflect.Bool).(bool)

			isSelected := isBoolean && checkValue
			isUnselected := isBoolean && !checkValue

			if isSelected && status == "selected" || isUnselected && status == "unselected" {
				return nil
			}

			return fmt.Errorf("%s radio button is not %s", radioId, status)
		}
	}

	validator := func(radioBtnName, _ string) core.ValidationErrors {
		vc := core.ValidationErrors{}
		radioLabel := formatLabel(radioBtnName)

		if !testsconfig.IsElementDefined(radioLabel) {
			vc.AddMissingElement(radioLabel)
		}

		return vc
	}

	statusType := shared.DocVarTypeEnum("selected", "unselected")

	return core.NewStepWithTwoVariables(
		[]string{`the {string} radio button should be (selected|unselected)`},
		definition,
		validator,
		core.StepDefDocParams{
			Description: "checks if the radio button is selected or unselected.",
			Variables: []shared.StepVariable{
				{Name: "radio button", Description: "The name of the radio button.", Type: shared.DocVarTypeString},
				{Name: "status", Description: "The status of the radio button.", Type: statusType},
			},
			Example:  `Then the "terms" radio button should be selected`,
			Category: shared.Form,
		},
	)
}
