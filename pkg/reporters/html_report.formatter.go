package reporters

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cucumber/godog"
)

//go:embed html_templates/report.template.html
var reportTemplate string

//go:embed html_templates/scenario.template.html
var scenarioTemplate string

//go:embed html_templates/step.template.html
var stepTemplate string

type templates struct {
	report, scenario, step string
}

type htmlReportFormatter struct{}

func (r htmlReportFormatter) fillReport(p fillHTMLReportParams) string {
	year, month, day := p.startDate.Date()
	dateTime := fmt.Sprintf("%d-%d-%d at %d:%d", month, day, year, p.startDate.Hour(), p.startDate.Minute())

	total := len(p.scenarios)
	var content string
	var succeed, failed int
	for _, sc := range p.scenarios {
		if len(sc.err) > 0 {
			failed++
		} else {
			succeed++
		}

		content += fmt.Sprintln(r.fillScenarioTemplate(sc, p.templates.scenario, p.templates.step))
	}

	reportVars := getReportVariables()
	testSuiteReport := r.setTemplateVar(p.templates.report, reportVars.scenariosTemplate, content)
	testSuiteReport = r.setTemplateVar(testSuiteReport, reportVars.executionDate, dateTime)
	testSuiteReport = r.setTemplateVar(testSuiteReport, reportVars.totalTests, strconv.Itoa(total))
	testSuiteReport = r.setTemplateVar(testSuiteReport, reportVars.succeededTests, strconv.Itoa(succeed))
	testSuiteReport = r.setTemplateVar(testSuiteReport, reportVars.failedTests, strconv.Itoa(failed))
	testSuiteReport = r.setTemplateVar(testSuiteReport, reportVars.appName, p.appName)
	testSuiteReport = r.setTemplateVar(testSuiteReport, reportVars.appVersion, p.appVersion)

	totalDuration := fmt.Sprintf("%ds", p.getTestSuiteDurationInSeconds())
	testSuiteReport = r.setTemplateVar(testSuiteReport, reportVars.totalExecutionTime, totalDuration)

	const totalRate = 100
	successRate := succeed * totalRate / total
	testSuiteReport = r.setTemplateVar(testSuiteReport, reportVars.successRate, strconv.Itoa(successRate))

	return testSuiteReport
}

func (r htmlReportFormatter) fillScenarioTemplate(sc Scenario, scenarioTemplate, stepTemplate string) string {
	vars := getScenarioVariables()
	formattedTemplate := scenarioTemplate
	formattedTemplate = r.setTemplateVar(formattedTemplate, vars.name, sc.title)

	duration := sc.duration.Seconds()
	duration = math.Max(math.Ceil(duration), duration)
	formattedTemplate = r.setTemplateVar(formattedTemplate, vars.duration, fmt.Sprintf("%vs", duration))
	result, color := "FAILED", "red"
	if sc.err == "" {
		result, color, sc.err = "SUCCEEDED", "green", "-"
	}
	formattedTemplate = r.setTemplateVar(formattedTemplate, vars.result, result)
	formattedTemplate = r.setTemplateVar(formattedTemplate, vars.statusColor, color)
	formattedTemplate = r.setTemplateVar(formattedTemplate, vars.errorMessage, sc.err)

	var filledSteps string
	for _, step := range sc.steps {
		filledSteps += fmt.Sprintln(r.fillStepTemplate(step, stepTemplate))
	}

	formattedTemplate = r.setTemplateVar(formattedTemplate, vars.stepsTemplate, filledSteps)
	return formattedTemplate
}

func (r htmlReportFormatter) fillStepTemplate(step Step, template string) string {
	getColor := func(status godog.StepResultStatus) string {
		if status == 0 {
			return "green"
		}

		if status == 1 {
			return "red"
		}

		return "gray"
	}

	vars := getStepVariables()
	tpl := r.setTemplateVar(template, vars.title, step.title)
	tpl = r.setTemplateVar(tpl, vars.status, step.status.String())
	tpl = r.setTemplateVar(tpl, vars.statusColor, getColor(step.status))
	tpl = r.setTemplateVar(tpl, vars.duration, fmt.Sprintf("%dms", step.duration.Milliseconds()))
	return strings.TrimSpace(tpl)
}

func (r htmlReportFormatter) setTemplateVar(template, variableName, value string) string {
	variable := fmt.Sprintf("{{ %s }}", variableName)
	return strings.ReplaceAll(template, variable, value)
}

func (r htmlReportFormatter) WriteReport(details testSuiteDetails) {
	content := r.fillReport(fillHTMLReportParams{
		testSuiteDetails: details,
		templates: templates{
			report:   reportTemplate,
			scenario: scenarioTemplate,
			step:     stepTemplate,
		},
	})

	file, err := os.Create("report.html")
	if err != nil {
		log.Panicf("cannot create reporters file in this folder ( %s )\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Panicf("error when reporters filling ( %s )", err)
	}
}

type fillHTMLReportParams struct {
	testSuiteDetails
	templates
}

func (params fillHTMLReportParams) getTestSuiteDurationInSeconds() int {
	var total time.Duration
	for _, sc := range params.scenarios {
		total += sc.duration
	}
	return int(total.Seconds())
}
