package gherkinparser

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"

	messages "github.com/cucumber/messages/go/v21"
)

type Macrohelpers struct {
	macros       []scenario
	macroNameSet map[string]struct{}
	macroPattern *regexp.Regexp
}

func NewMacroHelpers(macros []scenario) *Macrohelpers {
	macroNameSet := make(map[string]struct{}, len(macros))
	escapedTitles := make([]string, 0, len(macros))
	for _, macro := range macros {
		macroNameSet[macro.Name] = struct{}{}
		escapedTitles = append(escapedTitles, regexp.QuoteMeta(macro.Name))
	}

	// Keep a compiled pattern for compatibility/perf checks, but make it safe for
	// macro names containing regex metacharacters.
	pattern := regexp.MustCompile("^$")
	if len(escapedTitles) > 0 {
		pattern = regexp.MustCompile(strings.Join(escapedTitles, "|"))
	}

	return &Macrohelpers{
		macros:       macros,
		macroNameSet: macroNameSet,
		macroPattern: pattern,
	}
}

func (mh *Macrohelpers) ApplyMacroToFeature(feature Feature) (*Feature, error) {
	currContent := string(feature.Contents)

	if len(mh.macroNameSet) == 0 || !mh.macroPattern.MatchString(currContent) {
		return &feature, nil
	}

	if feature.background != nil && mh.containsMacro(feature.background.Steps) {
		f, err := mh.applyMacrosAndReparse(feature.background.Steps, currContent)
		if err != nil {
			return nil, err
		}
		feature = *f
	}

	for i := range len(feature.scenarios) {
		sc := feature.scenarios[i]
		if sc == nil || isMacroScenario(sc) || !mh.containsMacro(sc.Steps) {
			continue
		}

		f, err := mh.applyMacrosAndReparse(sc.Steps, string(feature.Contents))
		if errors.Is(err, errFeatureParse) {
			continue
		}

		if err != nil {
			return nil, err
		}
		feature = *f
	}

	return &feature, nil
}

func (mh *Macrohelpers) applyMacrosAndReparse(stepsToExpand []*messages.Step, featureContent string) (*Feature, error) {
	newFeatureContent, err := mh.applyMacro(stepsToExpand, featureContent)
	if err != nil {
		return nil, err
	}
	return parseFeatureContent(newFeatureContent)
}

func (mh *Macrohelpers) containsMacro(steps []*messages.Step) bool {
	for _, step := range steps {
		if step == nil {
			continue
		}
		if _, ok := mh.macroNameSet[step.Text]; ok {
			return true
		}
	}
	return false
}

// apply macro and return updated feature content.
func (mh *Macrohelpers) applyMacro(scenarioSteps []*messages.Step, featureContent string) (string, error) {
	featureContentLines := strings.Split(featureContent, "\n")
	for _, step := range scenarioSteps {
		macroIdx := slices.IndexFunc(mh.macros, func(macro scenario) bool {
			return macro.Name == step.Text
		})

		isMacroStep := macroIdx != -1
		if !isMacroStep {
			continue
		}

		// Convert DataTable to map for efficient variable substitution
		// Each row: [variable_name, value] → map entry
		expandedSteps, err := mh.expandMacroSteps(ExpandMacroParam{
			Macro:     mh.macros[macroIdx],
			Keyword:   step.Keyword,
			Variables: getMacroVariables(step.DataTable), // Convert to map here
		})
		if err != nil {
			return "", err
		}

		stepStartLine := int(step.Location.Line) - 1
		stepEndLine := mh.getStepEndLine(stepStartLine, featureContentLines)

		featureContentLines = slices.Delete(featureContentLines, stepStartLine, stepEndLine+1)

		featureContentLines = slices.Insert(featureContentLines, stepStartLine, expandedSteps...)
	}

	return strings.Join(featureContentLines, "\n"), nil
}

func (mh *Macrohelpers) getStepEndLine(stepStartLine int, featureContent []string) int {
	stepEndLine := stepStartLine
	stepKeywords := []string{"Given", "When", "Then", "And", "But"}
	structureKeywords := []string{"Scenario:", "Scenario Outline:", "Background:", "Feature:", "Examples:"}

	for i := stepStartLine + 1; i < len(featureContent); i++ {
		line := strings.TrimSpace((featureContent)[i])

		// Check if this is a new step
		isNewStep := slices.ContainsFunc(stepKeywords, func(keyword string) bool {
			return strings.HasPrefix(line, keyword)
		})

		// Check if this is a new Gherkin structure
		isNewStructure := slices.ContainsFunc(structureKeywords, func(keyword string) bool {
			return strings.HasPrefix(line, keyword)
		})

		if isNewStep || isNewStructure {
			break
		}

		// If this is an empty line, check if the next non-empty line is a structure keyword
		if line == "" && i+1 < len(featureContent) {
			nextLine := strings.TrimSpace(featureContent[i+1])
			isNextLineStructure := slices.ContainsFunc(structureKeywords, func(keyword string) bool {
				return strings.HasPrefix(nextLine, keyword)
			})
			if isNextLineStructure {
				break
			}
		}

		stepEndLine = i
	}

	return stepEndLine
}

// expandMacroSteps expands a macro scenario into concrete steps with variable substitution.
// Uses the map-based Variables from ExpandMacroParam for efficient ${variable} replacement.
// Handles step text, docstrings, and nested data tables within macro definitions.
func (mh *Macrohelpers) expandMacroSteps(params ExpandMacroParam) ([]string, error) {
	var expandedSteps []string
	for idx, macroStep := range params.Macro.Steps {
		keyword := strings.TrimSpace(params.Keyword)
		if idx > 0 {
			keyword = "And"
		}

		// Substitute variables using the map for O(1) lookup
		stepText, err := substituteVariables(macroStep.Text, params.Variables)
		if err != nil {
			return nil, err
		}
		expandedSteps = append(expandedSteps, fmt.Sprintf("%s %s", keyword, stepText))

		newLines, errDoc := mh.applyToDocString(macroStep.DocString, params.Variables)
		if errDoc != nil {
			return nil, errDoc
		}
		expandedSteps = append(expandedSteps, newLines...)

		if macroStep.DataTable != nil {
			datatableContent, errDataString := convertDatatableToString(macroStep.DataTable, params.Variables)
			if errDataString != nil {
				return nil, errDataString
			}
			expandedSteps = append(expandedSteps, datatableContent)
		}
	}
	return expandedSteps, nil
}

func (mh *Macrohelpers) applyToDocString(ds *messages.DocString, vars MacroVariables) ([]string, error) {
	var newParts []string
	if ds == nil {
		return newParts, nil
	}

	delimiter := ds.Delimiter
	if delimiter == "" {
		delimiter = "\"\"\""
	}
	newParts = append(newParts, delimiter)
	docStringContent, err := substituteVariables(ds.Content, vars)
	if err != nil {
		return nil, err
	}
	newParts = append(newParts, docStringContent)
	newParts = append(newParts, delimiter)

	return newParts, nil
}
