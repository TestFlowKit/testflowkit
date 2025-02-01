package visual

import (
	"etoolse/frontend/common"
	"etoolse/frontend/visual/table"
	"slices"
)

type steps struct {
}

func GetSteps() []common.FrontStep {
	handlers := steps{}

	var otherSteps = []common.FrontStep{
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
