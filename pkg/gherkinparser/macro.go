package gherkinparser

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func applyMacros(macros []*scenario, featuresContainingMacros []*Feature) {
	macroTitles := getMacroTitles(macros)
	for _, f := range featuresContainingMacros {
		featureContent := strings.Split(string(f.Contents), "\n")

		for _, sc := range f.scenarios {
			if sc == nil {
				continue
			}

			var scenarioContent string
			for _, step := range sc.Steps {
				scenarioContent += step.Text + "\n"
			}

			r := regexp.MustCompile(strings.Join(macroTitles, "|"))
			if !r.MatchString(scenarioContent) {
				continue
			}

			applyMacro(sc, macroTitles, macros, featureContent)
		}

		f.Contents = []byte(strings.Join(featureContent, "\n"))
	}
}

func applyMacro(sc *scenario, macroTitles []string, macros []*scenario, featureContent []string) {
	for _, step := range sc.Steps {
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

func getMacros(macroFeatures []*Feature) []*scenario {
	var macros []*scenario

	for _, sc := range macroFeatures {
		macros = append(macros, sc.scenarios...)
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
