package common

type FrontStepDefinition = any
type FrontStepDefinitionClosure = func(*TestSuiteContext) FrontStepDefinition

type FrontStep struct {
	Sentences  []string
	Definition FrontStepDefinitionClosure
}
