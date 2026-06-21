package restapi

import (
	"context"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/gherkinparser"
	"testflowkit/pkg/logger"
)

func (steps) setJSONBody() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`I set the JSON request body to:`},
		func(ctx context.Context, jsonBody string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			backend.SetRESTRequestBody([]byte(jsonBody))
			logger.Infof("JSON request body set and validated (%d bytes)", len(jsonBody))
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets a JSON request body for the REST API request with JSON validation. The body must be valid JSON.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "json",
					Description: "JSON body content",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: gherkinparser.DocStringExample(
				"Given I set the JSON request body to:",
				"{\n  \"name\": \"John Doe\",\n  \"email\": \"john@example.com\",\n  \"age\": 30,\n  "+
					"\"tags\": [\"developer\", \"golang\"]\n}",
			),
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI},
		},
	)
}
