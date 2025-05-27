package reporters

import (
	"fmt"
	"log"
	"testflowkit/pkg/logger"
	"time"
)

type scenarioResult string

const (
	succeeded scenarioResult = "succeeded"
	failed    scenarioResult = "failed"
)

type Report struct {
	appName, appVersion string
	scenarios           []Scenario
	startDate           time.Time
	formatter           formatter
	AreAllTestsPassed   bool
}

func (r *Report) AddScenario(sc Scenario) {
	r.scenarios = append(r.scenarios, sc)

	result := sc.Result
	addedScenarioLoggedMessage := fmt.Sprintf("'%s' %s in %fs", sc.Title, result, sc.Duration.Seconds())

	if result == failed {
		r.AreAllTestsPassed = false
		logger.Error(addedScenarioLoggedMessage, nil, nil)
	} else {
		logger.Success(addedScenarioLoggedMessage)
	}
}

func (r *Report) Start() {
	r.startDate = time.Now()
}

func (r *Report) Write() {
	ts := newTestSuiteDetails(
		r.appName,
		r.appVersion,
		r.startDate,
		r.scenarios,
	)

	r.formatter.WriteReport(*ts)
}

func (r *Report) HasScenarios() bool {
	return len(r.scenarios) > 0
}

func New(appName, appVersion string, formatType string) Report {
	reportFormatter := getFormatter(formatType)
	return Report{
		formatter:         reportFormatter,
		appName:           appName,
		appVersion:        appVersion,
		AreAllTestsPassed: true,
	}
}

func getFormatter(formatType string) formatter {
	switch formatType {
	case "html":
		return htmlReportFormatter{}
	case "json":
		return jsonReportFormatter{}
	default:
		log.Printf("'%s' report format not supported\n", formatType)
		return disabledFormatter{}
	}
}
