package navigation

import (
	"etoolse/frontend/common"
	"etoolse/internal/config/testsConfig"
	"fmt"
	"strings"
)

func (n navigation) iShouldBeNavigatedToPage() common.TestStep {
	return common.NewStepWithOneVariable(
		[]string{"^I should be navigated to {string} page$"},
		func(ctx *common.TestSuiteContext) func(string) error {
			return func(pageName string) error {
				page := ctx.GetCurrentPage()
				page.WaitLoading()
				url, err := testsConfig.GetPageURL(pageName)
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
		func(pageName string) common.ValidationErrors {
			vc := common.ValidationErrors{}
			if !testsConfig.IsPageDefined(pageName) {
				vc.AddMissingPage(pageName)
			}

			return vc
		},
	)
}
