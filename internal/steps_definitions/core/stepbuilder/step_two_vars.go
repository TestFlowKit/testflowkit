package stepbuilder

import (
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/shared"
)

type stepTwoVars[T stepSupportedTypes, U stepSupportedTypes] struct {
	sentences  []string
	definition func(*scenario.Context) func(T, U) error
	validator  func(T, U) ValidationErrors
	doc        StepDefDocParams
}

func (s stepTwoVars[T, U]) GetDocumentation() shared.StepDocumentation {
	return shared.StepDocumentation{
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

func NewStepWithTwoVariables[T stepSupportedTypes, U stepSupportedTypes](sentences []string,
	definition func(*scenario.Context) func(T, U) error,
	validator func(T, U) ValidationErrors,
	documentation StepDefDocParams,
) TestStep {
	return stepTwoVars[T, U]{
		sentences,
		definition,
		validator,
		documentation,
	}
}
