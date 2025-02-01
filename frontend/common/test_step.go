package common

type FrontStep interface {
	GetSentences() []string
	GetDefinition(*TestSuiteContext) any
	Validate(*ValidatorContext) any
}
