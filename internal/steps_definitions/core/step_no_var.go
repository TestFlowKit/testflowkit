package core

type stepWithoutVar struct {
	sentences  []string
	definition func(*TestSuiteContext) func() error
	validator  func() ValidationErrors
}

func (s stepWithoutVar) GetSentences() []string {
	return s.sentences
}

func (s stepWithoutVar) GetDefinition(ctx *TestSuiteContext) any {
	return s.definition(ctx)
}

func (s stepWithoutVar) Validate(vc *ValidatorContext) any {
	return func() {
		if s.validator == nil {
			return
		}

		validations := s.validator()
		if validations.HasErrors() {
			vc.AddValidationErrors(validations)
		}
	}
}

type noVarDef func(*TestSuiteContext) func() error
type noVarValidator func() ValidationErrors

func NewStepWithoutVariables(sentences []string, definition noVarDef, validator noVarValidator) TestStep {
	return stepWithoutVar{
		sentences,
		definition,
		validator,
	}
}
