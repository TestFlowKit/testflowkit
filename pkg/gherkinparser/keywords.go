package gherkinparser

import (
	"bytes"
	"strings"

	gherkin "github.com/cucumber/gherkin/go/v26"
	messages "github.com/cucumber/messages/go/v21"
)

// DocStringDelimiter is the default Gherkin doc string delimiter.
const DocStringDelimiter = "\"\"\""

const (
	gherkinKeywordGiven = "Given"
	gherkinKeywordWhen  = "When"
	gherkinKeywordThen  = "Then"
	gherkinKeywordAnd   = "And"
	gherkinKeywordBut   = "But"
)

// DocStringExample wraps body content in a Gherkin doc string block.
func DocStringExample(prefix, body string) string {
	return prefix + "\n" + DocStringDelimiter + "\n" + body + "\n" + DocStringDelimiter
}

// StepKeywordsByID parses a feature file and maps every step's AST node id to its keyword
// (e.g. "Given ", "And "). Ids are generated the same way godog does (a fresh incrementing
// generator per document), so they match the ids found in godog's pickle steps.
func StepKeywordsByID(content []byte) map[string]string {
	doc, err := gherkin.ParseGherkinDocument(bytes.NewReader(content), (&messages.Incrementing{}).NewId)
	if err != nil || doc.Feature == nil {
		return nil
	}

	keywords := map[string]string{}
	collectStepKeywords(doc.Feature.Children, keywords)
	return keywords
}

func collectStepKeywords(children []*messages.FeatureChild, keywords map[string]string) {
	collect := func(steps []*messages.Step) {
		for _, st := range steps {
			keywords[st.Id] = strings.TrimSpace(st.Keyword)
		}
	}

	for _, child := range children {
		switch {
		case child.Background != nil:
			collect(child.Background.Steps)
		case child.Scenario != nil:
			collect(child.Scenario.Steps)
		case child.Rule != nil:
			ruleChildren := make([]*messages.FeatureChild, len(child.Rule.Children))
			for i, rc := range child.Rule.Children {
				ruleChildren[i] = &messages.FeatureChild{Background: rc.Background, Scenario: rc.Scenario}
			}
			collectStepKeywords(ruleChildren, keywords)
		}
	}
}
