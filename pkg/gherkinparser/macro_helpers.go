package gherkinparser

import (
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
		currContent = mh.applyMacro(feature.background.Steps, currContent)
		fUpdated, err := parseFeatureContent(currContent)
		if err != nil {
			return nil, err
		}
		feature = *fUpdated
	}

	for _, sc := range feature.scenarios {
		if sc == nil || isMacroScenario(sc) || !mh.containsMacro(sc.Steps) {
			continue
		}

		newFeatureContent := mh.applyMacro(sc.Steps, string(feature.Contents))
		fUpdated, err := parseFeatureContent(newFeatureContent)
		if err != nil {
			continue
		}
		feature = *fUpdated
	}

	return &feature, nil
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
func (mh *Macrohelpers) applyMacro(scenarioSteps []*messages.Step, featureContent string) string {
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
		expandedSteps := mh.expandMacroSteps(ExpandMacroParam{
			Macro:     mh.macros[macroIdx],
			Keyword:   step.Keyword,
			Variables: getMacroVariables(step.DataTable), // Convert to map here
		})

		stepStartLine := int(step.Location.Line) - 1
		stepEndLine := mh.getStepEndLine(stepStartLine, featureContentLines)

		featureContentLines = slices.Delete(featureContentLines, stepStartLine, stepEndLine+1)

		featureContentLines = slices.Insert(featureContentLines, stepStartLine, expandedSteps...)
	}

	return strings.Join(featureContentLines, "\n")
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
func (mh *Macrohelpers) expandMacroSteps(params ExpandMacroParam) []string {
	var expandedSteps []string
	for idx, macroStep := range params.Macro.Steps {
		keyword := params.Keyword
		if idx > 0 {
			keyword = "And"
		}

		// Substitute variables using the map for O(1) lookup
		stepText := substituteVariables(macroStep.Text, params.Variables)
		expandedSteps = append(expandedSteps, fmt.Sprintf("%s %s", keyword, stepText))

		if macroStep.DocString != nil {
			ds := macroStep.DocString
			delimiter := ds.Delimiter
			if delimiter == "" {
				delimiter = "\"\"\""
			}
			expandedSteps = append(expandedSteps, delimiter)
			// Also substitute variables in docstring content
			expandedSteps = append(expandedSteps, substituteVariables(ds.Content, params.Variables))
			expandedSteps = append(expandedSteps, delimiter)
		}

		if macroStep.DataTable != nil {
			expandedSteps = append(expandedSteps, convertDatatableToString(macroStep.DataTable))
		}
	}
	return expandedSteps
}
