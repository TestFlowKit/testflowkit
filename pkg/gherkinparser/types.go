package gherkinparser

import messages "github.com/cucumber/messages/go/v21"

func newFeature(name, fileURL string, content []byte, scenarios []*scenario) *Feature {
	return &Feature{
		Name:      name,
		Contents:  content,
		uri:       fileURL,
		scenarios: scenarios,
	}
}

type Feature struct {
	Name      string
	Contents  []byte
	uri       string
	scenarios []*scenario
}

type scenario = messages.Scenario
