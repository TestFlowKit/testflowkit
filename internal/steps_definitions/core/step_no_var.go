package core

import "testflowkit/shared"

type stepWithoutVar struct {
	sentences  []string
	definition func(*TestSuiteContext) func() error
	validator  func() ValidationErrors
	doc        StepDefDocParams
}

func (s stepWithoutVar) GetSentences() []string {
	return s.sentences
}

func (s stepWithoutVar) GetDefinition(ctx *TestSuiteContext) any {
	return s.definition(ctx)
}

func (s stepWithoutVar) GetDocumentation() shared.StepDocumentation {
	return shared.StepDocumentation{
		Sentence:    s.sentences[0],
		Description: s.doc.Description,
		Example:     s.doc.Example,
		Category:    s.doc.Category,
		Variables:   s.doc.Variables,
	}
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

func NewStepWithoutVariables(
	sentences []string,
	definition noVarDef,
	validator noVarValidator,
	documentation StepDefDocParams,
) TestStep {
	return stepWithoutVar{
		sentences,
		definition,
		validator,
		documentation,
	}
}
