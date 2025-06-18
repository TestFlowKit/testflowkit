package keyboard

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type keyboardSteps struct {
}

func GetSteps() []stepbuilder.TestStep {
	steps := keyboardSteps{}

	return []stepbuilder.TestStep{
		steps.userPressButton(),
	}
}
