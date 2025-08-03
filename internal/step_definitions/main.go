package stepdefinitions

import (
	"slices"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/frontend"
	"testflowkit/internal/step_definitions/restapi"
	"testflowkit/internal/step_definitions/variables"
)

func GetAll() []stepbuilder.Step {
	allSteps := slices.Concat(
		frontend.GetAllSteps(),
		restapi.GetAllSteps(),
		restapi.GetAssertionSteps(),
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
