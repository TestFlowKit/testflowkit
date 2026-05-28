package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"testflowkit/pkg/logger"
	"testflowkit/scripts/shared"
)

const defaultOutputFile = "config-schema.json"

func main() {
	outputFile := flag.String("output-file", "", "Path to the output JSON file. Defaults to ./config-schema.json")
	appVersion := flag.String("app-version", "unknown", "CLI version embedded in the schema (e.g. 1.4.2)")
	flag.Parse()

	if *outputFile == "" {
		wd, err := os.Getwd()
		if err != nil {
			logger.Fatal("Error getting working directory", err)
		}
		*outputFile = filepath.Join(wd, defaultOutputFile)
	}

	comments, err := parseConfigComments(filepath.Join("internal", "config"))
	if err != nil {
		logger.Fatal("Error parsing config package", err)
	}

	schema := buildSchema(*appVersion, comments)

	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		logger.Fatal("Error marshalling config schema", err)
	}

	if writeErr := shared.WriteFileWithDirectories(*outputFile, data); writeErr != nil {
		logger.Fatal("Error writing output file "+*outputFile, writeErr)
	}

	logger.Success("Config schema exported to " + *outputFile)
}
