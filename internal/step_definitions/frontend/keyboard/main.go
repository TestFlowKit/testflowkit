package keyboard

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type keyboardSteps struct {
}

func GetSteps() []stepbuilder.Step {
	steps := keyboardSteps{}

	return []stepbuilder.Step{
		steps.userPressButton(),
	}
}
