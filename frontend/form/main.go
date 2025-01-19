package form

import (
	"etoolse/frontend/common"
)

type steps struct {
}

func GetSteps() []common.FrontStep {
	handlers := steps{}

	return []common.FrontStep{
		handlers.iTypeXXXIntoInput(),
		handlers.iSelectXXXIntoDropdown(),
		handlers.checkCheckboxStatus(),
		handlers.theFieldShouldContains(),
		handlers.dropdownHaveValuesSelected(),
	}
}
