package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
	"testflowkit/scripts/shared"
)

const defaultOutputFile = "step-definitions.json"

func main() {
	const usage = "Path to the output JSON file (e.g. dir/step-definitions.json). Defaults to ./step-definitions.json"
	outputFile := flag.String("output-file", "", usage)
	flag.Parse()

	if *outputFile == "" {
		wd, err := os.Getwd()
		if err != nil {
			logger.Fatal("Error getting working directory", err)
		}
		*outputFile = filepath.Join(wd, defaultOutputFile)
	}

	allDocs := shared.GetAllDocs()
	entries := formatDocs(allDocs)

	jsonData, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		logger.Fatal("Error marshalling step definitions", err)
	}

	if writeErr := shared.WriteFileWithDirectories(*outputFile, jsonData); writeErr != nil {
		logger.Fatal("Error writing output file "+*outputFile, writeErr)
	}

	logger.Success("Step definitions exported to " + *outputFile)
}

type stepDefinition struct {
	Sentence    string         `json:"sentence"`
	Description string         `json:"description"`
	Categories  []string       `json:"categories"`
	Example     string         `json:"example"`
	Variables   []stepVariable `json:"variables,omitempty"`
}

type stepVariable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func formatDocs(docs []stepbuilder.Documentation) []stepDefinition {
	entries := make([]stepDefinition, 0, len(docs))

	for _, d := range docs {
		categories := make([]string, 0, len(d.Categories))
		for _, c := range d.Categories {
			categories = append(categories, string(c))
		}

		entry := stepDefinition{
			Sentence:    shared.CleanSentence(d.Sentence),
			Description: d.Description,
			Categories:  categories,
			Example:     d.Example,
		}

		for _, v := range d.Variables {
			entry.Variables = append(entry.Variables, stepVariable{
				Name:        v.Name,
				Description: v.Description,
				Type:        string(v.Type),
			})
		}

		entries = append(entries, entry)
	}

	return entries
}
