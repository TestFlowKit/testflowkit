package navigation

import (
	"fmt"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (n navigation) userIsOnHomepage() core.TestStep {
	const descriptionContext = "indicating that the user begins on the application’s primary or default page"
	const moreDetails = "It assumes a predefined base URL for the “homepage."
	return core.NewStepWithNoVariables(
		[]string{"the user is on the homepage"},
		func(ctx *core.TestSuiteContext) func() error {
			return func() error {
				const settingsVariable = "homepage"
				url, err := testsconfig.GetPageURL(settingsVariable)
				if err != nil {
					logger.Fatal(fmt.Sprintf("Url for page %s not configured", settingsVariable), err)
					return err
				}
				if ctx.GetCurrentPage() == nil {
					ctx.InitBrowser(false)
				}

				ctx.OpenNewPage(url)

				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: fmt.Sprintf("establishes the initial context, %s %s", descriptionContext, moreDetails),
			Variables:   nil,
			Example:     "Given the user is on the homepage",
			Category:    shared.Navigation,
		},
	)
}
