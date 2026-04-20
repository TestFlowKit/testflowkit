package main

import (
	"encoding/json"
	"os"
	"path"
	"regexp"
	"strings"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
	"testflowkit/scripts/shared"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logger.Fatal("Error getting working directory", err)
	}
	outputDir := path.Join(wd, "documentation", "content", "sentences")

	allDocs := shared.GetAllDocs()
	generateDocs(allDocs, outputDir)
}

func generateDocs(stepDocumentations []stepbuilder.Documentation, outputDir string) {
	err := os.RemoveAll(outputDir)
	if err != nil {
		logger.Fatal("Error removing directory "+outputDir, err)
	}

	docs := formatSentencesDocs(stepDocumentations)
	for _, documentation := range docs {
		jsonData, jsonEerr := json.Marshal(documentation)
		if jsonEerr != nil {
			logger.Fatal(documentation.Sentence+" documentation generation failed", err)
		}

		filename := formatFilename(documentation.Sentence)
		categoryDir := "uncategorized"
		if len(documentation.Categories) > 0 {
			categoryDir = documentation.Categories[0]
		}

		filePath := path.Join(outputDir, categoryDir, filename)

		fileCreationErr := createFileWithDirectories(filePath, jsonData)
		if fileCreationErr != nil {
			logger.Fatal("Error creating file "+filePath, err)
		}
	}
}

func formatSentencesDocs(sentences []stepbuilder.Documentation) (docs []doc) {
	for _, step := range sentences {
		var categories []string
		for _, c := range step.Categories {
			categories = append(categories, string(c))
		}

		curr := doc{
			Sentence:    shared.CleanSentence(step.Sentence),
			Description: step.Description,
			Categories:  categories,
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
	return strings.ToLower(sentence) + ".json"
}

func createFileWithDirectories(filePath string, data []byte) error {
	return shared.WriteFileWithDirectories(filePath, data)
}

type doc struct {
	Sentence    string   `json:"sentence"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Variables   []docVar `json:"variables"`
	Example     string   `json:"gherkinExample"`
}

type docVar struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
