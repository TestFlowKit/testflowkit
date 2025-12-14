package backend

import (
	"testflowkit/internal/step_definitions/backend/commonbackendsteps"
	"testflowkit/internal/step_definitions/backend/restapi"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func GetAllSteps() []stepbuilder.Step {
	var allSteps []stepbuilder.Step
	allSteps = append(allSteps, commonbackendsteps.GetSteps()...)
	allSteps = append(allSteps, restapi.GetSteps()...)
	return allSteps
}

func GetDocs() []stepbuilder.Documentation {
	var docs []stepbuilder.Documentation
	for _, step := range GetAllSteps() {
		docs = append(docs, step.GetDocumentation())
	}
	return docs
}
