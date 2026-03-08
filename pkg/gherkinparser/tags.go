package gherkinparser

import (
	"sort"
	"strings"

	"testflowkit/pkg/logger"

	messages "github.com/cucumber/messages/go/v21"
	tagexpressions "github.com/cucumber/tag-expressions/go/v6"
)

// collectTagNames merges tag names from multiple slices of *messages.Tag.
func collectTagNames(tagSlices ...[]*messages.Tag) []string {
	var result []string
	for _, slice := range tagSlices {
		for _, t := range slice {
			if t != nil {
				result = append(result, t.Name)
			}
		}
	}
	return result
}

// scenarioFirstLine returns the 1-based line number of the first line that
// belongs to the scenario (its earliest tag line, or its keyword line).
func scenarioFirstLine(sc *scenario) int {
	first := int(sc.Location.Line)
	for _, tag := range sc.Tags {
		if tag != nil && int(tag.Location.Line) < first {
			first = int(tag.Location.Line)
		}
	}
	return first
}

// filterFeatureContent removes the toRemove scenarios from f's raw content,
// re-parses the result, and returns the updated Feature.
// It processes removals from bottom to top to keep line numbers stable.
func filterFeatureContent(f *Feature, toRemove []*scenario) (*Feature, error) {
	// Collect first line of every top-level block (background + all scenarios).
	var siblingFirstLines []int
	if f.background != nil {
		siblingFirstLines = append(siblingFirstLines, int(f.background.Location.Line))
	}
	for _, sc := range f.scenarios {
		siblingFirstLines = append(siblingFirstLines, scenarioFirstLine(sc))
	}
	sort.Ints(siblingFirstLines)

	lines := strings.Split(string(f.Contents), "\n")

	// Sort removal candidates bottom-to-top so indices stay valid.
	type removal struct{ start, end int }
	var removals []removal

	for _, sc := range toRemove {
		startLine := scenarioFirstLine(sc) // 1-based
		// Find the next sibling's first line.
		endLine := len(lines) // default: remove to EOF
		for _, sib := range siblingFirstLines {
			if sib > startLine {
				endLine = sib
				break
			}
		}
		removals = append(removals, removal{startLine, endLine})
	}

	// Sort descending by start so bottom-up removal keeps indices intact.
	sort.Slice(removals, func(i, j int) bool {
		return removals[i].start > removals[j].start
	})

	for _, r := range removals {
		// Convert to 0-based slice indices.
		from := r.start - 1
		to := r.end - 1
		if from < 0 {
			from = 0
		}
		if to > len(lines) {
			to = len(lines)
		}
		lines = append(lines[:from], lines[to:]...)
	}

	newContent := strings.Join(lines, "\n")
	return parseFeatureContent(newContent)
}

// filterFeatures filters features using a Cucumber tag expression.
// Features with no surviving scenarios are dropped entirely.
// An empty expression returns all features unchanged.
// An invalid expression is treated as a fatal configuration error.
func filterFeatures(features []*Feature, expr string) []*Feature {
	parsed, err := tagexpressions.Parse(expr)
	if err != nil {
		logger.Fatal("Invalid tag expression: "+expr, err)
		return features // unreachable
	}

	var result []*Feature

	for _, f := range features {
		var kept []*scenario
		var removed []*scenario

		for _, sc := range f.scenarios {
			allTags := collectTagNames(f.featureTags, sc.Tags)
			if parsed.Evaluate(allTags) {
				kept = append(kept, sc)
			} else {
				removed = append(removed, sc)
			}
		}

		if len(kept) == 0 {
			// No scenarios survive → drop the whole feature.
			continue
		}

		if len(removed) == 0 {
			// All scenarios survive → return unchanged.
			result = append(result, f)
			continue
		}

		// Some removed → rebuild Contents without the removed scenarios.
		updated, errFilter := filterFeatureContent(f, removed)
		if errFilter != nil {
			// If rebuild fails, keep the original feature to avoid data loss.
			result = append(result, f)
			continue
		}
		result = append(result, updated)
	}

	return result
}
