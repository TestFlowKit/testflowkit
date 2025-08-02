package gherkinparser

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	messages "github.com/cucumber/messages/go/v21"
)

const MacroTag = "@macro"

func applyMacros(macros []*scenario, featuresContainingMacros []*Feature) {
	macroTitles := getMacroTitles(macros)
	mustContainsMacro := regexp.MustCompile(strings.Join(macroTitles, "|"))
	for _, f := range featuresContainingMacros {
		content := string(f.Contents)
		if !mustContainsMacro.MatchString(content) {
			continue
		}

		featureContent := strings.Split(content, "\n")

		if f.background != nil {
			applyMacro(f.background.Steps, macroTitles, macros, featureContent)
		}

		for _, sc := range f.scenarios {
			if sc == nil {
				continue
			}

			var scenarioContent string
			for _, step := range sc.Steps {
				scenarioContent += step.Text + "\n"
			}

			if !mustContainsMacro.MatchString(scenarioContent) {
				continue
			}

			applyMacro(sc.Steps, macroTitles, macros, featureContent)
		}

		f.Contents = []byte(strings.Join(featureContent, "\n"))
	}
}

func applyMacro(steps []*messages.Step, macroTitles []string, macros []*scenario, featureContent []string) {
	for _, step := range steps {
		isMacroStep := slices.Contains(macroTitles, step.Text)
		if !isMacroStep {
			continue
		}

		macroIdx := slices.IndexFunc(macroTitles, func(title string) bool {
			return title == step.Text
		})

		macro := macros[macroIdx]
		var steps []string
		for idx, macroStep := range macro.Steps {
			keyword := step.Keyword
			if idx > 0 {
				keyword = "And"
			}

			steps = append(steps, fmt.Sprintf("%s %s", keyword, macroStep.Text))
		}

		featureContent[step.Location.Line-1] = strings.Join(steps, "\n")
	}
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
