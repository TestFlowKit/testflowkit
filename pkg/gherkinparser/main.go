package gherkinparser

import (
	"errors"
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
	return applyMacros(macros, features)
}

// Filter applies a tag expression to an already-parsed list of features.
// It is the filtering step of ParseWithFilter, usable when features are parsed once
// and filtered multiple times with different expressions.
func Filter(features []*Feature, expr string) []*Feature {
	groups := parseTagExpression(expr)
	if groups == nil {
		return features
	}
	return filterFeatures(features, groups)
}

func ParseWithFilter(featureFilesLocation, expr string) []*Feature {
	return Filter(Parse(featureFilesLocation), expr)
}

// ParseContent parses a single Gherkin feature document from its raw string content.
// It is exported primarily for use in tests that need Feature values without filesystem I/O.
func ParseContent(content string) (*Feature, error) {
	return parseFeatureContent(content)
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

		if filepath.Ext(path) == ".feature" {
			featuresPaths = append(featuresPaths, path)
		}
		return nil
	})
	return featuresPaths, getFeaturesErr
}

func parseFeatureFile(featurePath string) *Feature {
	fileContent, readFileErr := os.ReadFile(featurePath)
	if readFileErr != nil {
		msg := "Error reading fileContent: " + featurePath
		logger.Warn(msg, []string{"Please check the file read permissions"})
	}

	featureParsed, err := parseFeatureContent(string(fileContent))
	if err != nil && errors.Is(err, errFeatureParse) {
		logger.Warn("Error parsing feature file: "+featurePath, []string{"Please check the file syntax"})
		return nil
	}

	return featureParsed
}

func parseFeatureContent(content string) (*Feature, error) {
	gherkinDoc, gherkinParseErr := gherkin.ParseGherkinDocument(strings.NewReader(content), func() string {
		return uuid.Must(uuid.NewV4()).String()
	})

	if gherkinParseErr != nil {
		return nil, errFeatureParse
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

	return newFeature(NewFeatureParams{
		Name:       gherkinDoc.Feature.Name,
		Content:    []byte(content),
		Scenarios:  scenarios,
		Background: background,
		Tags:       gherkinDoc.Feature.Tags,
	}), nil
}

var errFeatureParse = errors.New("error parsing feature content")
