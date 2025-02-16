package gherkinparser

import (
	"etoolse/pkg/logger"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	gherkin "github.com/cucumber/gherkin/go/v26"
	"github.com/gofrs/uuid/v5"
)

func Parse(featureFilesLocation string) []*Feature {
	features := getFeaturesAndMacros(featureFilesLocation)
	macroFeatures, testFeatures := separateMacroAndTestsFeatures(features)

	macros := getMacros(macroFeatures)
	featuresContainingMacros, featuresWithoutMacros := separateFeaturesContainingMacrosOrNot(macros, testFeatures)

	applyMacros(macros, featuresContainingMacros)

	return slices.Concat(featuresWithoutMacros, featuresContainingMacros)
}

func getFeaturesAndMacros(featureFilesLocation string) []*Feature {
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
	for _, child := range gherkinDoc.Feature.Children {
		if child.Scenario == nil {
			continue
		}
		scenarios = append(scenarios, child.Scenario)
	}

	return newFeature(
		gherkinDoc.Feature.Name,
		featurePath,
		fileContent,
		scenarios,
	)
}

func separateMacroAndTestsFeatures(features []*Feature) ([]*Feature, []*Feature) {
	var macroFeatures, testFeatures []*Feature
	for _, feature := range features {
		isMacroFeature := strings.HasSuffix(feature.uri, ".macro.feature")
		if isMacroFeature {
			macroFeatures = append(macroFeatures, feature)
		} else {
			testFeatures = append(testFeatures, feature)
		}
	}
	return macroFeatures, testFeatures
}

func separateFeaturesContainingMacrosOrNot(macros []*scenario, testFeatures []*Feature) ([]*Feature, []*Feature) {
	if len(macros) == 0 {
		return []*Feature{}, testFeatures
	}

	var featuresContainingMacros, featuresWithoutMacros []*Feature
	macrostitles := getMacroTitles(macros)
	r := regexp.MustCompile(strings.Join(macrostitles, "|"))

	for _, f := range testFeatures {
		if r.Match(f.Contents) {
			featuresContainingMacros = append(featuresContainingMacros, f)
		} else {
			featuresWithoutMacros = append(featuresWithoutMacros, f)
		}
	}
	return featuresContainingMacros, featuresWithoutMacros
}
