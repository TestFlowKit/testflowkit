package stepbuilder

import (
	"testflowkit/internal/steps_definitions/core/scenario"
)

type Step interface {
	GetDocumentation() Documentation
	GetSentences() []string
	GetDefinition(*scenario.Context) any
	Validate(*ValidatorContext) any
}
