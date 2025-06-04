package form

import (
	"testflowkit/internal/steps_definitions/core"
)

type steps struct {
}

func GetSteps() []core.TestStep {
	handlers := steps{}

	return []core.TestStep{
		handlers.userEntersTextIntoField(),
		handlers.iSelectXXXIntoDropdown(),
		handlers.checkCheckboxStatus(),
		handlers.theFieldShouldContains(),
		handlers.radioButtonShouldBeSelectedOrNot(),
		handlers.dropdownHaveValuesSelected(),
		handlers.userChecksCheckbox(),
		handlers.userUnchecksCheckbox(),
		handlers.userSelectsRadioButton(),
		handlers.userClearsFormField(),
	}
}
