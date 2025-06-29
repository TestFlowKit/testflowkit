package navigation

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) userShouldBeNavigatedToPage() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{"the user should be navigated to {string} page"},
		func(ctx context.Context, pageName string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			const maxRetries = 10
			page := scenarioCtx.GetCurrentPageOnly()
			page.WaitLoading()

			var url string
			var err error
			var currentURL string

			appConfig := scenarioCtx.GetConfig()
			for range maxRetries {
				url, err = appConfig.GetFrontendURL(pageName)
				if err != nil {
					return ctx, err
				}
				page = scenarioCtx.GetCurrentPageOnly()
				currentURL = page.GetInfo().URL
				if strings.HasPrefix(currentURL, url) || strings.HasPrefix(url, currentURL) {
					return ctx, nil
				}

				page.WaitLoading()
			}

			return ctx, fmt.Errorf("navigation check failed: current url is %s but %s expected", currentURL, url)
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
