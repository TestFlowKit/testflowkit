package visual

import (
	"etoolse/frontend/common"
	"fmt"
	"strings"
)

func (s steps) iShouldSeeOnPage() common.FrontStep {
	return common.NewStepWithOneVariable(
		[]string{`^I should see "{string}" on the page$`},
		func(ctx *common.TestSuiteContext) func(string) error {
			return func(word string) error {
				elt, err := ctx.GetCurrentPage().GetOneBySelector("body")
				if err != nil {
					return err
				}
				if !strings.Contains(elt.TextContent(), word) {
					return fmt.Errorf("%s should be visible", word)
				}
				return nil
			}
		},
		nil,
	)
}
