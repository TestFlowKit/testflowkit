package reporters

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatSuccessCaseScenarioTemplate(t *testing.T) {
	reportFormatter := htmlReportFormatter{}

	first := newScenario("first", []Step{
		{
			title:  "je",
			status: 0,
		},
		{
			title:  "fais",
			status: 0,
		},
	}, 10*time.Second, nil)

	scenarioTpl := `<details>
{{ SCENARIO_NAME }}{{ SCENARIO_RESULT }}{{ SCENARIO_DURATION }}{{ SCENARIO_ERROR_MESSAGE }}
<details>`
	firstTpl := reportFormatter.fillScenarioTemplate(first, scenarioTpl, "")

	assert.Equal(t, firstTpl, strings.TrimSpace(`<details>
firstSUCCEEDED10s-
<details>`))
}

func TestFormatFailedCaseScenarioTemplate(t *testing.T) {
	htmlFormatter := htmlReportFormatter{}
	failed := newScenario("connect", make([]Step, 0), 3*time.Second, errors.New("error"))

	tpl := `
		<details>
		{{ SCENARIO_NAME }}
		{{ SCENARIO_RESULT }}
        {{ SCENARIO_DURATION }}
        {{ SCENARIO_ERROR_MESSAGE }}
		<details>
`
	firstTpl := htmlFormatter.fillScenarioTemplate(failed, tpl, "")

	assert.Contains(t, firstTpl, strings.TrimSpace(`
		<details>
		connect
		FAILED
        3s
        error
		<details>
	`))
}

func TestFormatStepTemplate(t *testing.T) {
	htmlFormatter := htmlReportFormatter{}
	step := Step{title: "etape", status: 0, duration: 61 * time.Millisecond}
	tpl := "{{ STEP_TITLE }} / {{ STEP_STATUS_COLOR }} / {{ STEP_STATUS }} / {{ STEP_DURATION }}"
	expected := "etape / green / passed / 61ms"

	assert.Equal(t, expected, htmlFormatter.fillStepTemplate(step, tpl))
}

func TestFormatReport(t *testing.T) {
	htmlFormatter := htmlReportFormatter{}

	reportTpl := `
{{ APP_NAME }} {{ APP_VERSION }} {{ TOTAL_EXECUTION_TIME }}
{{ EXECUTION_DATE }}-{{ TOTAL_TESTS }}{{ SUCCEEDED_TESTS }}{{ FAILED_TESTS }}-{{ SUCCESS_RATE }}-{{ SCENARIOS }}
`
	scTpl := `{{ SCENARIO_NAME }}{{ STEPS }}`
	stepTpl := "{{ STEP_TITLE }}"

	const expected = `
My app 1.0.0 10s
12-10-2024 at 10:0-110-100-SCetape
`

	sc := Scenario{title: "SC", steps: []Step{{title: "etape", status: 0}}, duration: 10 * time.Second}

	startDate := time.Date(2024, 12, 10, 10, 00, 00, 00, time.Local)

	reportFormatted := htmlFormatter.fillReport(fillHTMLReportParams{
		testSuiteDetails: testSuiteDetails{
			appName:    "My app",
			appVersion: "1.0.0",
			startDate:  startDate,
			scenarios:  []Scenario{sc},
		},
		templates: templates{
			report:   reportTpl,
			scenario: scTpl,
			step:     stepTpl,
		},
	})

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(reportFormatted))
}

func TestTotalDuration(t *testing.T) {
	params := fillHTMLReportParams{}
	params.scenarios = append(params.scenarios, Scenario{duration: 2 * time.Second})
	params.scenarios = append(params.scenarios, Scenario{duration: 10 * time.Second})
	params.scenarios = append(params.scenarios, Scenario{duration: 5 * time.Minute})

	assert.Equal(t, 312, params.getTestSuiteDurationInSeconds())
}

func newScenario(title string, steps []Step, duration time.Duration, err error) Scenario {
	sc := Scenario{
		title:    title,
		steps:    steps,
		duration: duration,
	}

	if err != nil {
		sc.err = err.Error()
	}

	return sc
}
