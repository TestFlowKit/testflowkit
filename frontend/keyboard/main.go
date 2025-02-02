package keyboard

import "etoolse/frontend/common"

type keyboardSteps struct {
}

func GetSteps() []common.TestStep {
	steps := keyboardSteps{}

	return []common.TestStep{
		steps.iPressButton(),
	}
}
