package stepbuilder

import (
	"context"
)

type stepThreeVars[T supportedTypes, U supportedTypes, V supportedTypes] struct {
	sentences  []string
	definition func(context.Context, T, U, V) (context.Context, error)
	validator  func(T, U, V) ValidationErrors
	doc        DocParams
}

func (s stepThreeVars[T, U, V]) GetDocumentation() Documentation {
	return Documentation{
		Sentence:    s.sentences[0],
		Description: s.doc.Description,
		Example:     s.doc.Example,
		Category:    s.doc.Category,
		Variables:   s.doc.Variables,
	}
}

func (s stepThreeVars[T, U, V]) GetSentences() []string {
	return s.sentences
}

func (s stepThreeVars[T, U, V]) GetDefinition() any {
	return s.definition
}

func (s stepThreeVars[T, U, V]) Validate(vc *ValidatorContext) any {
	return func(t T, u U, v V) {
		if s.validator == nil {
			return
		}

		validations := s.validator(t, u, v)
		if validations.HasErrors() {
			vc.AddValidationErrors(validations)
		}
	}
}

func NewWithThreeVariables[T supportedTypes, U supportedTypes, V supportedTypes](sentences []string,
	definition func(context.Context, T, U, V) (context.Context, error),
	validator func(T, U, V) ValidationErrors,
	documentation DocParams,
) Step {
	return stepThreeVars[T, U, V]{
		sentences,
		definition,
		validator,
		documentation,
	}
}
