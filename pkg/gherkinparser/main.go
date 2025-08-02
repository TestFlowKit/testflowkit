package gherkinparser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testflowkit/pkg/logger"

	gherkin "github.com/cucumber/gherkin/go/v26"
	messages "github.com/cucumber/messages/go/v21"
	"github.com/gofrs/uuid/v5"
)

func Parse(featureFilesLocation string) []*Feature {
	features := getFeatures(featureFilesLocation)
	macros := getMacros(features)

	applyMacros(macros, features)

	return features
}

func getFeatures(featureFilesLocation string) []*Feature {
	featuresPaths, getFeaturesErr := getFeaturesPaths(featureFilesLocation)
	if getFeaturesErr != nil {
		logger.Fatal("Error getting features paths", getFeaturesErr)
	}

	var features []*Feature
	for _, featurePath := range featuresPaths {
		feature := parseFeatureFile(featurePath)
		if feature == nil {
			continue
		}

		features = append(features, feature)
	}

	return features
}

func getFeaturesPaths(featureFilesLocation string) ([]string, error) {
	var featuresPaths []string
	getFeaturesErr := filepath.Walk(featureFilesLocation, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Fatal("Error getting features paths", err)
		}
		if info.IsDir() {
			return nil
		}

		featuresPaths = append(featuresPaths, path)
		return nil
	})
	return featuresPaths, getFeaturesErr
}

func parseFeatureFile(featurePath string) *Feature {
	fileContent, readFileErr := os.ReadFile(featurePath)
	if readFileErr != nil {
		msg := fmt.Sprintf("Error reading fileContent: %s", featurePath)
		logger.Warn(msg, []string{"Please check the file read permissions"})
	}
	gherkinDoc, gherkinParseErr := gherkin.ParseGherkinDocument(strings.NewReader(string(fileContent)), func() string {
		return uuid.Must(uuid.NewV4()).String()
	})

	if gherkinParseErr != nil {
		logger.Warn(fmt.Sprintf("Error parsing feature file: %s", featurePath), []string{"Please check the file syntax"})
		return nil
	}

	var scenarios []*scenario
	var background *messages.Background
	for _, child := range gherkinDoc.Feature.Children {
		if child.Scenario != nil {
			scenarios = append(scenarios, child.Scenario)
		}
		if child.Background != nil {
			background = child.Background
		}
	}

	return newFeature(
		gherkinDoc.Feature.Name,
		featurePath,
		fileContent,
		scenarios,
		background,
	)
}
