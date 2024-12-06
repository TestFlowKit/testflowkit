package visual

import (
	"cucumber/frontend/common"
	"cucumber/frontend/visual/table"
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
		handlers.iShouldSeeElementWitchContains(),
		handlers.iShouldSeeXElements(),
		handlers.iShouldSeeDetailsOnPage(),
	}
	return slices.Concat(table.GetSteps(), otherSteps)
}
