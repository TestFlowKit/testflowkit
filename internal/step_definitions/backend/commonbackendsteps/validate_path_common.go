package commonbackendsteps

import (
	"context"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/queryable"
)

func responsePathDocVariable(description string) stepbuilder.DocVariable {
	return stepbuilder.DocVariable{
		Name:        "path",
		Description: description,
		Type:        stepbuilder.VarTypeString,
	}
}

func newResponseFieldDocParams(
	description string,
	pathDescription string,
	valueName string,
	valueDescription string,
	example string,
) stepbuilder.DocParams {
	return stepbuilder.DocParams{
		Description: description,
		Variables: []stepbuilder.DocVariable{
			responsePathDocVariable(pathDescription),
			{Name: valueName, Description: valueDescription, Type: stepbuilder.VarTypeString},
		},
		Example:    example,
		Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
	}
}

func (s steps) newJSONPathStringStep(
	sentence string,
	shouldMatch bool,
	doc stepbuilder.DocParams,
	validator func(queryable.Queryable, string, string, bool) error,
) stepbuilder.Step {
	return stepbuilder.NewWithTwoVariables(
		[]string{sentence},
		func(ctx context.Context, jsonPath, expected string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !backend.HasResponse() {
				return ctx, apperrors.ErrNoResponseAvailable
			}

			jsonPath = scenario.ReplaceVariablesInString(scenarioCtx, jsonPath)
			expected = scenario.ReplaceVariablesInString(scenarioCtx, expected)

			engine, err := newResponseEngine(backend)
			if err != nil {
				return ctx, err
			}

			return ctx, validator(engine, jsonPath, expected, shouldMatch)
		},
		nil,
		doc,
	)
}
