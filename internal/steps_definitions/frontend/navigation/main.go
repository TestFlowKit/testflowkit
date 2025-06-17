package navigation

import (
	"testflowkit/internal/steps_definitions/core"
)

type navigation struct {
}

func GetSteps() []core.TestStep {
	handlers := navigation{}

	return []core.TestStep{
		handlers.userNavigateToPage(),
		handlers.userWait(),
		handlers.refreshPage(),
		handlers.theUserNavigateBack(),
		handlers.userNavigateToURL(),
		handlers.openNewBrowserTab(),
		handlers.openNewPrivateBrowserTab(),
		handlers.userIsOnHomepage(),
		// TODO: window handling e2e tests
		handlers.waitAMomentForNewWindow(),
		handlers.switchToMostRecentlyOpenedWindow(),
		handlers.switchToOriginalWindow(),
		handlers.switchToNewOpenedWindow(),
	}
}
