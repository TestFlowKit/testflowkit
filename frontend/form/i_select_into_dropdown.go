package form

import (
	"etoolse/config"
	"etoolse/frontend/common"
	"etoolse/frontend/common/browser"
	"etoolse/utils"
	"fmt"
)

func (s steps) iSelectXXXIntoDropdown() common.FrontStep {
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
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			return vc
		},
	)
}
