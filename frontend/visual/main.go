package visual

import (
	"cucumber/frontend/common"
)

type steps struct {
}

func GetSteps() []common.FrontStep {
	handlers := steps{}
	return []common.FrontStep{
		handlers.elementShouldBeVisible(),
		handlers.elementShouldNotBeVisible(),
		handlers.iClickOn(),
		handlers.iClickOnElementWhichContains(),
		handlers.iShouldSeeOnPage(),
		handlers.iShouldNotSeeOnPage(),
		handlers.iShouldSeeElementWitchContains(),
		handlers.iShouldSeeXElements(),
	}
}
