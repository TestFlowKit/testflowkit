package form

import (
	"etoolse/config"
	"etoolse/frontend/common"
	"etoolse/utils"
	"fmt"
)

func (s steps) dropdownHaveValuesSelected() common.FrontStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}
	return common.NewStepWithTwoVariables(
		[]string{`^the {string} dropdown should have "{string}" selected$`},
		func(ctx *common.TestSuiteContext) func(string, string) error {
			return func(dropdownId, optionLabels string) error {
				selector, err := config.FrontConfig{}.GetHTMLElementSelectors(formatVar(dropdownId))
				if err != nil {
					return err
				}

				labels := utils.String{}.SplitAndTrim(optionLabels, ",")

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
		func(dropdownId, _ string) common.ValidationErrors {
			vErr := common.ValidationErrors{}
			label := formatVar(dropdownId)
			if !config.IsElementDefined(label) {
				vErr.AddMissingElement(label)
			}

			return vErr
		},
	)
}
