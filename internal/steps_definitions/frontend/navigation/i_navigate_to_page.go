package navigation

import (
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
	"log"
)

func (n navigation) iNavigateToPage() core.TestStep {
	return core.NewStepWithOneVariable[string](
		[]string{`^I navigate to {string} page$`},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(page string) error {
				url, err := testsconfig.GetPageURL(page)
				if err != nil {
					log.Fatalf("Url for page %s not configured", page)
					return err
				}
				ctx.OpenNewPage(url)
				return nil
			}
		},
		func(page string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsPageDefined(page) {
				vc.AddMissingPage(page)
			}

			return vc
		},
	)
}
