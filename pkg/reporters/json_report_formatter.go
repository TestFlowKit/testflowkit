package reporters

import (
	"encoding/json"
	"log"
)

const reportJSONPath = "report/report.json"

type jsonReportFormatter struct{}

func (f jsonReportFormatter) WriteReport(details testSuiteDetails) {
	order, groups := groupByFeature(details.Scenarios)

	features := make([]cucumberFeature, 0, len(order))
	for _, uri := range order {
		feature := cucumberFeature{
			URI:      uri,
			ID:       uri,
			Keyword:  "Feature",
			Name:     uri,
			Elements: make([]cucumberElement, 0, len(groups[uri])),
		}
		for _, sc := range groups[uri] {
			feature.Elements = append(feature.Elements, toCucumberElement(sc))
		}
		features = append(features, feature)
	}

	jsonData, err := json.MarshalIndent(features, "", "  ")
	if err != nil {
		log.Printf("error while serializing the JSON report: %v\n", err)
		return
	}

	writeReportFile(reportJSONPath, jsonData)
}

func toCucumberElement(sc Scenario) cucumberElement {
	element := cucumberElement{
		ID:      sc.ID,
		Keyword: "Scenario",
		Type:    "scenario",
		Name:    sc.Title,
		Steps:   make([]cucumberStep, len(sc.Steps)),
	}
	for _, tag := range sc.Tags {
		element.Tags = append(element.Tags, cucumberTag{Name: tag})
	}

	for i, step := range sc.Steps {
		cs := cucumberStep{
			Keyword: step.Keyword,
			Name:    step.Title,
			Result: cucumberResult{
				Status:   step.Status,
				Duration: step.Duration.Nanoseconds(),
			},
		}
		if step.Status == stepStatusFailed {
			cs.Result.ErrorMessage = sc.ErrorMsg
		}
		if step.ScreenshotBase64 != "" {
			cs.Embeddings = []cucumberEmbedding{{MimeType: "image/png", Data: step.ScreenshotBase64}}
		}
		element.Steps[i] = cs
	}
	return element
}
