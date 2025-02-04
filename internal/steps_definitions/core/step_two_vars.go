package core

type stepTwoVars[T stepSupportedTypes, U stepSupportedTypes] struct {
	sentences  []string
	definition func(*TestSuiteContext) func(T, U) error
	validator  func(T, U) ValidationErrors
}

func (s stepTwoVars[T, U]) GetSentences() []string {
	return s.sentences
}

func (s stepTwoVars[T, U]) GetDefinition(ctx *TestSuiteContext) any {
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
	definition func(*TestSuiteContext) func(T, U) error,
	validator func(T, U) ValidationErrors) TestStep {
	return stepTwoVars[T, U]{
		sentences,
		definition,
		validator,
	}
}
