package assertions

import (
	"testflowkit/internal/steps_definitions/core"
)

type steps struct {
}

func GetSteps() []core.TestStep {
	handlers := steps{}

	return []core.TestStep{
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
