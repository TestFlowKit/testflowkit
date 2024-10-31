package report

import (
	"log"
	"time"
)

type formatter interface {
	WriteReport(details testSuiteDetails)
}

type disabledFormatter struct {
}

func (f disabledFormatter) WriteReport(details testSuiteDetails) {
	log.Printf("%d tests executed successfully at %s / report generation disabled\n", len(details.scenarios), details.startDate)
}

type testSuiteDetails struct {
	appName, appVersion string
	startDate           time.Time
	scenarios           []Scenario
}
