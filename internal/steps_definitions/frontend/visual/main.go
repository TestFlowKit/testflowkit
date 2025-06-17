package visual

import (
	"slices"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/steps_definitions/frontend/visual/table"
)

type steps struct {
}

func GetSteps() []core.TestStep {
	handlers := steps{}

	var otherSteps = []core.TestStep{
		handlers.elementShouldBeVisible(),
		handlers.elementShouldNotBeVisible(),
		handlers.elementShouldNotExist(),
		handlers.elementShouldExist(),
		handlers.shouldSeeOnPage(),
		handlers.shouldNotSeeOnPage(),
		handlers.shouldSeeElementWhichContains(),
		handlers.shouldSeeOnPageXElements(),
		handlers.shouldSeeDetailsOnPage(),
		handlers.scrollToElement(),
	}
	return slices.Concat(table.GetSteps(), otherSteps)
}
