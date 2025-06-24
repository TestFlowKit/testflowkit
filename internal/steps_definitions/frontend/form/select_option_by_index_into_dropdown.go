package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) userSelectOptionByIndexIntoDropdown() stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user selects the option at index {number} from the {string} dropdown$`},
		func(ctx *scenario.Context) func(int, string) error {
			return func(index int, dropdownId string) error {
				currentPage, pageName := ctx.GetCurrentPage()
				input, err := browser.GetElementByLabel(currentPage, pageName, dropdownId+"_dropdown")
				if err != nil {
					return err
				}

				return input.SelectByIndex(index)
			}
		},
		func(_ int, dropdownId string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := dropdownId + "_dropdown"
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "Selects an option from a dropdown list based on its index.",
			Variables: []stepbuilder.DocVariable{
				{Name: "index", Description: "The index of the option to select.", Type: stepbuilder.VarTypeInt},
				{Name: "name", Description: "The logical name of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user selects the option at index 2 from the "Country" dropdown`,
			Category: stepbuilder.Form,
		},
	)
}
