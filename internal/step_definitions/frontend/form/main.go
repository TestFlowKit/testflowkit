package form

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.Step {
	handlers := steps{}

	return []stepbuilder.Step{
		handlers.userEntersTextIntoField(),
		handlers.selectOptionWithTextIntoDropdown(),
		handlers.selectMultipleOptionsByTextIntoDropdown(),
		handlers.userSelectMultipleOptionsByValueIntoDropdown(),
		handlers.userSelectOptionWithValueIntoDropdown(),
		handlers.userSelectOptionByIndexIntoDropdown(),
		handlers.checkCheckbox(),
		handlers.uncheckCheckbox(),
		handlers.selectRadioButton(),
		handlers.clearField(),
		handlers.userUploadsFileIntoField(),
		handlers.userUploadsMultipleFiles(),
	}
}
