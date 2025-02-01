package navigation

import (
	"etoolse/config"
	"etoolse/frontend/common"
	"fmt"
	"strings"
)

func (n navigation) iShouldBeNavigatedToPage() common.FrontStep {
	return common.NewStepWithOneVariable(
		[]string{"^I should be navigated to {string} page$"},
		func(ctx *common.TestSuiteContext) func(string) error {
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
		nil,
	)
}
