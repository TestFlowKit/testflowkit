package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) userSelectOptionByIndexIntoDropdown() core.TestStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}

	return core.NewStepWithTwoVariables(
		[]string{`^the user selects the option at index {number} from the {string} dropdown$`},
		func(ctx *core.TestSuiteContext) func(int, string) error {
			return func(optionIndex int, dropdownId string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatVar(dropdownId))
				if err != nil {
					return err
				}
				return input.SelectByIndex(optionIndex)
			}
		},
		func(_ int, dropdownId string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			label := formatVar(dropdownId)
			if !testsconfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
		core.StepDefDocParams{
			Description: "Selects an option from a dropdown list based on its underlying ‘value’ attribute.",
			Variables: []shared.StepVariable{
				{Name: "option value", Description: "The value of the option to select.", Type: shared.DocVarTypeString},
				{Name: "name", Description: "The logical name of the dropdown.", Type: shared.DocVarTypeString},
			},
			Example:  `When the user selects the option with value “CIV” from the “Country” dropdown.`,
			Category: shared.Form,
		},
	)
}
