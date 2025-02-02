package form

import "etoolse/frontend/common"

type steps struct {
}

func GetSteps() []common.TestStep {
	handlers := steps{}

	return []common.TestStep{
		handlers.iTypeXXXIntoInput(),
		handlers.iSelectXXXIntoDropdown(),
		handlers.checkCheckboxStatus(),
		handlers.theFieldShouldContains(),
		handlers.dropdownHaveValuesSelected(),
	}
}
