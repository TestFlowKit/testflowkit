package stepbuilder

import "testflowkit/shared"

type stepOneVar[T stepSupportedTypes] struct {
	sentences  []string
	definition func(*TestSuiteContext) func(T) error
	validator  func(T) ValidationErrors
	doc        StepDefDocParams
}

func (s stepOneVar[T]) GetSentences() []string {
	return s.sentences
}

func (s stepOneVar[T]) GetDefinition(ctx *TestSuiteContext) any {
	return s.definition(ctx)
}

func (s stepOneVar[T]) GetDocumentation() shared.StepDocumentation {
	return shared.StepDocumentation{
		Sentence:    s.sentences[0],
		Description: s.doc.Description,
		Example:     s.doc.Example,
		Category:    s.doc.Category,
		Variables:   s.doc.Variables,
	}
}

func (s stepOneVar[T]) Validate(vc *ValidatorContext) any {
	return func(t T) {
		if s.validator == nil {
			return
		}

		validations := s.validator(t)
		if validations.HasErrors() {
			vc.AddValidationErrors(validations)
		}
	}
}

func NewStepWithOneVariable[T stepSupportedTypes](
	sentences []string,
	definition func(*TestSuiteContext) func(T) error,
	validator func(T) ValidationErrors,
	documentation StepDefDocParams,
) TestStep {
	return stepOneVar[T]{
		sentences,
		definition,
		validator,
		documentation,
	}
}
