package navigation

import (
	"cucumber/config"
	"cucumber/frontend/common"
	"log"
)

func (n navigation) iNavigateToPage() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^I navigate to {string} page$`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(page string) error {
				url, err := config.FrontConfig{}.GetPageURL(page)
				if err != nil {
					log.Fatalf("Url for page %s not configured", page)
					return err
				}
				ctx.OpenNewPage(url)
				return nil
			}
		},
	}
}
