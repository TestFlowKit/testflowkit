package stepbuilder

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/shared"
)

type TestStep interface {
	GetDocumentation() shared.StepDocumentation
	GetSentences() []string
	GetDefinition(*scenario.Context) any
	Validate(*ValidatorContext) any
}
