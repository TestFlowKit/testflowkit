package report

import (
	"log"
	"time"
)

type formatter interface {
	WriteReport(startDate time.Time, scenarios []Scenario)
}

type disabledFormatter struct {
}

func (f disabledFormatter) WriteReport(startDate time.Time, scenarios []Scenario) {
	log.Printf("%d tests executed successfully at %s / report generation disabled\n", len(scenarios), startDate)
}
