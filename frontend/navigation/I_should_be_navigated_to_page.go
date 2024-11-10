package navigation

import (
	"cucumber/config"
	"cucumber/frontend/common"
	"fmt"
	"strings"
)

func (n navigation) iShouldBeNavigatedToPage() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{"^I should be navigated to {string} page$"},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(pageName string) error {
				page := ctx.GetCurrentPage()
				page.WaitLoading()
				url, err := config.FrontConfig{}.GetPageURL(pageName)
				if err != nil {
					return err
				}

				currentURL := page.GetInfo().URL
				if strings.HasPrefix(currentURL, url) || strings.HasPrefix(url, currentURL) {
					return nil
				}

				return fmt.Errorf("navigation check failed: current url is %s but %s expected", currentURL, url)
			}
		},
	}
}
