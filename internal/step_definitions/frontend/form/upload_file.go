package form

import (
	"context"
	"fmt"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (steps) userUploadsFileIntoField() stepbuilder.Step {
	formatLabel := func(label string) string {
		return stringutils.SuffixWithUnderscore(label, "field")
	}

	filenameDesc := "The logical name of the file as defined in configuration."
	return stepbuilder.NewWithTwoVariables(
		[]string{`the user uploads the {string} file into the {string} field`},
		func(ctx context.Context, fileName, inputLabel string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			cfg := scenarioCtx.GetConfig()

			filePaths, getFPErr := cfg.GetFilesPaths([]string{fileName})
			if getFPErr != nil {
				return ctx, fmt.Errorf("failed to get file path for '%s': %w", fileName, getFPErr)
			}

			input, getInputErr := scenarioCtx.GetHTMLElementByLabel(formatLabel(inputLabel))
			if getInputErr != nil {
				return ctx, fmt.Errorf("failed to get input element for '%s': %w", inputLabel, getInputErr)
			}

			filePath := filePaths[0]
			uploadFileErr := input.UploadFile(filePath)
			if uploadFileErr != nil {
				return ctx, fmt.Errorf("failed to upload file '%s' to field '%s': %w", fileName, inputLabel, uploadFileErr)
			}

			return ctx, nil
		},
		func(fileName, inputLabel string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := formatLabel(inputLabel)
			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}
			cfg, err := config.Get()
			if err != nil {
				vc.AddError("configuration not loaded")
			} else if _, getFPErr := cfg.GetFilesPaths([]string{fileName}); getFPErr != nil {
				vc.AddError(fmt.Sprintf("file '%s' not defined: %v", fileName, getFPErr))
			}
			return vc
		},
		stepbuilder.DocParams{
			Description: "Uploads a single file into a file input field identified by its logical name.",
			Variables: []stepbuilder.DocVariable{
				{Name: "fileName", Description: filenameDesc, Type: stepbuilder.VarTypeString},
				{Name: "inputLabel", Description: "The logical name of the file input field.", Type: stepbuilder.VarTypeString},
			},
			Example:    `When the user uploads the "profile_image" file into the "Avatar" field`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Form},
		},
	)
}
