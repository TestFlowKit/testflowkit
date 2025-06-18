package mouse

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.TestStep {
	handlers := steps{}

	return []stepbuilder.TestStep{
		handlers.userClicksOnButton(),
		handlers.clickOnElementWhichContains(),
		handlers.doubleClickOn(),
		handlers.userClicksOnLink(),
		handlers.userClicksOnElement(),
		handlers.doubleClickOnElementWhichContains(),
		handlers.rightClickOn(),
		handlers.rightClickOnElementWhichContains(),
		handlers.hoverOnElement(),
		handlers.hoverOnElementWhichContains(),
	}
}
