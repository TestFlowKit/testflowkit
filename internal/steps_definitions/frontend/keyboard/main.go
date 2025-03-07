package keyboard

import (
	"testflowkit/internal/steps_definitions/core"
)

type keyboardSteps struct {
}

func GetSteps() []core.TestStep {
	steps := keyboardSteps{}

	return []core.TestStep{
		steps.iPressButton(),
	}
}
