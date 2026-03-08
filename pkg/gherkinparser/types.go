package gherkinparser

import messages "github.com/cucumber/messages/go/v21"

func newFeature(params NewFeatureParams) *Feature {
	return &Feature{
		Name:        params.Name,
		Contents:    params.Content,
		scenarios:   params.Scenarios,
		background:  params.Background,
		featureTags: params.Tags,
	}
}

type NewFeatureParams struct {
	Name       string
	Content    []byte
	Scenarios  []*scenario
	Background *messages.Background
	Tags       []*messages.Tag
}

type Feature struct {
	Name        string
	Contents    []byte
	scenarios   []*scenario
	background  *messages.Background
	featureTags []*messages.Tag
}

// tagGroup is an AND group: all tags must match.
// Tags prefixed with "~" are negations: the scenario must NOT have that tag.
type tagGroup = []string

type scenario = messages.Scenario

// HasAnyStep reports whether any step text across the background and all scenarios
// of the given features satisfies the match function. It returns true on the first match.
func HasAnyStep(features []*Feature, match func(string) bool) bool {
	for _, f := range features {
		if f == nil {
			continue
		}
		if f.background != nil {
			if stepMatches(f.background.Steps, match) {
				return true
			}
		}
		for _, sc := range f.scenarios {
			if sc == nil {
				continue
			}

			if stepMatches(sc.Steps, match) {
				return true
			}
		}
	}
	return false
}

func stepMatches(steps []*messages.Step, match func(string) bool) bool {
	for _, step := range steps {
		if step == nil {
			continue
		}
		if match(step.Text) {
			return true
		}
	}
	return false
}
