package keyboard

import "cucumber/frontend/common"

type keyboardSteps struct {
}

func GetSteps() []common.FrontStep {
	steps := keyboardSteps{}

	return []common.FrontStep{
		steps.iPressButton(),
	}
}
