package report

import (
	"fmt"
	"log"
	"time"
)

type Report struct {
	scenarios []Scenario
	startDate time.Time
}

func (r *Report) AddScenario(sc Scenario) {
	text := fmt.Sprintf("%s executed in %fs", sc.title, sc.duration.Seconds())
	log.Println(text)
	r.scenarios = append(r.scenarios, sc)
}

func (r *Report) Start() {
	r.startDate = time.Now()
}

func (r *Report) Write() {
	htmlReport := htmlReportFormatter{}
	htmlReport.WriteReport(r.startDate, r.scenarios)
}

func New() Report {
	return Report{}
}
