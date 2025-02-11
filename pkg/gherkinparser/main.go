package gherkinparser

import (
	"log"
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

	log.Println("Features with macro: ", len(featuresContainingMacros))
	log.Println("Features without macro: ", len(featuresWithoutMacros))

	applyMacros(macros, featuresContainingMacros)

	return slices.Concat(featuresWithoutMacros, featuresContainingMacros)
}

func getFeaturesAndMacros(featureFilesLocation string) []*Feature {
	featuresPaths, getFeaturesErr := getFeaturesPaths(featureFilesLocation)
	if getFeaturesErr != nil {
		log.Fatal(getFeaturesErr)
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
			log.Fatal(err)
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
		log.Println("Error reading fileContent: ", featurePath)
	}

	gherkinDoc, gherkinParseErr := gherkin.ParseGherkinDocument(strings.NewReader(string(fileContent)), func() string {
		return uuid.Must(uuid.NewV4()).String()
	})

	if gherkinParseErr != nil {
		log.Println("Error parsing gherkin document: ", gherkinParseErr)
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
