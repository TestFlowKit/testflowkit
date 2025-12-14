package commonbackendsteps

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct{}

func GetSteps() []stepbuilder.Step {
	s := steps{}
	return []stepbuilder.Step{
		s.prepareRequest(),
		s.sendRequest(),
		s.setHeaders(),
		s.storeResponseData(),
		s.validateStatusCode(),
		s.validateJSONPathExists(),
		s.validateJSONPathValue(),
		s.validateJSONPathContains(),
		s.validateJSONBodyEquals(),
		s.validateJSONBodyContains(),
	}
}
