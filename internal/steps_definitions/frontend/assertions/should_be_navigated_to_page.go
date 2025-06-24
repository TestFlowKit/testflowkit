package assertions

import (
	"fmt"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) userShouldBeNavigatedToPage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{"^the user should be navigated to {string} page$"},
		func(ctx *scenario.Context) func(string) error {
			return func(pageName string) error {
				const maxRetries = 10
				page := ctx.GetCurrentPageOnly()
				page.WaitLoading()

				var url string
				var err error
				var currentURL string

				appConfig := ctx.GetConfig()
				for range maxRetries {
					url, err = appConfig.GetFrontendURL(pageName)
					if err != nil {
						return err
					}
					page = ctx.GetCurrentPageOnly()
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
			if !config.IsPageDefined(pageName) {
				vc.AddMissingPage(pageName)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "checks if the user is navigated to a page.",
			Variables: []stepbuilder.DocVariable{
				{Name: "pageName", Description: "The name of the page to navigate to.", Type: stepbuilder.VarTypeString},
			},
			Example:  "Then I should be navigated to \"Home\" page",
			Category: stepbuilder.Navigation,
		},
	)
}
