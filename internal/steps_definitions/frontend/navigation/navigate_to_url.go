package navigation

import (
	"context"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (steps) userNavigateToURL() stepbuilder.Step {
	testDefinition := func(scenarioCtx *scenario.Context) func(context.Context, string) (context.Context, error) {
		return func(ctx context.Context, url string) (context.Context, error) {
			scenarioCtx.OpenNewPage(url)
			return ctx, nil
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
			Example:  "When the user navigates to the URL \"https://myapp.com/login\"",
			Category: stepbuilder.Navigation,
		},
	)
}
