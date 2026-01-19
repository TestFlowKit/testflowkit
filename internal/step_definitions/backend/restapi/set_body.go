package restapi

import (
	"context"
	"fmt"
	"os"
	"strings"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/fileutils"
	"testflowkit/pkg/logger"
)

func (steps) setRequestBody() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{
			`I set the request body to:`,
		},
		func(ctx context.Context, body string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			// Check if it looks like a file path (second sentence pattern)
			// File paths would be passed directly, multiline strings would have newlines

			backend.SetRESTRequestBody([]byte(body))
			logger.InfoFf("Request body set (%d bytes)", len(body))
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets the raw request body for the REST API request.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "body",
					Description: "Raw body content or file path",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example: `Given I set the request body to:
"""
{"name": "John", "email": "john@example.com"}
"""
`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI}},
	)
}

func (steps) setRequestBodyFromFile() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{
			`I set the request body from file {string}`,
		},
		func(ctx context.Context, filePath string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			backend := scenarioCtx.GetBackendContext()

			if !strings.HasSuffix(filePath, ".json") {
				return ctx, fmt.Errorf("only JSON files are supported for request body: %s", filePath)
			}

			errPath := fileutils.ValidatePath(filePath)
			if errPath != nil {
				return ctx, errPath
			}

			content, err := os.ReadFile(filePath)
			if err != nil {
				return ctx, fmt.Errorf("failed to read request body file '%s': %w", filePath, err)
			}

			if !jsonhelpers.IsValid(content) {
				return ctx, fmt.Errorf("invalid JSON in request body file '%s'", filePath)
			}

			backend.SetRESTRequestBody(content)
			logger.InfoFf("Request body set from file: %s (%d bytes)", filePath, len(content))
			return ctx, nil
		},
		nil,
		stepbuilder.DocParams{
			Description: "Sets the request body for the REST API request from a file.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "filePath",
					Description: "Path to the file containing the request body",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:    `Given I set the request body from file "data/request.json"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.RESTAPI}},
	)
}
