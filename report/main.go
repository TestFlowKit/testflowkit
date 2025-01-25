package report

import (
	"fmt"
	"log"
	"time"
)

type Report struct {
	appName, appVersion string
	scenarios           []Scenario
	startDate           time.Time
	formatter           formatter
}

func (r *Report) AddScenario(sc Scenario) {
	result := "succeeded"

	if len(sc.err) > 0 {
		result = "failed"
	}

	addedScenarioLoggedMessage := fmt.Sprintf("'%s' %s in %fs", sc.title, result, sc.duration.Seconds())
	log.Println(addedScenarioLoggedMessage)
	r.scenarios = append(r.scenarios, sc)
}

func (r *Report) Start() {
	r.startDate = time.Now()
}

func (r *Report) Write() {
	r.formatter.WriteReport(testSuiteDetails{
		appName:    r.appName,
		appVersion: r.appVersion,
		startDate:  r.startDate,
		scenarios:  r.scenarios,
	})
}

func (r *Report) HasScenarios() bool {
	return len(r.scenarios) > 0
}

func New(appName, appVersion string, formatType string) Report {
	reportFormatter := getFormatter(formatType)
	return Report{
		formatter:  reportFormatter,
		appName:    appName,
		appVersion: appVersion,
	}
}

func getFormatter(formatType string) formatter {
	switch formatType {
	case "html":
		return htmlReportFormatter{}
	default:
		log.Printf("'%s' report format not supported\n", formatType)
		return disabledFormatter{}
	}
}
