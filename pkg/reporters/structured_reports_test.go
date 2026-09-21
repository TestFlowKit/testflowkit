package reporters

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sampleDetails() testSuiteDetails {
	scenarios := []Scenario{
		{
			Title: "ok", ID: "1", URI: "login", Tags: []string{"@smoke"}, Type: scenario,
			Result: succeeded, Duration: 2 * time.Second,
			Steps: []Step{{Title: "open", Keyword: "Given", Status: "passed", Duration: time.Millisecond}},
		},
		{
			Title: "ko", ID: "2", URI: "login", Type: scenario,
			Result: failed, ErrorMsg: "boom", Duration: time.Second,
			Steps: []Step{{Title: "click", Status: "failed", ScreenshotBase64: "abc"}},
		},
		{Title: "skipped", ID: "3", URI: "cart", Type: scenario, Result: succeeded,
			Steps: []Step{{Title: "x", Status: "skipped"}}},
		{Type: beforeHook, Result: succeeded},
	}
	return *newTestSuiteDetails(time.Now(), scenarios)
}

func TestCucumberJSONReport(t *testing.T) {
	t.Chdir(t.TempDir())
	jsonReportFormatter{}.WriteReport(sampleDetails())

	data, err := os.ReadFile("report/report.json")
	require.NoError(t, err)

	var features []cucumberFeature
	require.NoError(t, json.Unmarshal(data, &features))
	require.Len(t, features, 2)
	assert.Equal(t, "login", features[0].URI)
	assert.Equal(t, "Feature", features[0].Keyword)
	require.Len(t, features[0].Elements, 2)
	assert.Equal(t, "@smoke", features[0].Elements[0].Tags[0].Name)
	assert.Equal(t, "Given", features[0].Elements[0].Steps[0].Keyword)
	assert.Equal(t, int64(time.Millisecond), features[0].Elements[0].Steps[0].Result.Duration)

	failedStep := features[0].Elements[1].Steps[0]
	assert.Equal(t, "failed", failedStep.Result.Status)
	assert.Equal(t, "boom", failedStep.Result.ErrorMessage)
	assert.Equal(t, "abc", failedStep.Embeddings[0].Data)
}

func TestJUnitReport(t *testing.T) {
	t.Chdir(t.TempDir())
	junitReportFormatter{}.WriteReport(sampleDetails())

	data, err := os.ReadFile("report/report.xml")
	require.NoError(t, err)

	var suites junitTestSuites
	require.NoError(t, xml.Unmarshal(data, &suites))
	assert.Equal(t, 3, suites.Tests)
	assert.Equal(t, 1, suites.Failures)
	require.Len(t, suites.Suites, 2)
	assert.Equal(t, 1, suites.Suites[0].Failures)
	assert.Equal(t, "boom", suites.Suites[0].TestCases[1].Failure.Message)
	assert.Equal(t, 1, suites.Suites[1].Skipped)
}
