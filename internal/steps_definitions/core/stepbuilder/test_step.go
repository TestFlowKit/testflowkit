package stepbuilder

import "testflowkit/shared"

type TestStep interface {
	GetDocumentation() shared.StepDocumentation
	GetSentences() []string
	GetDefinition(*TestSuiteContext) any
	Validate(*ValidatorContext) any
}
