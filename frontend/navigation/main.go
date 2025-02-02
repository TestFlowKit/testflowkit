package navigation

import "etoolse/frontend/common"

type navigation struct {
}

func GetSteps() []common.TestStep {
	handlers := navigation{}

	return []common.TestStep{
		handlers.iShouldBeNavigatedToPage(),
		handlers.iNavigateToPage(),
		handlers.iOpenNewBrowserTab(),
		handlers.iOpenNewPrivateBrowserTab(),
	}
}
