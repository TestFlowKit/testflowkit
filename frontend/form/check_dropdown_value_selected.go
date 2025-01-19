package form

import (
	"etoolse/config"
	"etoolse/frontend/common"
	"etoolse/utils"
	"fmt"
)

func (s steps) dropdownHaveValuesSelected() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^the {string} dropdown should have "{string}" selected$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(dropdownId, optionLabels string) error {
				selector, err := config.FrontConfig{}.GetHTMLElementSelectors(fmt.Sprintf("%s_dropdown", dropdownId))
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
	}
}
