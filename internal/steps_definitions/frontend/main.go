package frontend

import (
	"slices"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/steps_definitions/frontend/assertions"
	"testflowkit/internal/steps_definitions/frontend/form"
	"testflowkit/internal/steps_definitions/frontend/keyboard"
	"testflowkit/internal/steps_definitions/frontend/mouse"
	"testflowkit/internal/steps_definitions/frontend/navigation"
	"testflowkit/internal/steps_definitions/frontend/utils"
	"testflowkit/internal/steps_definitions/frontend/visual"
)

func GetAllSteps() []stepbuilder.Step {
	finalSteps := slices.Concat(
		form.GetSteps(),
		keyboard.GetSteps(),
		navigation.GetSteps(),
		visual.GetSteps(),
		mouse.GetSteps(),
		assertions.GetSteps(),
	)

	for idx, step := range finalSteps {
		finalSteps[idx] = utils.NewUpdatePageNameDecorator(step)
	}
	return finalSteps
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range GetAllSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
