package navigation

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type navigation struct {
}

func GetSteps() []stepbuilder.TestStep {
	handlers := navigation{}

	return []stepbuilder.TestStep{
		handlers.userNavigateToPage(),
		handlers.userWait(),
		handlers.refreshPage(),
		handlers.theUserNavigateBack(),
		handlers.userNavigateToURL(),
		handlers.openANewBrowserTab(),
		handlers.openANewPrivateBrowserTab(),
		handlers.userIsOnHomepage(),
		// TODO: window handling e2e tests
		handlers.waitAMomentForNewWindow(),
		handlers.switchToMostOpenedWindow(),
		handlers.switchToOriginalWindow(),
		handlers.switchToNewOpenedWindow(),
	}
}
