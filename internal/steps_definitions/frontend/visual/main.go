package visual

import (
	"etoolse/internal/steps_definitions/core"
	"etoolse/internal/steps_definitions/frontend/visual/table"
	"slices"
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
		handlers.iShouldSeeOnPage(),
		handlers.iShouldNotSeeOnPage(),
		handlers.iShouldSeeElementWhichContains(),
		handlers.iShouldSeeOnPageXElements(),
		handlers.iShouldSeeDetailsOnPage(),
	}
	return slices.Concat(table.GetSteps(), otherSteps)
}
