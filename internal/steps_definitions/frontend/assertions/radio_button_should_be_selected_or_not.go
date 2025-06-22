package assertions

import (
	"fmt"
	"reflect"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) radioButtonShouldBeSelectedOrNot() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "radio_button")
	}

	definition := func(ctx *scenario.Context) func(string, string) error {
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

	validator := func(radioBtnName, _ string) stepbuilder.ValidationErrors {
		vc := stepbuilder.ValidationErrors{}
		radioLabel := formatLabel(radioBtnName)

		if !testsconfig.IsElementDefined(radioLabel) {
			vc.AddMissingElement(radioLabel)
		}

		return vc
	}

	statusType := stepbuilder.DocVarTypeEnum("selected", "unselected")

	return stepbuilder.NewWithTwoVariables(
		[]string{`the {string} radio button should be (selected|unselected)`},
		definition,
		validator,
		stepbuilder.DocParams{
			Description: "checks if the radio button is selected or unselected.",
			Variables: []stepbuilder.DocVariable{
				{Name: "radio button", Description: "The name of the radio button.", Type: stepbuilder.VarTypeString},
				{Name: "status", Description: "The status of the radio button.", Type: statusType},
			},
			Example:  `Then the "terms" radio button should be selected`,
			Category: stepbuilder.Form,
		},
	)
}
