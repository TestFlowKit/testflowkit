package common

import "github.com/cucumber/godog"

type stepSupportedTypes interface {
	// Add supported types here
	string | int | float64 | bool | *godog.Table
}
