package table

import (
	"testflowkit/internal/steps_definitions/core"
)

type steps struct {
}

func GetSteps() []core.TestStep {
	handlers := steps{}

	return []core.TestStep{
		handlers.clickOnTheRowContainingTheFollowingElements(),
		handlers.shouldSeeRowContainingTheFollowingElements(),
		handlers.shouldNotSeeRowContainingTheFollowingElements(),
		handlers.tableShouldContainsTheFollowingHeaders(),
	}
}
