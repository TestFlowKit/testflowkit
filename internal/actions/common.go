package actions

import (
	"slices"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/steps_definitions/frontend"
	"testflowkit/internal/steps_definitions/restapi"
)

func GetAllSteps() []stepbuilder.Step {
	return slices.Concat(frontend.GetAllSteps(), restapi.GetAllSteps())
}
