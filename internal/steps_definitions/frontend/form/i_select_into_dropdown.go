package form

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) iSelectXXXIntoDropdown() core.TestStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}

	return core.NewStepWithTwoVariables(
		[]string{`^I select "{string}" into the {string} dropdown$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(options, dropdownId string) error {
				input, err := browser.GetElementByLabel(ctx.GetCurrentPage(), formatVar(dropdownId))
				if err != nil {
					return err
				}
				return input.Select(stringutils.SplitAndTrim(options, ","))
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
		core.StepDefDocParams{
			Description: "selects the specified options into the dropdown.",
			Variables: []shared.StepVariable{
				{Name: "options", Description: "The options to select.", Type: shared.DocVarTypeString},
				{Name: "dropdownId", Description: "The id of the dropdown.", Type: shared.DocVarTypeString},
			},
			Example:  `When I select "USA,Canada" into the "country" dropdown`,
			Category: shared.Form,
		},
	)
}
