package assertions

import (
	"fmt"
	"reflect"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) checkCheckboxStatus() stepbuilder.TestStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_checkbox", label)
	}
	definition := func(ctx *stepbuilder.TestSuiteContext) func(string, string) error {
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

	validator := func(checkboxId, _ string) stepbuilder.ValidationErrors {
		vc := stepbuilder.ValidationErrors{}
		checkboxLabel := formatVar(checkboxId)

		if !testsconfig.IsElementDefined(checkboxLabel) {
			vc.AddMissingElement(checkboxLabel)
		}

		return vc
	}

	return stepbuilder.NewStepWithTwoVariables(
		[]string{`the {string} checkbox should be (checked|unchecked)`},
		definition,
		validator,
		stepbuilder.StepDefDocParams{
			Description: "checks if the checkbox is checked or unchecked.",
			Variables: []shared.StepVariable{
				{Name: "checkboxId", Description: "The id of the checkbox.", Type: shared.DocVarTypeString},
				{Name: "status", Description: "The status of the checkbox.", Type: shared.DocVarTypeEnum("checked", "unchecked")},
			},
			Example:  `Then the "terms" checkbox should be checked`,
			Category: shared.Form,
		},
	)
}
