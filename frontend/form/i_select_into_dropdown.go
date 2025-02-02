package form

import (
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"etoolse/internal/config/testsConfig"
	"etoolse/utils"
	"fmt"
)

func (s steps) iSelectXXXIntoDropdown() common.TestStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}

	return common.NewStepWithTwoVariables(
		[]string{`^I select "{string}" into the {string} dropdown$`},
		func(ctx *common.TestSuiteContext) func(string, string) error {
			return func(options, dropdownId string) error {
				input, err := browser.GetElement(ctx.GetCurrentPage(), formatVar(dropdownId))
				if err != nil {
					return err
				}
				ctx.GetCurrentPage()
				return input.Select(utils.String{}.SplitAndTrim(options, ","))
			}
		},
		func(_, dropdownId string) common.ValidationErrors {
			vc := common.ValidationErrors{}
			label := formatVar(dropdownId)
			if !testsConfig.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
	)
}
