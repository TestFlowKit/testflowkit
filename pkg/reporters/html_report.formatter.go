package reporters

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/utils"
)

//go:embed html_report.template.html
var reportTemplate string

type htmlReportFormatter struct{}

func (r htmlReportFormatter) format(ts htmlTestSuiteDetails) string {
	tmpl, err := template.New("report").Parse(reportTemplate)
	if err != nil {
		logger.Fatal("cannot parse report template", err)
	}

	wr := utils.TextWriter{}
	err = tmpl.Execute(&wr, ts)
	if err != nil {
		logger.Fatal("cannot execute template", err)
	}

	return wr.String()
}

func (r htmlReportFormatter) WriteReport(details testSuiteDetails) {
	content := r.formatContent(details)
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

func (r htmlReportFormatter) formatContent(details testSuiteDetails) string {
	htmlScenarios := make([]htmlScenario, len(details.Scenarios))
	for i, sc := range details.Scenarios {
		htmlScenarios[i] = *newhtmlScenario(sc)
	}

	return r.format(htmlTestSuiteDetails{
		testSuiteDetails: details,
		Scenarios:        htmlScenarios,
	})
}

type htmlScenario struct {
	Scenario
	FmtDuration          string
	HTMLStatusColorClass string
}

type htmlTestSuiteDetails struct {
	testSuiteDetails
	Scenarios []htmlScenario
}

func newhtmlScenario(sc Scenario) *htmlScenario {
	color := "red"
	if sc.Result == succeeded {
		color = "green"
	}

	if sc.ErrorMsg == "" {
		sc.ErrorMsg = "-"
	}

	return &htmlScenario{
		Scenario:             sc,
		FmtDuration:          sc.Duration.String(),
		HTMLStatusColorClass: fmt.Sprintf("bg-%s-500", color),
	}
}
