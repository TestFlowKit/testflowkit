package stepbuilder

import (
	"context"
)

type stepOneVar[T supportedTypes] struct {
	sentences  []string
	definition func(context.Context, T) (context.Context, error)
	validator  func(T) ValidationErrors
	doc        DocParams
}

func (s stepOneVar[T]) GetSentences() []string {
	return s.sentences
}

func (s stepOneVar[T]) GetDefinition() any {
	return s.definition
}

func (s stepOneVar[T]) GetDocumentation() Documentation {
	return Documentation{
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

func NewWithOneVariable[T supportedTypes](
	sentences []string,
	definition func(context.Context, T) (context.Context, error),
	validator func(T) ValidationErrors,
	documentation DocParams,
) Step {
	return stepOneVar[T]{
		sentences,
		definition,
		validator,
		documentation,
	}
}
