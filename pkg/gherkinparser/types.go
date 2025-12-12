package gherkinparser

import messages "github.com/cucumber/messages/go/v21"

func newFeature(name string, content []byte, scenarios []*scenario, background *messages.Background) *Feature {
	return &Feature{
		Name:       name,
		Contents:   content,
		scenarios:  scenarios,
		background: background,
	}
}

type Feature struct {
	Name       string
	Contents   []byte
	scenarios  []*scenario
	background *messages.Background
}

type scenario = messages.Scenario
