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
	text := fmt.Sprintf("%s executed in %fs", sc.title, sc.duration.Seconds())
	log.Println(text)
	r.scenarios = append(r.scenarios, sc)
}

func (r *Report) Start() {
	r.startDate = time.Now()
}

func (r *Report) Write() {
	r.formatter.WriteReport(r.startDate, r.scenarios)
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
