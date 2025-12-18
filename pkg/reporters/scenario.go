package reporters

import (
	"fmt"
	"time"

	"github.com/cucumber/godog"
)

type Scenario struct {
	Title     string
	Steps     []Step
	ErrorMsg  string
	StartDate time.Time
	Duration  time.Duration
	Result    scenarioResult
}

func (s *Scenario) AddStep(title string, status godog.StepResultStatus, duration time.Duration, err error) {
	if err != nil {
		s.ErrorMsg = err.Error()
	}

	getColor := func(status godog.StepResultStatus) string {
		stepStatusColors := map[godog.StepResultStatus]string{
			godog.StepPassed:    "green",
			godog.StepFailed:    "red",
			godog.StepSkipped:   "yellow",
			godog.StepPending:   "gray",
			godog.StepUndefined: "gray",
			godog.StepAmbiguous: "gray",
		}

		if color, ok := stepStatusColors[status]; ok {
			return color
		}
		return "gray"
	}

	screenshotBase64 := ""
	if stepErr, ok := err.(interface{ ScreenshotBase64() string }); ok {
		screenshotBase64 = stepErr.ScreenshotBase64()
	}

	s.Steps = append(s.Steps, Step{
		Title:                title,
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
	Status               string
	HTMLStatusColorClass string
	Duration             time.Duration
	FmtDuration          string
	ScreenshotBase64     string
}

func NewScenario() Scenario {
	return Scenario{
		StartDate: time.Now(),
	}
}
