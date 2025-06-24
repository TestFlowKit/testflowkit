package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) selectOptionWithTextIntoDropdown() stepbuilder.Step {
	formatVar := func(variable string) string {
		return variable + "_dropdown"
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user selects the option with text {string} from the {string} dropdown$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(optionText, dropdownId string) error {
				currentPage, pageName := ctx.GetCurrentPage()
				input, err := browser.GetElementByLabel(currentPage, pageName, formatVar(dropdownId))
				if err != nil {
					return err
				}

				return input.SelectByText([]string{optionText})
			}
		},
		func(_, dropdownId string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatVar(dropdownId)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "selects an option from a dropdown by its text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "optionText", Description: "The text of the option to select.", Type: stepbuilder.VarTypeString},
				{Name: "dropdownId", Description: "The ID of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user selects the option with text \"United States\" from the \"Country\" dropdown",
			Category: stepbuilder.Form,
		},
	)
}
