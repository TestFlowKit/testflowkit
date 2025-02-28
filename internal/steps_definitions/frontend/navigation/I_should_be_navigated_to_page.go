package navigation

import (
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
	"etoolse/shared"
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
		core.StepDefDocParams{
			Description: "checks if the user is navigated to a page.",
			Variables: []shared.StepVariable{
				{Name: "pageName", Description: "The name of the page to navigate to.", Type: shared.DocVarTypeString},
			},
			Example:  "Then I should be navigated to \"Home\" page",
			Category: shared.Navigation,
		},
	)
}
