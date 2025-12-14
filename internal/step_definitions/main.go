package stepdefinitions

import (
	"slices"
	"testflowkit/internal/step_definitions/backend"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/frontend"
	"testflowkit/internal/step_definitions/variables"
)

func GetAll() []stepbuilder.Step {
	allSteps := slices.Concat(
		frontend.GetAllSteps(),
		backend.GetAllSteps(),
		variables.GetAllSteps(),
	)

	var decoratedSteps []stepbuilder.Step
	for _, step := range allSteps {
		decoratedSteps = append(decoratedSteps,
			stepbuilder.NewVariableSubstitutionDecorator(
				stepbuilder.NewUpdatePageNameDecorator(step)))
	}

	return decoratedSteps
}
