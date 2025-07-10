package mouse

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.Step {
	handlers := steps{}

	return []stepbuilder.Step{
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
