package stepbuilder

import (
	"context"
	"testflowkit/internal/steps_definitions/core/scenario"
)

type stepTwoVars[T supportedTypes, U supportedTypes] struct {
	sentences  []string
	definition func(*scenario.Context) func(context.Context, T, U) (context.Context, error)
	validator  func(T, U) ValidationErrors
	doc        DocParams
}

func (s stepTwoVars[T, U]) GetDocumentation() Documentation {
	return Documentation{
		Sentence:    s.sentences[0],
		Description: s.doc.Description,
		Example:     s.doc.Example,
		Category:    s.doc.Category,
		Variables:   s.doc.Variables,
	}
}

func (s stepTwoVars[T, U]) GetSentences() []string {
	return s.sentences
}

func (s stepTwoVars[T, U]) GetDefinition(ctx *scenario.Context) any {
	return s.definition(ctx)
}

func (s stepTwoVars[T, U]) Validate(vc *ValidatorContext) any {
	return func(t T, u U) {
		if s.validator == nil {
			return
		}

		validations := s.validator(t, u)
		if validations.HasErrors() {
			vc.AddValidationErrors(validations)
		}
	}
}

func NewWithTwoVariables[T supportedTypes, U supportedTypes](sentences []string,
	definition func(*scenario.Context) func(context.Context, T, U) (context.Context, error),
	validator func(T, U) ValidationErrors,
	documentation DocParams,
) Step {
	return stepTwoVars[T, U]{
		sentences,
		definition,
		validator,
		documentation,
	}
}
