package assertions

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct {
}

func GetSteps() []stepbuilder.Step {
	handlers := steps{}

	return []stepbuilder.Step{
		handlers.checkResponseStatusCode(),
		handlers.responseBodyShouldContain(),
		handlers.responseBodyPathShouldExist(),
	}
}
