package form

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
	"cucumber/utils"
	"fmt"
)

func (s steps) iSelectXXXIntoDropdown() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I select "{string}" into the {string} dropdown$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(options, dropdownId string) error {
				input, err := browser.GetElement(ctx.GetCurrentPage(), fmt.Sprintf("%s_dropdown", dropdownId))
				if err != nil {
					return err
				}
				ctx.GetCurrentPage()
				return input.Select(utils.String{}.SplitAndTrim(options, ","))
			}
		},
	}
}
