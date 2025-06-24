package navigation

import (
	"errors"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (n navigation) userNavigateToURL() stepbuilder.Step {
	testDefinition := func(ctx *scenario.Context) func(string) error {
		return func(URL string) error {
			if ctx.GetCurrentPageOnly() == nil {
				return errors.New("no browser opened")
			}
			ctx.OpenNewPage(URL)
			return nil
		}
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`^the user navigate to the URL {string}$`},
		testDefinition,
		nil,
		stepbuilder.DocParams{
			Description: "directs the browser to open the specified absolute URL",
			Variables: []stepbuilder.DocVariable{
				{Name: "URL", Description: "the absolute URL", Type: stepbuilder.VarTypeString},
			},
			Example:  "When the user navigates to the URL “https://myapp.com/login",
			Category: stepbuilder.Navigation,
		},
	)
}
