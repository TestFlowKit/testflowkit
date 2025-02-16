package navigation

import (
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
	"etoolse/pkg/logger"
	"fmt"
)

func (n navigation) iNavigateToPage() core.TestStep {
	return core.NewStepWithOneVariable[string](
		[]string{`^I navigate to {string} page$`},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(page string) error {
				url, err := testsconfig.GetPageURL(page)
				if err != nil {
					logger.Fatal(fmt.Sprintf("Url for page %s not configured", page), err)
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
