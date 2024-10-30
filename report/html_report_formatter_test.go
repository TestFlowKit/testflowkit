package report

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatSuccessCaseScenarioTemplate(t *testing.T) {
	formatter := htmlReportFormatter{}

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

	scenarioTemplate := `<details>
{{ SCENARIO_NAME }}{{ SCENARIO_RESULT }}{{ SCENARIO_DURATION }}{{ SCENARIO_ERROR_MESSAGE }}
<details>`
	firstTpl := formatter.fillScenarioTemplate(first, scenarioTemplate, "")

	assert.Equal(t, firstTpl, strings.TrimSpace(`<details>
firstSUCCEEDED10s-
<details>`))
}

func TestFormatFailedCaseScenarioTemplate(t *testing.T) {
	formatter := htmlReportFormatter{}
	failed := newScenario("connect", make([]Step, 0), 3*time.Second, errors.New("error"))

	tpl := `
		<details>
		{{ SCENARIO_NAME }}
		{{ SCENARIO_RESULT }}
        {{ SCENARIO_DURATION }}
        {{ SCENARIO_ERROR_MESSAGE }}
		<details>
`
	firstTpl := formatter.fillScenarioTemplate(failed, tpl, "")

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
	step := Step{title: "etape", status: 0}
	tpl := "{{ STEP_TITLE }} / {{ STEP_STATUS_COLOR }} / {{ STEP_STATUS }}"
	expected := "etape / green / passed"

	assert.Equal(t, expected, htmlFormatter.fillStepTemplate(step, tpl))
}

func TestFormatReport(t *testing.T) {
	htmlFormatter := htmlReportFormatter{}

	reportTpl := `
{{ EXECUTION_DATE }}-{{ TOTAL_TESTS }}{{ SUCCEEDED_TESTS }}{{ FAILED_TESTS }}-{{ SUCCESS_RATE }}-{{ SCENARIOS }}
`
	scTpl := `{{ SCENARIO_NAME }}{{ STEPS }}`
	stepTpl := "{{ STEP_TITLE }}"

	const expected = "12-10-2024 at 10:0-110-100-SCetape"

	sc := Scenario{title: "SC", steps: []Step{{title: "etape", status: 0}}}

	startDate := time.Date(2024, 12, 10, 10, 00, 00, 00, time.Local)

	reportFormatted := htmlFormatter.fillReport(startDate, []Scenario{sc}, templates{
		report:   reportTpl,
		scenario: scTpl,
		step:     stepTpl,
	})

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(reportFormatted))
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
