package graphql

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct{}

func GetSteps() []stepbuilder.Step {
	s := steps{}
	return []stepbuilder.Step{
		s.setGraphQLVariables(),
		s.validateHaveErrors(),
		s.validateNoErrors(),
		s.validateErrorMessage(),
		s.storeGraphQLError(),
		s.storeGraphQLErrorMessage(),
	}
}
