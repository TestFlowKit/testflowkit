package form

import (
	"context"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/label"
)

func (steps) selectOptionWithTextIntoDropdown() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`the user selects the option with text {string} from the {string} dropdown`},
		func(ctx context.Context, optionText, dropdownName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			dropdown, err := scenarioCtx.GetHTMLElementByLabel(label.Dropdown(dropdownName))
			if err != nil {
				return ctx, err
			}

			err = dropdown.SelectByText([]string{optionText})
			return ctx, err
		},
		func(_, dropdownName string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			fLabel := label.Dropdown(dropdownName)
			if !config.IsElementDefined(fLabel) {
				vc.AddMissingElement(fLabel)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "selects an option from a dropdown by its text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "text", Description: "The text of the option to select.", Type: stepbuilder.VarTypeString},
				{Name: "name", Description: "The logical name of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:    "When the user selects the option with text \"United States\" from the \"Country\" dropdown",
			Categories: []stepbuilder.StepCategory{stepbuilder.Form},
		},
	)
}
