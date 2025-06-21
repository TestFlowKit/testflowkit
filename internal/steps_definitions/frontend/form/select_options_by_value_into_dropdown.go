package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) userSelectMultipleOptionsByValueIntoDropdown() stepbuilder.TestStep {
	return selectOptionsByValueIntoDropdownBuilder(
		[]string{`^the user selects the options with values {string} from the {string} dropdown$`},
		stepbuilder.StepDefDocParams{
			Description: "Selects multiple options by values from a dropdown.",
			Variables: []shared.StepVariable{
				{Name: "options values", Description: "The options values to select.", Type: shared.DocVarTypeString},
				{Name: "dropdownId", Description: "The id of the dropdown.", Type: shared.DocVarTypeString},
			},
			Example:  `When the user selects the options with values "UK,US" from the "Country" dropdown.`,
			Category: shared.Form,
		},
	)
}

func (s steps) userSelectOptionWithValueIntoDropdown() stepbuilder.TestStep {
	return selectOptionsByValueIntoDropdownBuilder(
		[]string{`^the user selects the option with value {string} from the {string} dropdown$`},
		stepbuilder.StepDefDocParams{
			Description: "Selects an option from a dropdown list based on its underlying ‘value’ attribute.",
			Variables: []shared.StepVariable{
				{Name: "option value", Description: "The value of the option to select.", Type: shared.DocVarTypeString},
				{Name: "name", Description: "The logical name of the dropdown.", Type: shared.DocVarTypeString},
			},
			Example:  `When the user selects the option with value "CIV" from the "Country" dropdown.`,
			Category: shared.Form,
		})
}

func selectOptionsByValueIntoDropdownBuilder(phrases []string, doc stepbuilder.StepDefDocParams) stepbuilder.TestStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}

	return stepbuilder.NewStepWithTwoVariables(
		phrases,
		func(ctx *scenario.Context) func(string, string) error {
			return func(optionValues, dropdownId string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatVar(dropdownId))
				if err != nil {
					return err
				}
				return input.SelectByValue(stringutils.SplitAndTrim(optionValues, ","))
			}
		},
		func(_, dropdownId string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatVar(dropdownId)
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		doc,
	)
}
