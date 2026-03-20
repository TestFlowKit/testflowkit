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
		s.setHeader(),
		s.storeResponseData(),
		s.validateStatusCode(),
		s.validateStatusCodeNot(),
		s.validateJSONPathExists(),
		s.validateJSONPathNotExists(),
		s.validateJSONPathValue(),
		s.validateJSONPathValueNot(),
		s.validateJSONPathContains(),
		s.validateJSONPathNotContains(),
		s.validateJSONPathType(),
		s.validateJSONPathNotType(),
		s.validateJSONPathPattern(),
		s.validateJSONPathNotPattern(),
		s.validateJSONBodyEquals(),
		s.validateJSONBodyContains(),
		s.validateJSONBodyNotContains(),
		s.validateResponseHeaderEquals(),
		s.validateResponseHeaderNotEquals(),
		s.validateResponseHeaderPattern(),
		s.validateResponseHeaderNotPattern(),
		s.debugAPIResponse(),
	}
}
