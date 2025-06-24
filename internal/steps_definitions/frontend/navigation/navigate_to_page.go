package navigation

import (
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (n navigation) userNavigateToPage() stepbuilder.Step {
	testDefinition := func(ctx *scenario.Context) func(string) error {
		return func(page string) error {
			url, err := ctx.GetConfig().GetFrontendURL(page)
			if err != nil {
				logger.Fatal(fmt.Sprintf("Url for page %s not configured", page), err)
				return err
			}
			ctx.OpenNewPage(url)
			return nil
		}
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user goes to the {string} page$`},
		testDefinition,
		func(page string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsPageDefined(page) {
				vc.AddMissingPage(page)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "Navigates to a page identified by a logical name.",
			Variables: []stepbuilder.DocVariable{
				{Name: "page", Description: "The name of the page to navigate to.", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user goes to the “Login” page",
			Category: stepbuilder.Navigation,
		},
	)
}
