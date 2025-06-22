package stepbuilder

import (
	"testflowkit/internal/steps_definitions/core/scenario"
)

type stepWithoutVar struct {
	sentences  []string
	definition func(*scenario.Context) func() error
	validator  func() ValidationErrors
	doc        DocParams
}

func (s stepWithoutVar) GetSentences() []string {
	return s.sentences
}

func (s stepWithoutVar) GetDefinition(ctx *scenario.Context) any {
	return s.definition(ctx)
}

func (s stepWithoutVar) GetDocumentation() Documentation {
	return Documentation{
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

type noVarDef func(*scenario.Context) func() error
type noVarValidator func() ValidationErrors

func NewWithNoVariables(
	sentences []string,
	definition noVarDef,
	validator noVarValidator,
	documentation DocParams,
) Step {
	return stepWithoutVar{
		sentences,
		definition,
		validator,
		documentation,
	}
}
