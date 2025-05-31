package navigation

import (
	"errors"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (n navigation) userNavigateToURL() core.TestStep {
	testDefinition := func(ctx *core.TestSuiteContext) func(string) error {
		return func(URL string) error {
			if ctx.GetCurrentPage() == nil {
				return errors.New("no browser opened")
			}
			ctx.OpenNewPage(URL)
			return nil
		}
	}

	return core.NewStepWithOneVariable(
		[]string{`^the user navigate to the URL {string}$`},
		testDefinition,
		nil,
		core.StepDefDocParams{
			Description: "directs the browser to open the specified absolute URL",
			Variables: []shared.StepVariable{
				{Name: "URL", Description: "the absolute URL", Type: shared.DocVarTypeString},
			},
			Example:  "When the user navigates to the URL “https://myapp.com/login",
			Category: shared.Navigation,
		},
	)
}
