package main

import (
	"encoding/json"
	"etoolse/internal/steps_definitions/frontend"
	"etoolse/pkg/logger"
	"etoolse/shared"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logger.Fatal("Error getting working directory", err)
	}
	outputDir := path.Join(wd, "documentation", "content", "sentences")

	generateDocs(frontend.GetDocs(), outputDir)
}

func generateDocs(stepDocumentations []shared.StepDocumentation, outputDir string) {
	err := os.RemoveAll(outputDir)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error removing directory %s", outputDir), err)
	}

	docs := formatSentencesDocs(stepDocumentations)
	for _, documentation := range docs {
		jsonData, jsonEerr := json.Marshal(documentation)
		if jsonEerr != nil {
			logger.Fatal(fmt.Sprintf("%s documentation generation failed", documentation.Sentence), err)
		}

		filename := formatFilename(documentation.Sentence)
		filePath := path.Join(outputDir, documentation.Category, filename)

		fileCreationErr := createFileWithDirectories(filePath, jsonData)
		if fileCreationErr != nil {
			logger.Fatal(fmt.Sprintf("Error creating file %s", filePath), err)
		}
	}
}

func formatSentencesDocs(sentences []shared.StepDocumentation) (docs []doc) {
	re := regexp.MustCompile(`[$^]`)

	for _, step := range sentences {
		curr := doc{
			Sentence:    re.ReplaceAllString(step.Sentence, ""),
			Description: step.Description,
			Category:    string(step.Category),
			Example:     step.Example,
		}

		for _, v := range step.Variables {
			curr.Variables = append(curr.Variables, docVar{
				Description: v.Description,
				Name:        v.Name,
				Type:        string(v.Type),
			})
		}

		docs = append(docs, curr)
	}

	return docs
}

func formatFilename(sentence string) string {
	re := regexp.MustCompile(`[!@#$%^&*(){}\[\]\"]`)
	sentence = re.ReplaceAllString(sentence, "")
	sentence = strings.ReplaceAll(sentence, " ", "_")
	sentence = strings.ReplaceAll(sentence, "'", "")
	return fmt.Sprintf("%s.json", strings.ToLower(sentence))
}

func createFileWithDirectories(filePath string, data []byte) error {
	dir := filepath.Dir(filePath)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	mkdErr := os.WriteFile(filePath, data, 0600)
	if mkdErr != nil {
		return err
	}

	return nil
}

type doc struct {
	Sentence    string   `json:"sentence"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Variables   []docVar `json:"variables"`
	Example     string   `json:"gherkinExample"`
}

type docVar struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
