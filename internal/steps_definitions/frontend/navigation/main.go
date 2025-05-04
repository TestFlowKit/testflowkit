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
		handlers.iSwitchToTheWindowForPage(),
		handlers.iWaitForNewWindow(),
		handlers.iSwitchToNewWindow(),
		handlers.iSwitchToOriginalWindow(),
	}
}
