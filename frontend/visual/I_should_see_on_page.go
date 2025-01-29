package visual

import (
	"etoolse/frontend/common"
	"fmt"
	"strings"
)

func (s steps) iShouldSeeOnPage() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I should see {string} on the page$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(word string) error {
				elt, err := ctx.GetCurrentPage().GetOneBySelector("body")
				if err != nil {
					return err
				}
				if !strings.Contains(elt.TextContent(), word) {
					return fmt.Errorf("%s not found", word)
				}
				return nil
			}
		},
	}
}
