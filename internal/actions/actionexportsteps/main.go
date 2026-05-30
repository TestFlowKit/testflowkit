package actionexportsteps

import (
	"encoding/json"
	"fmt"
	"os"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
	"testflowkit/scripts/shared"
)

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

func Execute(format string) {
	if format != "json" {
		fmt.Fprintf(os.Stderr, "error: unsupported format %q — only \"json\" is supported\n", format)
		os.Exit(1)
	}

	allDocs := shared.GetAllDocs()
	entries := formatDocs(allDocs)

	jsonData, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		logger.Fatal("Error marshalling step definitions", err)
	}

	if _, writeErr := os.Stdout.Write(jsonData); writeErr != nil {
		logger.Fatal("Error writing step definitions to stdout", writeErr)
	}
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
