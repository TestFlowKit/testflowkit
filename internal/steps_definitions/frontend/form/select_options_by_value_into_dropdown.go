package form

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (steps) userSelectMultipleOptionsByValueIntoDropdown() stepbuilder.Step {
	return selectOptionsByValueIntoDropdownBuilder(
		[]string{`^the user selects the options with values {string} from the {string} dropdown$`},
		stepbuilder.DocParams{
			Description: "Selects multiple options by values from a dropdown.",
			Variables: []stepbuilder.DocVariable{
				{Name: "options values", Description: "The options values to select.", Type: stepbuilder.VarTypeString},
				{Name: "dropdownId", Description: "The id of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user selects the options with values "UK,US" from the "Country" dropdown.`,
			Category: stepbuilder.Form,
		},
	)
}

func (steps) userSelectOptionWithValueIntoDropdown() stepbuilder.Step {
	return selectOptionsByValueIntoDropdownBuilder(
		[]string{`^the user selects the option with value {string} from the {string} dropdown$`},
		stepbuilder.DocParams{
			Description: "Selects an option from a dropdown list based on its underlying 'value' attribute.",
			Variables: []stepbuilder.DocVariable{
				{Name: "option value", Description: "The value of the option to select.", Type: stepbuilder.VarTypeString},
				{Name: "name", Description: "The logical name of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user selects the option with value "CIV" from the "Country" dropdown.`,
			Category: stepbuilder.Form,
		})
}

func selectOptionsByValueIntoDropdownBuilder(phrases []string, doc stepbuilder.DocParams) stepbuilder.Step {
	formatVar := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "dropdown")
	}

	return stepbuilder.NewWithTwoVariables(
		phrases,
		func(ctx context.Context, optionValues, dropdownId string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			input, err := scenarioCtx.GetHTMLElementByLabel(formatVar(dropdownId))
			if err != nil {
				return ctx, err
			}
			return ctx, input.SelectByValue(stringutils.SplitAndTrim(optionValues, ","))
		},
		func(_, dropdownId string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatVar(dropdownId)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		doc,
	)
}
