package stepbuilder

type Step interface {
	GetDocumentation() Documentation
	GetSentences() []string
	GetDefinition() any
	Validate(*ValidatorContext) any
}
