package gherkinparser

import (
	"sort"
	"strings"

	messages "github.com/cucumber/messages/go/v21"
)

// parseTagExpression converts a tag expression string into an OR-of-AND-groups structure.
//
// Syntax:
//   - "||" separates OR groups
//   - "&&" separates AND tags within a group
//   - e.g. "@smoke || @critical && @api" → [["@smoke"], ["@critical", "@api"]]
//
// An empty expression returns nil (no filter → pass all).
func parseTagExpression(expr string) []tagGroup {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil
	}

	var groups []tagGroup
	for _, orPart := range strings.Split(expr, "||") {
		orPart = strings.TrimSpace(orPart)
		if orPart == "" {
			continue
		}
		var group tagGroup
		for _, tag := range strings.Split(orPart, "&&") {
			tag = strings.TrimSpace(tag)
			if tag != "" {
				group = append(group, tag)
			}
		}
		if len(group) > 0 {
			groups = append(groups, group)
		}
	}

	if len(groups) == 0 {
		return nil
	}
	return groups
}

// matchesFilter reports whether tags satisfies the OR-of-AND filter groups.
// Tags prefixed with "~" are negations: the scenario must NOT have that tag.
// Empty groups means no filter → always true.
func matchesFilter(tags []string, groups []tagGroup) bool {
	if len(groups) == 0 {
		return true
	}

	tagSet := make(map[string]struct{}, len(tags))
	for _, t := range tags {
		tagSet[t] = struct{}{}
	}

	for _, group := range groups {
		allMatch := true
		for _, term := range group {
			if strings.HasPrefix(term, "~") {
				// negation: tag must NOT be present
				if _, ok := tagSet[term[1:]]; ok {
					allMatch = false
					break
				}
			} else {
				if _, ok := tagSet[term]; !ok {
					allMatch = false
					break
				}
			}
		}
		if allMatch {
			return true
		}
	}
	return false
}

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

	// Build a set of first lines to remove.
	removeSet := make(map[int]struct{}, len(toRemove))
	for _, sc := range toRemove {
		removeSet[scenarioFirstLine(sc)] = struct{}{}
	}

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

// filterFeatures filters features by a parsed tag expression (OR-of-AND groups).
// Negated tags (~@tag) mean the scenario must NOT have that tag.
// Features with no surviving scenarios are dropped entirely.
func filterFeatures(features []*Feature, groups []tagGroup) []*Feature {
	var result []*Feature

	for _, f := range features {
		var kept []*scenario
		var removed []*scenario

		for _, sc := range f.scenarios {
			allTags := collectTagNames(f.featureTags, sc.Tags)
			if matchesFilter(allTags, groups) {
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
		updated, err := filterFeatureContent(f, removed)
		if err != nil {
			// If rebuild fails, keep the original feature to avoid data loss.
			result = append(result, f)
			continue
		}
		result = append(result, updated)
	}

	return result
}
