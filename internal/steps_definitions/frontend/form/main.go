package form

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.TestStep {
	handlers := steps{}

	return []stepbuilder.TestStep{
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
