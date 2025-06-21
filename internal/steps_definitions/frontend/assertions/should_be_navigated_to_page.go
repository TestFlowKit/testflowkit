package assertions

import (
	"fmt"
	"strings"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) userShouldBeNavigatedToPage() stepbuilder.TestStep {
	return stepbuilder.NewStepWithOneVariable(
		[]string{"^the user should be navigated to {string} page$"},
		func(ctx *scenario.Context) func(string) error {
			return func(pageName string) error {
				const maxRetries = 10
				page := ctx.GetCurrentPage()
				page.WaitLoading()

				var url string
				var err error
				var currentURL string

				for range maxRetries {
					url, err = testsconfig.GetPageURL(pageName)
					if err != nil {
						return err
					}
					page = ctx.GetCurrentPage()
					currentURL = page.GetInfo().URL
					if strings.HasPrefix(currentURL, url) || strings.HasPrefix(url, currentURL) {
						return nil
					}

					page.WaitLoading()
				}

				return fmt.Errorf("navigation check failed: current url is %s but %s expected", currentURL, url)
			}
		},
		func(pageName string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !testsconfig.IsPageDefined(pageName) {
				vc.AddMissingPage(pageName)
			}

			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "checks if the user is navigated to a page.",
			Variables: []shared.StepVariable{
				{Name: "pageName", Description: "The name of the page to navigate to.", Type: shared.DocVarTypeString},
			},
			Example:  "Then I should be navigated to \"Home\" page",
			Category: shared.Navigation,
		},
	)
}
