package navigation

import (
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
	"fmt"
	"strings"
)

func (n navigation) iShouldBeNavigatedToPage() core.TestStep {
	return core.NewStepWithOneVariable(
		[]string{"^I should be navigated to {string} page$"},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(pageName string) error {
				page := ctx.GetCurrentPage()
				page.WaitLoading()
				url, err := testsconfig.GetPageURL(pageName)
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
		func(pageName string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsPageDefined(pageName) {
				vc.AddMissingPage(pageName)
			}

			return vc
		},
	)
}
