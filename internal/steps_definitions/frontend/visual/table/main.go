package table

import (
	"testflowkit/internal/steps_definitions/core"
)

type steps struct {
}

func GetSteps() []core.TestStep {
	handlers := steps{}

	return []core.TestStep{
		// TODO: a tester tout ca pour etre sur qu'on est bons une fois pour toute
		handlers.iClickOnTheRowContainingTheFollowingElements(),
		handlers.iShouldSeeRowContainingTheFollowingElements(),
		handlers.iShouldNotSeeRowContainingTheFollowingElements(),
		handlers.tableShouldContainsTheFollowingHeaders(),
	}
}
