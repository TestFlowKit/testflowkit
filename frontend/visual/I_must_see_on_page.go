package visual

import (
	"cucumber/frontend/common"
	"fmt"
	"strings"
)

var iMustSeeOnThePage = common.FrontStep{
	Sentences: []string{`^I must see {string} on the page$`},
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
