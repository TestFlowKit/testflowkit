package core

import (
	"testflowkit/shared"

	"github.com/cucumber/godog"
)

type stepSupportedTypes interface {
	// Add supported types here
	string | int | float64 | bool | *godog.Table
}

type StepDefDocParams struct {
	Description string
	Variables   []shared.StepVariable
	Example     string
	Category    shared.TestCategory
}
