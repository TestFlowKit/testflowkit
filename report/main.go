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

func New(appName, appVersion string, enabled bool, formatType string) Report {
	if !enabled {
		formatType = "DISABLED"
	}

	reportFormatter, err := getFormatter(formatType)
	if err != nil {
		log.Panic(err)
	}

	return Report{
		formatter:  reportFormatter,
		appName:    appName,
		appVersion: appVersion,
	}
}

func getFormatter(formatType string) (f formatter, err error) {
	switch formatType {
	case "html":
		f = htmlReportFormatter{}
	case "DISABLED":
		f = disabledFormatter{}
	default:
		return f, fmt.Errorf("'%s' report format not supported", formatType)
	}

	return f, nil
}
