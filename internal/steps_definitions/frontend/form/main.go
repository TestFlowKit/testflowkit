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
		handlers.userSelectOptionWithTextIntoDropdown(),
		handlers.userSelectMultipleOptionsWithTextsIntoDropdown(),
		handlers.userSelectOptionWithValueIntoDropdown(),
		handlers.userSelectMultipleOptionsByValueIntoDropdown(),
		handlers.userSelectOptionByIndexIntoDropdown(),
		handlers.userChecksCheckbox(),
		handlers.userUnchecksCheckbox(),
		handlers.userSelectsRadioButton(),
		handlers.userClearsFormField(),
	}
}
