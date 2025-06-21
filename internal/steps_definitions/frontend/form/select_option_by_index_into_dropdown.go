package form

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) userSelectOptionByIndexIntoDropdown() stepbuilder.TestStep {
	return stepbuilder.NewStepWithTwoVariables(
		[]string{`^the user selects the option at index {number} from the {string} dropdown$`},
		func(ctx *scenario.Context) func(int, string) error {
			return func(index int, dropdownId string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), dropdownId+"_dropdown")
				if err != nil {
					return err
				}
				return input.SelectByIndex(index)
			}
		},
		func(_ int, dropdownId string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := dropdownId + "_dropdown"
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "Selects an option from a dropdown list based on its index.",
			Variables: []shared.StepVariable{
				{Name: "index", Description: "The index of the option to select.", Type: shared.DocVarTypeInt},
				{Name: "name", Description: "The logical name of the dropdown.", Type: shared.DocVarTypeString},
			},
			Example:  `When the user selects the option at index 2 from the "Country" dropdown`,
			Category: shared.Form,
		},
	)
}
