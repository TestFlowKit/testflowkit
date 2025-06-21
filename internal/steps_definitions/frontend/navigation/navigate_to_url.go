package navigation

import (
	"errors"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (n navigation) userNavigateToURL() stepbuilder.TestStep {
	testDefinition := func(ctx *scenario.Context) func(string) error {
		return func(URL string) error {
			if ctx.GetCurrentPage() == nil {
				return errors.New("no browser opened")
			}
			ctx.OpenNewPage(URL)
			return nil
		}
	}

	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user navigate to the URL {string}$`},
		testDefinition,
		nil,
		stepbuilder.StepDefDocParams{
			Description: "directs the browser to open the specified absolute URL",
			Variables: []shared.StepVariable{
				{Name: "URL", Description: "the absolute URL", Type: shared.DocVarTypeString},
			},
			Example:  "When the user navigates to the URL “https://myapp.com/login",
			Category: shared.Navigation,
		},
	)
}
