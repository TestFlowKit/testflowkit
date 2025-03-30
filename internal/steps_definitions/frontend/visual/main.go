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
		handlers.iClickOn(),
		handlers.iClickOnElementWhichContains(),
		handlers.iDoubleClickOn(),
		handlers.iDoubleClickOnElementWhichContains(),
		handlers.iShouldSeeOnPage(),
		handlers.iShouldNotSeeOnPage(),
		handlers.iShouldSeeElementWhichContains(),
		handlers.iShouldSeeOnPageXElements(),
		handlers.iShouldSeeDetailsOnPage(),
	}
	return slices.Concat(table.GetSteps(), otherSteps)
}
