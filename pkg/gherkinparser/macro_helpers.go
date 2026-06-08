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

	// Always compute BackgroundStepCount so the scenario tracker can offset correctly.
	bgStepCount := 0
	if feature.background != nil {
		bgStepCount = mh.expandedStepCount(feature.background.Steps)
	}

	if len(mh.macroNameSet) == 0 || !mh.macroPattern.MatchString(currContent) {
		feature.BackgroundStepCount = bgStepCount
		return &feature, nil
	}

	// Accumulate macro entries across reparsing passes (feature = *f resets zero fields).
	var bgMacros []MacroExpansionEntry
	scMacros := make(map[string][]MacroExpansionEntry)

	if feature.background != nil && mh.containsMacro(feature.background.Steps) {
		res, err := mh.applyMacrosAndReparse(feature.background.Steps, currContent)
		if err != nil {
			return nil, err
		}
		bgMacros = res.entries
		bgStepCount = res.totalStepCount
		feature = *res.feature
	}

	for i := range len(feature.scenarios) {
		sc := feature.scenarios[i]
		if sc == nil || isMacroScenario(sc) || !mh.containsMacro(sc.Steps) {
			continue
		}

		res, err := mh.applyMacrosAndReparse(sc.Steps, string(feature.Contents))
		if errors.Is(err, errFeatureParse) {
			continue
		}
		if err != nil {
			return nil, err
		}
		scMacros[sc.Name] = res.entries
		feature = *res.feature
	}

	feature.BackgroundMacros = bgMacros
	feature.BackgroundStepCount = bgStepCount
	feature.ScenarioMacros = scMacros
	return &feature, nil
}

func (mh *Macrohelpers) applyMacrosAndReparse(
	stepsToExpand []*messages.Step, featureContent string,
) (macroExpansionResult, error) {
	res, err := mh.applyMacro(stepsToExpand, featureContent)
	if err != nil {
		return macroExpansionResult{}, err
	}
	f, parseErr := parseFeatureContent(res.content)
	res.feature = f
	return res, parseErr
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

// applyMacro expands all macro calls in scenarioSteps within featureContent.
// Returns a macroExpansionResult with the updated content, expansion entries, and
// total expanded step count. The feature field is not populated here; use
// applyMacrosAndReparse when the re-parsed Feature is also needed.
func (mh *Macrohelpers) applyMacro(
	scenarioSteps []*messages.Step, featureContent string,
) (macroExpansionResult, error) {
	var entries []MacroExpansionEntry
	runningIdx := 0
	featureContentLines := strings.Split(featureContent, "\n")

	// Precompute the absolute expanded step index of each macro call. This
	// count includes normal steps and expanded macro steps so the tracker can
	// collapse the correct step sequence.
	stepOffsets := make(map[*messages.Step]int, len(scenarioSteps))
	for _, step := range scenarioSteps {
		if step == nil {
			continue
		}
		macroIdx := slices.IndexFunc(mh.macros, func(macro scenario) bool {
			return macro.Name == step.Text
		})
		if macroIdx == -1 {
			runningIdx++
			continue
		}
		stepOffsets[step] = runningIdx
		expandedCount := len(mh.macros[macroIdx].Steps)
		runningIdx += expandedCount
	}

	// Collect only the steps that are macro calls, preserving their macro index.
	type macroMatch struct {
		step     *messages.Step
		macroIdx int
	}
	var macroMatches []macroMatch
	for _, step := range scenarioSteps {
		if step == nil {
			continue
		}
		macroIdx := slices.IndexFunc(mh.macros, func(macro scenario) bool {
			return macro.Name == step.Text
		})
		if macroIdx != -1 {
			macroMatches = append(macroMatches, macroMatch{step: step, macroIdx: macroIdx})
		}
	}

	// Process macro calls bottom-to-top so that replacing a lower call does not
	// shift the line numbers of calls that appear earlier in the file.
	slices.SortFunc(macroMatches, func(a, b macroMatch) int {
		return int(b.step.Location.Line) - int(a.step.Location.Line)
	})

	for _, match := range macroMatches {
		expandedCount := len(mh.macros[match.macroIdx].Steps)
		callText := strings.TrimSpace(match.step.Keyword) + " " + match.step.Text
		entries = append(entries, MacroExpansionEntry{
			CallText:  callText,
			StartIdx:  stepOffsets[match.step],
			StepCount: expandedCount,
		})

		// Convert DataTable to map for efficient variable substitution
		// Each row: [variable_name, value] → map entry
		expandedSteps, err := mh.expandMacroSteps(ExpandMacroParam{
			Macro:     mh.macros[match.macroIdx],
			Keyword:   match.step.Keyword,
			Variables: getMacroVariables(match.step.DataTable),
		})
		if err != nil {
			return macroExpansionResult{}, err
		}

		stepStartLine := int(match.step.Location.Line) - 1
		stepEndLine := mh.getStepEndLine(stepStartLine, featureContentLines)

		featureContentLines = slices.Delete(featureContentLines, stepStartLine, stepEndLine+1)
		featureContentLines = slices.Insert(featureContentLines, stepStartLine, expandedSteps...)
	}

	return macroExpansionResult{
		content:        strings.Join(featureContentLines, "\n"),
		entries:        entries,
		totalStepCount: runningIdx,
	}, nil
}

// expandedStepCount returns the total number of godog step-hook calls that will be
// produced when all steps in the slice are executed (macro calls count as their
// expanded step count, regular steps count as one each).
func (mh *Macrohelpers) expandedStepCount(steps []*messages.Step) int {
	count := 0
	for _, step := range steps {
		if step == nil {
			continue
		}
		macroIdx := slices.IndexFunc(mh.macros, func(m scenario) bool { return m.Name == step.Text })
		if macroIdx == -1 {
			count++
		} else {
			count += len(mh.macros[macroIdx].Steps)
		}
	}
	return count
}

func (mh *Macrohelpers) getStepEndLine(stepStartLine int, featureContent []string) int {
	stepEndLine := stepStartLine
	stepKeywords := []string{
		gherkinKeywordGiven,
		gherkinKeywordWhen,
		gherkinKeywordThen,
		gherkinKeywordAnd,
		gherkinKeywordBut,
	}
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
			keyword = gherkinKeywordAnd
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
		delimiter = DocStringDelimiter
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
