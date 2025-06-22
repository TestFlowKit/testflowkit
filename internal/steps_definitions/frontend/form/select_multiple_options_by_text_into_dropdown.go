package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) userSelectMultipleOptionsWithTextsIntoDropdown() stepbuilder.Step {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`^the user selects the options with text {string} from the {string} dropdown$`},
		func(ctx *scenario.Context) func(string, string) error {
			return func(options, dropdownId string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatVar(dropdownId))
				if err != nil {
					return err
				}
				return input.SelectByText(stringutils.SplitAndTrim(options, ","))
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
		stepbuilder.DocParams{
			Description: "Selects multiple options from a dropdown list based on their visible text.",
			Variables: []stepbuilder.DocVariable{
				{Name: "options", Description: "The options to select.", Type: stepbuilder.VarTypeString},
				{Name: "name", Description: "The logical name of the dropdown.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user selects the options with text "Konoha,Hidden Leaf Village" from the "Country" dropdown`,
			Category: stepbuilder.Form,
		},
	)
}
