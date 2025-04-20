package mouse

import (
	"testflowkit/internal/steps_definitions/core"
)

type steps struct {
}

func GetSteps() []core.TestStep {
	handlers := steps{}

	return []core.TestStep{
		handlers.iClickOn(),
		handlers.iClickOnElementWhichContains(),
		handlers.iDoubleClickOn(),
		handlers.iDoubleClickOnElementWhichContains(),
		handlers.iRightClickOn(),
		handlers.iRightClickOnElementWhichContains(),
		handlers.iHoverOnElement(),
		handlers.iHoverOnElementWhichContains(),
	}
}
