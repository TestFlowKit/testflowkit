package form

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/label"
)

func (steps) userSelectOptionWithValueIntoDropdown() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the user selects the option with value {string} from the {string} dropdown`},
		func(ctx context.Context, optionValue, dropdownId string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			input, err := scenarioCtx.GetHTMLElementByLabel(label.Dropdown(dropdownId))
			if err != nil {
				return ctx, err
			}
			return ctx, input.SelectByValue([]string{optionValue})
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
			Description: "Selects an option from a dropdown list based on its underlying 'value' attribute.",
			Variables: []stepbuilder.DocVariable{
				{Name: "option value", Description: "The value of the option to select.", Type: stepbuilder.VarTypeString},
				{Name: "name", Description: "The logical name of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:    `When the user selects the option with value "CIV" from the "Country" dropdown.`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Form},
		},
	)
}
