package table

import "cucumber/frontend/common"

type steps struct {
}

func GetSteps() []common.FrontStep {
	handlers := steps{}

	return []common.FrontStep{
		handlers.iClickOnTheRowContainingTheFollowingElements(),
		handlers.iShouldSeeRowContainingTheFollowingElements(),
		handlers.iShouldNotSeeRowContainingTheFollowingElements(),
		handlers.tableShouldContainsTheFollowingHeaders(),
	}
}
