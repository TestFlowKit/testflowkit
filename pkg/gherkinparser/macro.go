package gherkinparser

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"testflowkit/pkg/logger"
	"unicode/utf8"

	messages "github.com/cucumber/messages/go/v21"
)

const MacroTag = "@macro"

var macroVarPattern = regexp.MustCompile(`\$\{([^}]+)\}`)

type MacroVariable struct {
	Name  string
	Value string
}

type MacroVariables = map[string]string

func getMacroVariables(table *messages.DataTable) MacroVariables {
	if table == nil {
		return make(map[string]string)
	}

	variables := make(map[string]string)
	const headerAndFirstRow = 2
	if len(table.Rows) >= headerAndFirstRow {
		headers := table.Rows[0].Cells
		dataRow := table.Rows[1].Cells

		for i, header := range headers {
			varName := strings.TrimSpace(header.Value)
			varValue := strings.TrimSpace(dataRow[i].Value)

			variables[varName] = varValue
		}
	}

	return variables
}

func substituteVariables(stepContent string, variables MacroVariables) string {
	return macroVarPattern.ReplaceAllStringFunc(stepContent, func(match string) string {
		varName := strings.TrimSpace(match[2 : len(match)-1])
		if value, exists := variables[varName]; exists {
			return value
		}

		logger.Warn("Macro variable not found: "+varName, nil)
		return match
	})
}

// TODO: refactor for better understanding.
func applyMacros(macros []scenario, features []*Feature) []*Feature {
	newFeatures := make([]*Feature, 0, len(features))
	for _, f := range features {
		currFeature, err := applyMacroToFeature(macros, *f)
		if err != nil {
			logger.Warn("Error applying macros to feature: "+f.Name, []string{err.Error()})
		}
		newFeatures = append(newFeatures, currFeature)
	}

	return newFeatures
}

func applyMacroToFeature(macros []scenario, feature Feature) (*Feature, error) {
	macroTitles := getMacroTitles(macros)
	currContent := string(feature.Contents)

	canContainsMacroRegex := regexp.MustCompile(strings.Join(macroTitles, "|"))
	if !canContainsMacroRegex.MatchString(currContent) {
		return &feature, nil
	}

	if feature.background != nil && containsMacro(feature.background.Steps, macroTitles) {
		currContent = applyMacro(feature.background.Steps, macros, currContent)
		fUpdated, err := parseFeatureContent(currContent)
		if err != nil {
			return nil, err
		}
		feature = *fUpdated
	}

	for _, sc := range feature.scenarios {
		if sc == nil || isMacroScenario(sc) || !containsMacro(sc.Steps, macroTitles) {
			continue
		}

		newFeatureContent := applyMacro(sc.Steps, macros, string(feature.Contents))
		fUpdated, err := parseFeatureContent(newFeatureContent)
		if err != nil {
			continue
		}
		feature = *fUpdated
	}

	return &feature, nil
}

func containsMacro(steps []*messages.Step, macroNames []string) bool {
	reg := regexp.MustCompile(strings.Join(macroNames, "|"))
	var stepTextSb strings.Builder
	for _, step := range steps {
		stepTextSb.WriteString(step.Text + "\n")
	}
	return reg.MatchString(stepTextSb.String())
}

// apply macro and return updated feature content.
func applyMacro(scenarioSteps []*messages.Step, macros []scenario, featureContent string) string {
	featureContentLines := strings.Split(featureContent, "\n")
	for _, step := range scenarioSteps {
		macroIdx := slices.IndexFunc(macros, func(macro scenario) bool {
			return macro.Name == step.Text
		})

		isMacroStep := macroIdx != -1
		if !isMacroStep {
			continue
		}

		expandedSteps := expandMacroSteps(ExpandMacroParam{
			Macro:     macros[macroIdx],
			Keyword:   step.Keyword,
			DataTable: step.DataTable,
		})

		stepStartLine := int(step.Location.Line) - 1
		stepEndLine := getStepEndLine(stepStartLine, featureContentLines)

		featureContentLines = slices.Delete(featureContentLines, stepStartLine, stepEndLine+1)

		featureContentLines = slices.Insert(featureContentLines, stepStartLine, expandedSteps...)
	}

	return strings.Join(featureContentLines, "\n")
}

func expandMacroSteps(params ExpandMacroParam) []string {
	variables := getMacroVariables(params.DataTable)

	var expandedSteps []string
	for idx, macroStep := range params.Macro.Steps {
		keyword := params.Keyword
		if idx > 0 {
			keyword = "And"
		}

		stepText := substituteVariables(macroStep.Text, variables)
		expandedSteps = append(expandedSteps, fmt.Sprintf("%s %s", keyword, stepText))

		if macroStep.DocString != nil {
			ds := macroStep.DocString
			delimiter := ds.Delimiter
			if delimiter == "" {
				delimiter = "\"\"\""
			}
			expandedSteps = append(expandedSteps, delimiter)
			expandedSteps = append(expandedSteps, substituteVariables(ds.Content, variables))
			expandedSteps = append(expandedSteps, delimiter)
		}

		if macroStep.DataTable != nil {
			expandedSteps = append(expandedSteps, convertDatatableToString(macroStep.DataTable))
		}
	}
	return expandedSteps
}

func convertDatatableToString(dt *messages.DataTable) string {
	if dt == nil || len(dt.Rows) == 0 {
		return ""
	}

	colWidths := calculateColWidths(dt)

	var sb strings.Builder

	// 2. Build the string
	for _, row := range dt.Rows {
		sb.WriteString(strings.Repeat(" ", int(row.Location.Column)-1))

		sb.WriteString("|")
		for i, cell := range row.Cells {
			val := cell.Value
			width := colWidths[i]
			currentLen := utf8.RuneCountInString(val)

			sb.WriteString(" ")
			sb.WriteString(val)

			// Pad with spaces
			sb.WriteString(strings.Repeat(" ", width-currentLen))
			sb.WriteString(" |")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func calculateColWidths(dt *messages.DataTable) map[int]int {
	colWidths := make(map[int]int)

	for _, row := range dt.Rows {
		for i, cell := range row.Cells {
			// Use RuneCountInString to handle multi-byte characters correctly (e.g. emojis)
			length := utf8.RuneCountInString(cell.Value)
			if length > colWidths[i] {
				colWidths[i] = length
			}
		}
	}
	return colWidths
}

func getStepEndLine(stepStartLine int, featureContent []string) int {
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

func getMacros(features []*Feature) []scenario {
	var macros []scenario

	for _, f := range features {
		if !isFileContainsMacros(f) {
			continue
		}

		for _, sc := range f.scenarios {
			if isMacroScenario(sc) {
				macros = append(macros, *sc)
			}
		}
	}

	return macros
}

func getMacroTitles(macros []scenario) []string {
	var titles []string
	for _, macro := range macros {
		titles = append(titles, macro.Name)
	}

	return titles
}

func isFileContainsMacros(feature *Feature) bool {
	for _, sc := range feature.scenarios {
		for _, tag := range sc.Tags {
			if tag.Name == MacroTag {
				return true
			}
		}
	}
	return false
}

func isMacroScenario(scenario *messages.Scenario) bool {
	if scenario == nil || len(scenario.Tags) == 0 {
		return false
	}
	for _, tag := range scenario.Tags {
		if strings.ToLower(tag.Name) == MacroTag {
			return true
		}
	}
	return false
}

type ExpandMacroParam struct {
	Macro     scenario
	Keyword   string
	DataTable *messages.DataTable
}
