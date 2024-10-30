package navigation

import (
	"cucumber/config"
	"cucumber/frontend/common"
	"fmt"
	"strings"
)

var iAmRedirectedToPage = common.FrontStep{
	Sentences: []string{"^I am redirected to {string} page$"},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(pageName string) error {
			page := ctx.GetCurrentPage()
			page.WaitLoading()
			url, err := config.FrontConfig{}.GetPageURL(pageName)
			if err != nil {
				return err
			}

			currentURL := page.GetInfo().URL
			if !strings.HasPrefix(currentURL, url) {
				return fmt.Errorf("redirection failed current url is %s", currentURL)
			}

			return nil
		}
	},
}
