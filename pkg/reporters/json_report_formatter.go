package reporters

import (
	"encoding/json"
	"log"
	"os"
	"testflowkit/internal/utils/fileutils"
)

type jsonReportFormatter struct{}

func (f jsonReportFormatter) WriteReport(details testSuiteDetails) {
	scenariosReports := make([]jsonScenarioReport, len(details.Scenarios))
	for i, sc := range details.Scenarios {
		scenariosReports[i] = jsonScenarioReport{
			Title:        sc.Title,
			Duration:     sc.Duration.String(),
			Result:       string(sc.Result),
			Steps:        make([]jsonScenarioStepReport, len(sc.Steps)),
			ErrorMessage: sc.ErrorMsg,
		}

		for j, step := range sc.Steps {
			scenariosReports[i].Steps[j] = jsonScenarioStepReport{
				Title:          step.Title,
				Status:         step.Status,
				Duration:       step.Duration.String(),
				ScreenshotPath: step.ScreenshotPath,
			}
		}
	}

	report := jsonReport{
		Scenarios: scenariosReports,
		StartDate: details.ExecutionDate,
		Duration:  details.TotalExecutionTime,
	}

	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Printf("Erreur lors de la sérialisation en JSON : %v\n", err)
		return
	}

	if mkdirErr := os.MkdirAll("report", fileutils.DirPermission); mkdirErr != nil {
		log.Panicf("cannot create report directory ( %s )\n", mkdirErr)
	}

	file, reportCreationErr := os.Create("report/report.json")
	if reportCreationErr != nil {
		log.Panicf("cannot create reporters file in this folder ( %s )\n", reportCreationErr)
	}
	defer file.Close()

	_, jsonWriteErr := file.Write(jsonData)
	if jsonWriteErr != nil {
		log.Panicf("error when reporters filling ( %s )", jsonWriteErr)
	}
}
