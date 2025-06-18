package table

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.TestStep {
	handlers := steps{}

	return []stepbuilder.TestStep{
		handlers.clickOnTheRowContainingTheFollowingElements(),
		handlers.shouldSeeRowContainingTheFollowingElements(),
		handlers.shouldNotSeeRowContainingTheFollowingElements(),
		handlers.tableShouldContainsTheFollowingHeaders(),
	}
}
