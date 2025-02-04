package reporters

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
	const sentence = "%d tests executed successfully at %s"
	log.Printf(sentence, len(details.scenarios), details.startDate)
}

type testSuiteDetails struct {
	appName, appVersion string
	startDate           time.Time
	scenarios           []Scenario
}
