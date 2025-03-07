package reporters

import (
	_ "embed"
	"html/template"
	"log"
	"os"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/utils"
)

//go:embed html_report.template.html
var reportTemplate string

type htmlReportFormatter struct{}

func (r htmlReportFormatter) format(ts testSuiteDetails) string {
	tmpl, err := template.New("report").Parse(reportTemplate)
	if err != nil {
		logger.Fatal("cannot parse report template", err)
	}

	ts.resume()

	wr := utils.TextWriter{}
	err = tmpl.Execute(&wr, ts)
	if err != nil {
		logger.Fatal("cannot execute template", err)
	}

	return wr.String()
}

func (r htmlReportFormatter) WriteReport(details testSuiteDetails) {
	content := r.format(details)

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
