package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) userSelectMultipleOptionsByValueIntoDropdown() core.TestStep {
	return selectOptionsByValueIntoDropdownBuilder(
		[]string{`^the user selects the options with values {string} from the {string} dropdown$`},
		core.StepDefDocParams{
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

func (s steps) userSelectOptionWithValueIntoDropdown() core.TestStep {
	return selectOptionsByValueIntoDropdownBuilder(
		[]string{`^the user selects the option with value {string} from the {string} dropdown$`},
		core.StepDefDocParams{
			Description: "Selects an option from a dropdown list based on its underlying ‘value’ attribute.",
			Variables: []shared.StepVariable{
				{Name: "option value", Description: "The value of the option to select.", Type: shared.DocVarTypeString},
				{Name: "dropdownId", Description: "The id of the dropdown.", Type: shared.DocVarTypeString},
			},
			Example:  `When the user selects the option with value “CIV” from the “Country” dropdown.`,
			Category: shared.Form,
		})
}

func selectOptionsByValueIntoDropdownBuilder(phrases []string, doc core.StepDefDocParams) core.TestStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}

	return core.NewStepWithTwoVariables(
		phrases,
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(optionValues, dropdownId string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatVar(dropdownId))
				if err != nil {
					return err
				}
				return input.SelectByValue(stringutils.SplitAndTrim(optionValues, ","))
			}
		},
		func(_, dropdownId string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			label := formatVar(dropdownId)
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		doc,
	)
}
