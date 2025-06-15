package assertions

import (
	"fmt"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/shared"
)

func (s steps) dropdownHaveValuesSelected() core.TestStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}

	doc := core.StepDefDocParams{
		Description: "checks if the dropdown has the specified values selected.",
		Variables: []shared.StepVariable{
			{Name: "dropdownId", Description: "The id of the dropdown.", Type: shared.DocVarTypeString},
			{Name: "optionLabels", Description: "The labels of the options to check.", Type: shared.DocVarTypeString},
		},
		Example:  `Then the "country" dropdown should have "USA,Canada" selected`,
		Category: shared.Form,
	}

	return core.NewStepWithTwoVariables(
		[]string{`^the {string} dropdown should have "{string}" selected$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(dropdownId, optionLabels string) error {
				selector, err := testsconfig.GetHTMLElementSelectors(formatVar(dropdownId))
				if err != nil {
					return err
				}

				labels := stringutils.SplitAndTrim(optionLabels, ",")

				result := ctx.GetCurrentPage().ExecuteJS(`(selector, labels) => {
					const selectedOpts = Array.from(document.querySelector(selector).selectedOptions).map(opt => opt.label)
					return labels.every(label => selectedOpts.includes(label))
				}`, selector, labels)

				if result == "true" {
					return nil
				}
				return fmt.Errorf("%s value is not selected in %s dropdown", optionLabels, dropdownId)
			}
		},
		func(dropdownId, _ string) core.ValidationErrors {
			vErr := core.ValidationErrors{}
			label := formatVar(dropdownId)
			if !testsconfig.IsElementDefined(label) {
				vErr.AddMissingElement(label)
			}

			return vErr
		},
		doc,
	)
}
