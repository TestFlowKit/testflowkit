package gherkinparser

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	messages "github.com/cucumber/messages/go/v21"
)

const MacroTag = "@macro"

type MacroVariable struct {
	Name  string
	Value string
}

type MacroVariables = map[string]string

func getMacroVariables(step *messages.Step) MacroVariables {
	if step.DataTable == nil {
		return make(map[string]string)
	}

	table := step.DataTable
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
	result := stepContent
	for varName, varValue := range variables {
		placeholder := fmt.Sprintf("|%s|", varName)
		result = strings.ReplaceAll(result, placeholder, varValue)
	}

	return result
}

func getCompleteStepContentWhithoutKeyword(step *messages.Step) string {
	result := step.Text

	if step.DocString != nil {
		ds := step.DocString
		docString := fmt.Sprintf("%s\n%s\n%s", ds.Delimiter, ds.Content, ds.Delimiter)

		result = fmt.Sprintf("%s\n%s", result, docString)
	}

	// TODO: in datatable

	return result
}

func applyMacros(macros []*scenario, features []*Feature) {
	macroTitles := getMacroTitles(macros)
	mustContainsMacro := regexp.MustCompile(strings.Join(macroTitles, "|"))

	for _, f := range features {
		content := string(f.Contents)
		if !mustContainsMacro.MatchString(content) {
			continue
		}

		featureContent := strings.Split(content, "\n")
		if f.background != nil {
			applyMacro(f.background.Steps, macros, &featureContent)
		}

		for _, sc := range f.scenarios {
			if sc == nil {
				continue
			}

			var scenarioStepTexts string
			var scenarioStepTextsSb90 strings.Builder
			for _, step := range sc.Steps {
				scenarioStepTextsSb90.WriteString(step.Text + "\n")
			}
			scenarioStepTexts += scenarioStepTextsSb90.String()

			if !mustContainsMacro.MatchString(scenarioStepTexts) {
				continue
			}

			applyMacro(sc.Steps, macros, &featureContent)
		}

		f.Contents = []byte(strings.Join(featureContent, "\n"))
	}
}

func applyMacro(steps []*messages.Step, macros []*scenario, featureContent *[]string) {
	for _, step := range steps {
		macroIdx := slices.IndexFunc(macros, func(macro *scenario) bool {
			return macro.Name == step.Text
		})

		isMacroStep := macroIdx != -1
		if !isMacroStep {
			continue
		}

		expandedSteps := expandMacroSteps(macros[macroIdx], step)

		stepStartLine := int(step.Location.Line) - 1
		stepEndLine := getStepEndLine(stepStartLine, *featureContent)

		*featureContent = slices.Delete(*featureContent, stepStartLine, stepEndLine+1)

		*featureContent = slices.Insert(*featureContent, stepStartLine, expandedSteps...)
	}
}

func expandMacroSteps(macro *scenario, step *messages.Step) []string {
	variables := getMacroVariables(step)

	var expandedSteps []string
	for idx, macroStep := range macro.Steps {
		keyword := step.Keyword
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

		// TODO: Handle datatable if needed
	}
	return expandedSteps
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

func getMacros(features []*Feature) []*scenario {
	var macros []*scenario

	for _, f := range features {
		if !isFileContainsMacros(f) {
			continue
		}

		for _, sc := range f.scenarios {
			if sc == nil || len(sc.Tags) == 0 {
				continue
			}

			if isMacroScenario(sc) {
				macros = append(macros, sc)
			}
		}
	}

	return macros
}

func getMacroTitles(macros []*scenario) []string {
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
	for _, tag := range scenario.Tags {
		if tag.Name == MacroTag {
			return true
		}
	}
	return false
}
