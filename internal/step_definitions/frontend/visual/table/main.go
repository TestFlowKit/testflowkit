package table

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.Step {
	handlers := steps{}

	return []stepbuilder.Step{
		handlers.clickOnTheRowContainingTheFollowingElements(),
		handlers.shouldSeeRowContainingTheFollowingElements(),
		handlers.shouldNotSeeRowContainingTheFollowingElements(),
		handlers.tableShouldContainsTheFollowingHeaders(),
	}
}
