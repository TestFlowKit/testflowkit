package navigation

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type navigation struct {
}

func GetSteps() []stepbuilder.Step {
	handlers := navigation{}

	return []stepbuilder.Step{
		handlers.userNavigateToPage(),
		handlers.userWait(),
		handlers.refreshPage(),
		handlers.theUserNavigateBack(),
		handlers.userNavigateToURL(),
		handlers.openANewBrowserTab(),
		handlers.openANewPrivateBrowserTab(),
		handlers.userIsOnHomepage(),
		handlers.userShouldBeNavigatedToPage(),
		// TODO: window handling e2e tests
		handlers.waitAMomentForNewWindow(),
		handlers.switchToMostOpenedWindow(),
		handlers.switchToOriginalWindow(),
		handlers.switchToNewOpenedWindow(),
	}
}
