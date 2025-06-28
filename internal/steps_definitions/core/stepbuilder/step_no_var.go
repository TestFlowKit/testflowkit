package stepbuilder

import (
	"context"
)

type stepWithoutVar struct {
	sentences  []string
	definition func(context.Context) (context.Context, error)
	validator  func() ValidationErrors
	doc        DocParams
}

func (s stepWithoutVar) GetSentences() []string {
	return s.sentences
}

func (s stepWithoutVar) GetDefinition() any {
	return s.definition
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

type noVarDef func(context.Context) (context.Context, error)
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
