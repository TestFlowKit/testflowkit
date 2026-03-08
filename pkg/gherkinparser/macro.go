package gherkinparser

import (
	"regexp"
	"strings"
	"testflowkit/pkg/logger"
	"unicode/utf8"

	messages "github.com/cucumber/messages/go/v21"
)

const MacroTag = "@macro"
const excludeMacroTagExpr = "not " + MacroTag

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
	macroHelper := NewMacroHelpers(macros)
	for _, f := range features {
		currFeature, err := macroHelper.ApplyMacroToFeature(*f)
		if err != nil {
			logger.Warn("Error applying macros to feature: "+f.Name, []string{err.Error()})
		}
		newFeatures = append(newFeatures, currFeature)
	}

	return newFeatures
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

func getMacros(features []*Feature) []scenario {
	var macros []scenario

	feats := filterFeatures(features, MacroTag)

	for _, f := range feats {
		for _, sc := range f.scenarios {
			macros = append(macros, *sc)
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

func isMacroScenario(scenario *scenario) bool {
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
