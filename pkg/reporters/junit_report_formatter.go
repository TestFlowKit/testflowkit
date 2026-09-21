package reporters

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	reportJUnitPath = "report/report.xml"
	hooksSuiteName  = "hooks"
)

type junitReportFormatter struct{}

func (f junitReportFormatter) WriteReport(details testSuiteDetails) {
	order, groups := groupByFeature(details.Scenarios)

	report := junitTestSuites{}
	timestamp := details.StartDate.Format(time.RFC3339)

	for _, uri := range order {
		suite := junitTestSuite{Name: uri, Timestamp: timestamp}
		for _, sc := range groups[uri] {
			suite.addCase(uri, sc.Title, sc)
		}
		report.add(suite)
	}

	// Hooks are only reported when they fail, so a broken setup/teardown is visible in CI.
	hooksSuite := junitTestSuite{Name: hooksSuiteName, Timestamp: timestamp}
	for _, sc := range details.Scenarios {
		if sc.Type != scenario && sc.Result == failed {
			hooksSuite.addCase(hooksSuiteName, string(sc.Type), sc)
		}
	}
	if hooksSuite.Tests > 0 {
		report.add(hooksSuite)
	}

	data, err := xml.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Printf("error while serializing the JUnit report: %v\n", err)
		return
	}

	writeReportFile(reportJUnitPath, append([]byte(xml.Header), data...))
}

func (s *junitTestSuite) addCase(className, name string, sc Scenario) {
	tc := junitTestCase{ClassName: className, Name: name, Time: sc.Duration.Seconds()}

	switch {
	case sc.Result == failed:
		tc.Failure = &junitFailure{Message: sc.ErrorMsg, Type: "AssertionError", Content: failureDetails(sc)}
		s.Failures++
	case allStepsSkipped(sc):
		tc.Skipped = &struct{}{}
		s.Skipped++
	}

	s.Tests++
	s.Time += tc.Time
	s.TestCases = append(s.TestCases, tc)
}

func (r *junitTestSuites) add(s junitTestSuite) {
	r.Tests += s.Tests
	r.Failures += s.Failures
	r.Time += s.Time
	r.Suites = append(r.Suites, s)
}

func allStepsSkipped(sc Scenario) bool {
	if len(sc.Steps) == 0 {
		return false
	}
	for _, st := range sc.Steps {
		if st.Status != stepStatusSkipped {
			return false
		}
	}
	return true
}

func failureDetails(sc Scenario) string {
	var b strings.Builder
	for _, st := range sc.Steps {
		fmt.Fprintf(&b, "%s: %s\n", st.Status, st.Title)
	}
	return b.String()
}
