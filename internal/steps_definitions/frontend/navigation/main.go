package navigation

import (
	"testflowkit/internal/steps_definitions/core"
)

type navigation struct {
}

func GetSteps() []core.TestStep {
	handlers := navigation{}

	return []core.TestStep{
		handlers.iShouldBeNavigatedToPage(),
		handlers.iNavigateToPage(),
		handlers.iWait(),
		handlers.iRefreshPage(),
		handlers.iNavigateBack(),
		handlers.iOpenNewBrowserTab(),
		handlers.iOpenNewPrivateBrowserTab(),
		// TODO: window handling e2e tests
		handlers.iWaitAMomentForNewWindow(),
		handlers.iSwitchToMostRecentlyOpenedWindow(),
		handlers.iSwitchToOriginalWindow(),
		handlers.iSwitchToNewOpenedWindow(),
	}
}
