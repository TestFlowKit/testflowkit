package keyboard

import "etoolse/frontend/common"

type keyboardSteps struct {
}

func GetSteps() []common.FrontStep {
	steps := keyboardSteps{}

	return []common.FrontStep{
		steps.iPressButton(),
	}
}
