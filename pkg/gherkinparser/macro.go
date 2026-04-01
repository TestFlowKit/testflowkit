package gherkinparser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testflowkit/pkg/logger"

	messages "github.com/cucumber/messages/go/v21"
)

const MacroTag = "@macro"
const excludeMacroTagExpr = "not " + MacroTag
const errMsg = "Invalid macro variables table: each row must have exactly 2 cells (variable name and value)"

var macroVarPattern = regexp.MustCompile(`\$\{([^}]+)\}`)
var errMacroVariableNotFound = errors.New("macro variable not found")

type MacroVariable struct {
	Name  string
	Value string
}

type MacroVariables = map[string]string

// getMacroVariables converts a Gherkin DataTable to a map for variable substitution.
// Uses vertical/two-column format where each row contains: [variable_name, value]
// Example DataTable:
//
//	| email    | user@example.com |
//	| password | secret123        |
//
// Returns: {email: "user@example.com", password: "secret123"}.
func getMacroVariables(table *messages.DataTable) MacroVariables {
	if table == nil {
		return make(map[string]string)
	}

	const expectedCellsPerRow = 2

	variables := make(map[string]string)
	for _, r := range table.Rows {
		cells := r.Cells
		if len(cells) != expectedCellsPerRow {
			logger.Warn(errMsg, nil)
			continue
		}

		varName := strings.TrimSpace(cells[0].Value)
		varValue := strings.TrimSpace(cells[1].Value)

		variables[varName] = varValue
	}

	return variables
}

// substituteVariables replaces ${variable_name} placeholders in step content with values from the map.
// Returns an error when a referenced variable is missing so macro expansion cannot continue partially.
func substituteVariables(stepContent string, variables MacroVariables) (string, error) {
	missingVar := ""

	result := macroVarPattern.ReplaceAllStringFunc(stepContent, func(match string) string {
		if missingVar != "" {
			return match
		}

		varName := strings.TrimSpace(match[2 : len(match)-1])
		if value, exists := variables[varName]; exists {
			return value
		}

		missingVar = varName
		return match
	})

	if missingVar != "" {
		return "", fmt.Errorf("%w: %s", errMacroVariableNotFound, missingVar)
	}

	return result, nil
}

// TODO: refactor for better understanding.
func applyMacros(macros []scenario, features []*Feature) []*Feature {
	newFeatures := make([]*Feature, 0, len(features))
	macroHelper := NewMacroHelpers(macros)
	for _, f := range features {
		currFeature, err := macroHelper.ApplyMacroToFeature(*f)
		if err != nil {
			logger.Warn("Error applying macros to feature: "+f.Name, []string{err.Error()})
			newFeatures = append(newFeatures, f)
			continue
		}
		newFeatures = append(newFeatures, currFeature)
	}

	return newFeatures
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

// convertDatatableToString converts a DataTable to its string representation for output.
// Each cell value is expanded using the same ${variable} substitution rules as step text.
func convertDatatableToString(dt *messages.DataTable, variables MacroVariables) (string, error) {
	if dt == nil || len(dt.Rows) == 0 {
		return "", nil
	}

	var sb strings.Builder
	for _, row := range dt.Rows {
		indent := 0
		if row.Location != nil && row.Location.Column > 0 {
			indent = int(row.Location.Column) - 1
		}
		sb.WriteString(strings.Repeat(" ", indent))
		sb.WriteString("|")
		for _, cell := range row.Cells {
			cellValue, err := substituteVariables(cell.Value, variables)
			if err != nil {
				return "", err
			}
			sb.WriteString(" ")
			sb.WriteString(cellValue)
			sb.WriteString(" |")
		}
		sb.WriteString("\n")
	}
	return sb.String(), nil
}

// ExpandMacroParam encapsulates parameters for macro expansion.
// Variables field contains the map of variable names to values for substitution,
// converted from the DataTable at the call site for efficient variable lookup.
type ExpandMacroParam struct {
	Macro     scenario
	Keyword   string
	Variables MacroVariables // Map-based variables for efficient ${variable} substitution
}
