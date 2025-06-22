package reporters

import (
	"fmt"
	"strconv"
	"time"
)

type testSuiteDetails struct {
	ExecutionDate      string
	TotalExecutionTime string
	TotalTests         string
	SucceededTests     string
	FailedTests        string
	SuccessRate        string
	StartDate          time.Time
	Scenarios          []Scenario
}

func (ts *testSuiteDetails) getTestSuiteDurationInSeconds() int {
	var total time.Duration
	for _, sc := range ts.Scenarios {
		total += sc.Duration
	}
	return int(total.Seconds())
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

func newTestSuiteDetails(startDate time.Time, scenarios []Scenario) *testSuiteDetails {
	ts := testSuiteDetails{
		StartDate: startDate,
		Scenarios: scenarios,
	}

	dateTime := ts.StartDate.Format("01-02-2006 at 15:04")

	total := len(ts.Scenarios)
	succeedSc, failedSc := ts.getScenarioResults()

	ts.TotalTests = strconv.Itoa(total)
	ts.SucceededTests = strconv.Itoa(succeedSc)
	ts.FailedTests = strconv.Itoa(failedSc)
	ts.ExecutionDate = dateTime
	ts.TotalExecutionTime = fmt.Sprintf("%ds", ts.getTestSuiteDurationInSeconds())
	if total > 0 {
		ts.SuccessRate = strconv.Itoa(succeedSc * 100 / total)
	} else {
		ts.SuccessRate = "0"
	}
	return &ts
}
