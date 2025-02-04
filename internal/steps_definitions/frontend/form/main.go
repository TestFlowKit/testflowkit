package form

import (
	"etoolse/internal/steps_definitions/core"
)

type steps struct {
}

func GetSteps() []core.TestStep {
	handlers := steps{}

	return []core.TestStep{
		handlers.iTypeXXXIntoInput(),
		handlers.iSelectXXXIntoDropdown(),
		handlers.checkCheckboxStatus(),
		handlers.theFieldShouldContains(),
		handlers.dropdownHaveValuesSelected(),
	}
}
