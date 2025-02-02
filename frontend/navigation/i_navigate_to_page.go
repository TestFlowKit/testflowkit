package navigation

import (
	"etoolse/frontend/common"
	"etoolse/internal/config/testsConfig"
	"log"
)

func (n navigation) iNavigateToPage() common.TestStep {
	return common.NewStepWithOneVariable[string](
		[]string{`^I navigate to {string} page$`},
		func(ctx *common.TestSuiteContext) func(string) error {
			return func(page string) error {
				url, err := testsConfig.GetPageURL(page)
				if err != nil {
					log.Fatalf("Url for page %s not configured", page)
					return err
				}
				ctx.OpenNewPage(url)
				return nil
			}
		},
		func(page string) common.ValidationErrors {
			vc := common.ValidationErrors{}
			if !testsConfig.IsPageDefined(page) {
				vc.AddMissingPage(page)
			}

			return vc
		},
	)
}
