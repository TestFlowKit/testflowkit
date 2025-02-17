package reporters

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type formatter interface {
	WriteReport(details testSuiteDetails)
}

type disabledFormatter struct {
}

func (f disabledFormatter) WriteReport(details testSuiteDetails) {
	const sentence = "%d tests executed successfully at %s"
	log.Printf(sentence, len(details.Scenarios), details.StartDate)
}

type testSuiteDetails struct {
	AppName, AppVersion string
	ExecutionDate       string
	TotalExecutionTime  string
	TotalTests          string
	SucceededTests      string
	FailedTests         string
	SuccessRate         string
	StartDate           time.Time
	Scenarios           []Scenario
}

func (ts *testSuiteDetails) getTestSuiteDurationInSeconds() int {
	var total time.Duration
	for _, sc := range ts.Scenarios {
		total += sc.Duration
	}
	return int(total.Seconds())
}

func (ts *testSuiteDetails) resume() {
	year, month, day := ts.StartDate.Date()
	dateTime := fmt.Sprintf("%d-%d-%d at %d:%d", month, day, year, ts.StartDate.Hour(), ts.StartDate.Minute())

	total := len(ts.Scenarios)
	succeedSc, failedSc := ts.getScenarioResults()

	ts.TotalTests = strconv.Itoa(total)
	ts.SucceededTests = strconv.Itoa(succeedSc)
	ts.FailedTests = strconv.Itoa(failedSc)
	ts.ExecutionDate = dateTime
	ts.TotalExecutionTime = fmt.Sprintf("%ds", ts.getTestSuiteDurationInSeconds())
	ts.SuccessRate = strconv.Itoa(succeedSc * 100 / total)
}

func (ts *testSuiteDetails) getScenarioResults() (int, int) {
	var succeedSc, failedSc int
	for _, sc := range ts.Scenarios {
		if sc.Result == succeeded {
			succeedSc++
		} else {
			failedSc++
		}
	}
	return succeedSc, failedSc
}
