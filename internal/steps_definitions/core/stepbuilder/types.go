package stepbuilder

import (
	"github.com/cucumber/godog"
)

type supportedTypes interface {
	// Add supported types here
	string | *godog.Table
}

type DocParams struct {
	Description string
	Variables   []DocVariable
	Example     string
	Category    StepCategory
}
