package navigation

import (
	"context"
	"fmt"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

func (steps) userIsOnHomepage() stepbuilder.Step {
	const descriptionContext = "indicating that the user begins on the application's primary or default page"
	const moreDetails = "It assumes a predefined base URL for the \"homepage.\""
	return stepbuilder.NewWithNoVariables(
		[]string{"the user is on the homepage"},
		func(ctx context.Context) (context.Context, error) {
			const settingsVariable = "homepage"
			scenarioCtx := scenario.MustFromContext(ctx)
			url, err := scenarioCtx.GetConfig().GetFrontendURL(settingsVariable)
			if err != nil {
				logger.Fatal(fmt.Sprintf("Url for page %s not configured", settingsVariable), err)
				return ctx, err
			}
			scenarioCtx.OpenNewPage(url)
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: fmt.Sprintf("establishes the initial context, %s %s", descriptionContext, moreDetails),
			Variables:   nil,
			Example:     "Given the user is on the homepage",
			Category:    stepbuilder.Navigation,
		},
	)
}
