package assertions

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.TestStep {
	handlers := steps{}

	return []stepbuilder.TestStep{
		handlers.checkCheckboxStatus(),
		handlers.theFieldShouldContains(),
		handlers.radioButtonShouldBeSelectedOrNot(),
		handlers.dropdownHaveValuesSelected(),
		handlers.userShouldBeNavigatedToPage(),
		handlers.elementShouldContainText(),
		handlers.elementShouldNotContainText(),
		handlers.elementShouldContainExactText(),
	}
}
