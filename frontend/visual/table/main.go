package table

import "etoolse/frontend/common"

type steps struct {
}

func GetSteps() []common.TestStep {
	handlers := steps{}

	return []common.TestStep{
		handlers.iClickOnTheRowContainingTheFollowingElements(),
		handlers.iShouldSeeRowContainingTheFollowingElements(),
		handlers.iShouldNotSeeRowContainingTheFollowingElements(),
		handlers.tableShouldContainsTheFollowingHeaders(),
	}
}
