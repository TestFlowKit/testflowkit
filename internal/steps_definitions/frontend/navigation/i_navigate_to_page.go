package navigation

import (
	"fmt"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (n navigation) iNavigateToPage() core.TestStep {
	testDefinition := func(ctx *core.TestSuiteContext) func(string) error {
		return func(page string) error {
			url, err := testsconfig.GetPageURL(page)
			if err != nil {
				logger.Fatal(fmt.Sprintf("Url for page %s not configured", page), err)
				return err
			}
			ctx.OpenNewPage(url)
			return nil
		}
	}

	return core.NewStepWithOneVariable[string](
		[]string{`^I navigate to {string} page$`},
		testDefinition,
		func(page string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsPageDefined(page) {
				vc.AddMissingPage(page)
			}

			return vc
		},
		core.StepDefDocParams{
			Description: "navigates to a page.",
			Variables: []shared.StepVariable{
				{Name: "page", Description: "The name of the page to navigate to.", Type: shared.DocVarTypeString},
			},
			Example:  "When I navigate to \"Home\" page",
			Category: shared.Navigation,
		},
	)
}
