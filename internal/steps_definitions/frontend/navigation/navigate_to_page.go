package navigation

import (
	"fmt"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (n navigation) userNavigateToPage() stepbuilder.TestStep {
	testDefinition := func(ctx *stepbuilder.TestSuiteContext) func(string) error {
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

	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user goes to the {string} page$`},
		testDefinition,
		func(page string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !testsconfig.IsPageDefined(page) {
				vc.AddMissingPage(page)
			}

			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "Navigates to a page identified by a logical name.",
			Variables: []shared.StepVariable{
				{Name: "page", Description: "The name of the page to navigate to.", Type: shared.DocVarTypeString},
			},
			Example:  "When the user goes to the “Login” page",
			Category: shared.Navigation,
		},
	)
}
