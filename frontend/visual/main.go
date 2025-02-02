package visual

import (
	"etoolse/frontend/common"
	"etoolse/frontend/visual/table"
	"slices"
)

type steps struct {
}

func GetSteps() []common.TestStep {
	handlers := steps{}

	var otherSteps = []common.TestStep{
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
