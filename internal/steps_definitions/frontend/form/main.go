package form

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.Step {
	handlers := steps{}

	return []stepbuilder.Step{
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
