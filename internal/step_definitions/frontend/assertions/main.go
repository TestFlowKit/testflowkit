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
		handlers.theFieldValueShouldEqualsTo(),
		handlers.theFieldValueShouldNotEqualsTo(),
		handlers.radioButtonShouldBeSelectedOrNot(),
		handlers.dropdownHasValuesSelected(),
		handlers.dropdownShouldNotHaveValuesSelected(),
		handlers.elementShouldContainsText(),
		handlers.elementShouldNotContainsText(),
		handlers.elementShouldContainsExactText(),
		handlers.elementShouldNotContainsExactText(),
		handlers.elementShouldBeVisible(),
		handlers.elementShouldNotExist(),
		handlers.elementShouldNotBeVisible(),
		handlers.elementShouldExist(),
		handlers.pageTitleShouldBe(),
		handlers.pageTitleShouldNotBe(),
		handlers.currentURLShouldContain(),
		handlers.currentURLShouldNotContain(),
		handlers.elementAttributeShouldBe(),
		handlers.elementAttributeShouldNotBe(),
	}
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range GetSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
