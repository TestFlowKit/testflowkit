package mouse

import (
	"testflowkit/internal/steps_definitions/core"
)

type steps struct {
}

func GetSteps() []core.TestStep {
	handlers := steps{}

	return []core.TestStep{
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
