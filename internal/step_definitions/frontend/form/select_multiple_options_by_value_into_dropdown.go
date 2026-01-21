package form

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/label"
	"testflowkit/internal/utils/stringutils"
)

func (steps) userSelectMultipleOptionsByValueIntoDropdown() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the user selects the options with values {string} from the {string} dropdown`},
		func(ctx context.Context, optionValues, dropdownId string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			input, err := scenarioCtx.GetHTMLElementByLabel(label.Dropdown(dropdownId))
			if err != nil {
				return ctx, err
			}
			return ctx, input.SelectByValue(stringutils.SplitAndTrim(optionValues, ","))
		},
		func(_, dropdownId string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			fLabel := label.Dropdown(dropdownId)
			if !config.IsElementDefined(fLabel) {
				vc.AddMissingElement(fLabel)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "Selects multiple options by values from a dropdown.",
			Variables: []stepbuilder.DocVariable{
				{Name: "options values", Description: "The options values to select.", Type: stepbuilder.VarTypeString},
				{Name: "dropdownId", Description: "The id of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:    `When the user selects the options with values "UK,US" from the "Country" dropdown.`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Form},
		},
	)
}
