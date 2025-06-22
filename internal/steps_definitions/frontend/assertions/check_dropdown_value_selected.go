package assertions

import (
	"fmt"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (s steps) dropdownHasValuesSelected() stepbuilder.Step {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}

	doc := stepbuilder.DocParams{
		Description: "checks if the dropdown has the specified values selected.",
		Variables: []stepbuilder.DocVariable{
			{Name: "dropdownId", Description: "The id of the dropdown.", Type: stepbuilder.VarTypeString},
			{Name: "optionLabels", Description: "The labels of the options to check.", Type: stepbuilder.VarTypeString},
		},
		Example:  `Then the "country" dropdown should have "USA,Canada" selected`,
		Category: stepbuilder.Form,
	}

	return stepbuilder.NewWithTwoVariables(
		[]string{`^the {string} dropdown should have "{string}" selected$`},
		func(ctx *scenario.Context) func(string, string) error {
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
		func(dropdownId, _ string) stepbuilder.ValidationErrors {
			vErr := stepbuilder.ValidationErrors{}
			label := formatVar(dropdownId)
			if !testsconfig.IsElementDefined(label) {
				vErr.AddMissingElement(label)
			}

			return vErr
		},
		doc,
	)
}
