package assertions

import (
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.Step {
	handlers := steps{}

	return []stepbuilder.Step{
		handlers.checkCheckboxStatus(),
		handlers.theFieldShouldContain(),
		handlers.radioButtonShouldBeSelectedOrNot(),
		handlers.dropdownHasValuesSelected(),
		handlers.userShouldBeNavigatedToPage(),
		handlers.elementShouldContainsText(),
		handlers.elementShouldNotContainsText(),
		handlers.elementShouldContainsExactText(),
		handlers.elementShouldBeVisible(),
		handlers.elementShouldNotExist(),
		handlers.elementShouldNotBeVisible(),
		handlers.elementShouldExist(),
	}
}
