package frontend

import (
	"slices"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/frontend/assertions"
	"testflowkit/internal/step_definitions/frontend/form"
	"testflowkit/internal/step_definitions/frontend/keyboard"
	"testflowkit/internal/step_definitions/frontend/mouse"
	"testflowkit/internal/step_definitions/frontend/navigation"
	"testflowkit/internal/step_definitions/frontend/visual"
)

func GetAllSteps() []stepbuilder.Step {
	return slices.Concat(
		form.GetSteps(),
		keyboard.GetSteps(),
		navigation.GetSteps(),
		visual.GetSteps(),
		mouse.GetSteps(),
		assertions.GetSteps(),
	)
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range GetAllSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
