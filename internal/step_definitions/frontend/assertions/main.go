package assertions

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
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
		handlers.elementShouldContainsText(),
		handlers.elementShouldNotContainsText(),
		handlers.elementShouldContainsExactText(),
		handlers.elementShouldBeVisible(),
		handlers.elementShouldNotExist(),
		handlers.elementShouldNotBeVisible(),
		handlers.elementShouldExist(),
		handlers.pageTitleShouldBe(),
		handlers.currentURLShouldContain(),
		handlers.elementAttributeShouldBe(),
	}
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range GetSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
