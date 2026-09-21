package reporters

import (
	"fmt"
	"time"

	"github.com/cucumber/godog"
)

const colorGray = "gray"

type scenarioType string

const (
	beforeHook scenarioType = "beforeHook"
	scenario   scenarioType = "scenario"
	afterHook  scenarioType = "afterHook"
)

type Scenario struct {
	Title     string
	ID        string
	URI       string
	Tags      []string
	Steps     []Step
	ErrorMsg  string
	StartDate time.Time
	Duration  time.Duration
	Result    scenarioResult
	Type      scenarioType
}

func (s *Scenario) AddStep(title, astNodeID string, status godog.StepResultStatus, duration time.Duration, err error) {
	if err != nil {
		s.ErrorMsg = err.Error()
	}

	getColor := func(status godog.StepResultStatus) string {
		stepStatusColors := map[godog.StepResultStatus]string{
			godog.StepPassed:    "green",
			godog.StepFailed:    "red",
			godog.StepSkipped:   "yellow",
			godog.StepPending:   colorGray,
			godog.StepUndefined: colorGray,
			godog.StepAmbiguous: colorGray,
		}

		if color, ok := stepStatusColors[status]; ok {
			return color
		}
		return colorGray
	}

	screenshotBase64 := ""
	if stepErr, ok := err.(interface{ ScreenshotBase64() string }); ok {
		screenshotBase64 = stepErr.ScreenshotBase64()
	}

	s.Steps = append(s.Steps, Step{
		Title:                title,
		astNodeID:            astNodeID,
		Status:               status.String(),
		HTMLStatusColorClass: fmt.Sprintf("text-%s-500", getColor(status)),
		Duration:             duration,
		FmtDuration:          fmt.Sprintf("%dms", duration.Milliseconds()),
		ScreenshotBase64:     screenshotBase64,
	})
}

func (s *Scenario) SetTitle(title string) {
	s.Title = title
}

// SetMetadata stores the feature and identification data used by structured reports.
func (s *Scenario) SetMetadata(id, uri string, tags []string) {
	s.ID = id
	s.URI = uri
	s.Tags = tags
}

// ResolveKeywords sets each step's Gherkin keyword from a map of step AST node id to keyword.
func (s *Scenario) ResolveKeywords(keywordsByID map[string]string) {
	for i := range s.Steps {
		s.Steps[i].Keyword = keywordsByID[s.Steps[i].astNodeID]
	}
}

func (s *Scenario) End() {
	duration := time.Since(s.StartDate)

	result, err := failed, s.ErrorMsg
	if len(err) == 0 {
		result, err = succeeded, ""
	}

	s.ErrorMsg = err
	s.Duration = duration
	s.Result = result
}

type Step struct {
	Title                string
	Keyword              string
	astNodeID            string
	Status               string
	HTMLStatusColorClass string
	Duration             time.Duration
	FmtDuration          string
	ScreenshotBase64     string
}

func NewMainScenario() Scenario {
	return Scenario{
		StartDate: time.Now(),
		Type:      scenario,
	}
}

func NewBeforeAllHook() Scenario {
	return Scenario{
		StartDate: time.Now(),
		Type:      beforeHook,
	}
}

func NewAfterAllHook() Scenario {
	return Scenario{
		StartDate: time.Now(),
		Type:      afterHook,
	}
}
