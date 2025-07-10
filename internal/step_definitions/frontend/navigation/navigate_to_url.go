package navigation

import (
	"context"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) userNavigateToURL() stepbuilder.Step {
	testDefinition := func(ctx context.Context, url string) (context.Context, error) {
		scenarioCtx := scenario.MustFromContext(ctx)
		scenarioCtx.OpenNewPage(url)
		return ctx, nil
	}

	return stepbuilder.NewWithOneVariable(
		[]string{`the user navigate to the URL {string}`},
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
