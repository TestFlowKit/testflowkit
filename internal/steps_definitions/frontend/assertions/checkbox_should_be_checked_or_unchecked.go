package assertions

import (
	"fmt"
	"reflect"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	sb "testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) checkCheckboxStatus() sb.Step {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_checkbox", label)
	}
	definition := func(ctx *scenario.Context) func(string, string) error {
		return func(checkboxId, status string) error {
			input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatVar(checkboxId))
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

	validator := func(checkboxId, _ string) sb.ValidationErrors {
		vc := sb.ValidationErrors{}
		checkboxLabel := formatVar(checkboxId)

		if !testsconfig.IsElementDefined(checkboxLabel) {
			vc.AddMissingElement(checkboxLabel)
		}

		return vc
	}

	return sb.NewWithTwoVariables(
		[]string{`the {string} checkbox should be (checked|unchecked)`},
		definition,
		validator,
		sb.DocParams{
			Description: "checks if the checkbox is checked or unchecked.",
			Variables: []sb.DocVariable{
				{Name: "checkboxId", Description: "The id of the checkbox.", Type: sb.VarTypeString},
				{Name: "status", Description: "The status of the checkbox.", Type: sb.DocVarTypeEnum("checked", "unchecked")},
			},
			Example:  `Then the "terms" checkbox should be checked`,
			Category: sb.Form,
		},
	)
}
