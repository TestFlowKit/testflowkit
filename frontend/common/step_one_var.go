package common

type stepOneVar[T stepSupportedTypes] struct {
	sentences  []string
	definition func(*TestSuiteContext) func(T) error
	validator  func(T) ValidationErrors
}

func (s stepOneVar[T]) GetSentences() []string {
	return s.sentences
}

func (s stepOneVar[T]) GetDefinition(ctx *TestSuiteContext) any {
	return s.definition(ctx)
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

func NewStepWithOneVariable[T stepSupportedTypes](sentences []string,
	definition func(*TestSuiteContext) func(T) error,
	validator func(T) ValidationErrors) TestStep {
	return stepOneVar[T]{
		sentences,
		definition,
		validator,
	}
}
