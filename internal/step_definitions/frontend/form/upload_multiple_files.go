package form

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/utils/stringutils"
)

func (steps) userUploadsMultipleFiles() stepbuilder.Step {
	filenamesDesc := "Comma-separated list of logical file names as defined in configuration."

	return stepbuilder.NewWithTwoVariables(
		[]string{`the user uploads the {string} files into the {string} field`},
		userUploadsMultipleFilesHandler,
		func(fileNames, inputLabel string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			label := userUploadsMultipleFilesFormatLabel(inputLabel)

			if !config.IsElementDefined(label) {
				vc.AddMissingElement(label)
			}

			fileNameList := strings.Split(fileNames, ",")
			for _, fileName := range fileNameList {
				fileName = strings.TrimSpace(fileName)
				if !config.IsFileDefined(fileName) {
					vc.AddMissingFile(fileName)
				}
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "Uploads multiple files into a file input field identified by its logical name.",
			Variables: []stepbuilder.DocVariable{
				{Name: "fileNames", Description: filenamesDesc, Type: stepbuilder.VarTypeString},
				{Name: "inputLabel", Description: "The logical name of the file input field.", Type: stepbuilder.VarTypeString},
			},
			Example:    `When the user uploads the "image1, image2, image3" files into the "Gallery" field`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Form}},
	)
}

func userUploadsMultipleFilesHandler(ctx context.Context, fileNames, inputLabel string) (context.Context, error) {
	scenarioCtx := scenario.MustFromContext(ctx)
	cfg := scenarioCtx.GetConfig()

	fileNameList := strings.Split(fileNames, ",")
	for i, name := range fileNameList {
		fileNameList[i] = strings.TrimSpace(name)
	}

	filePaths, err := cfg.GetFilesPaths(fileNameList)
	if err != nil {
		return ctx, fmt.Errorf("failed to get file paths for '%s': %w", fileNames, err)
	}

	input, err := scenarioCtx.GetHTMLElementByLabel(userUploadsMultipleFilesFormatLabel(inputLabel))
	if err != nil {
		return ctx, err
	}

	err = input.UploadMultipleFiles(filePaths)
	if err != nil {
		return ctx, fmt.Errorf("failed to upload files '%s' to field '%s': %w", fileNames, inputLabel, err)
	}

	return ctx, nil
}

func userUploadsMultipleFilesFormatLabel(label string) string {
	return stringutils.SuffixWithUnderscore(label, "field")
}
