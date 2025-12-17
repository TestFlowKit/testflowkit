package restapi

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct{}

func GetSteps() []stepbuilder.Step {
	s := steps{}
	return []stepbuilder.Step{
		s.setQueryParams(),
		s.setPathParams(),
		s.setRequestBody(),
		s.setRequestBodyFromFile(),
		s.setJSONBody(),
		s.debugRequest(),
	}
}
