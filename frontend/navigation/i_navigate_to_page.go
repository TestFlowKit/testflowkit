package navigation

import (
	"etoolse/config"
	"etoolse/frontend/common"
	"log"
)

func (n navigation) iNavigateToPage() common.FrontStep {
	return common.NewStepWithOneVariable[string](
		[]string{`^I navigate to {string} page$`},
		func(ctx *common.TestSuiteContext) func(string) error {
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
		nil,
	)
}
