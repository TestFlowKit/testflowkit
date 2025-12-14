package restapi

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct{}

// GetSteps returns all REST API-specific steps for documentation.
func GetSteps() []stepbuilder.Step {
	s := steps{}
	return []stepbuilder.Step{
		s.setQueryParams(),
		s.setPathParams(),
		s.setRequestBody(),
		s.setJSONBody(),
		s.debugRequest(),
	}
}
