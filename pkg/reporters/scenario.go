package reporters

import (
	"fmt"
	"math"
	"time"

	"github.com/cucumber/godog"
)

type Scenario struct {
	Title                string
	Steps                []Step
	ErrorMsg             string
	StartDate            time.Time
	Duration             time.Duration
	FmtDuration          string
	HTMLStatusColorClass string
	Result               scenarioResult
}

func (s *Scenario) AddStep(title string, status godog.StepResultStatus, duration time.Duration, err error) {
	if err != nil {
		s.ErrorMsg = err.Error()
	}

	getColor := func(status godog.StepResultStatus) string {
		stepStatusColors := map[godog.StepResultStatus]string{
			godog.StepPassed: "green",
			godog.StepFailed: "red",
		}

		if color, ok := stepStatusColors[status]; ok {
			return color
		}
		return "gray"
	}

	s.Steps = append(s.Steps, Step{
		Title:                title,
		Status:               status.String(),
		HTMLStatusColorClass: fmt.Sprintf("text-%s-500", getColor(status)),
		Duration:             duration,
		FmtDuration:          fmt.Sprintf("%dms", duration.Milliseconds()),
	})
}

func (s *Scenario) SetTitle(title string) {
	s.Title = title
}

func (s *Scenario) End() {
	duration := time.Since(s.StartDate)

	durationInS := duration.Seconds()
	durationInS = math.Max(math.Ceil(durationInS), durationInS)

	result, color, err := failed, "red", s.ErrorMsg
	if len(err) == 0 {
		result, color, err = succeeded, "green", "-"
	}

	s.ErrorMsg = err
	s.Duration = duration
	s.FmtDuration = fmt.Sprintf("%vs", durationInS)
	s.Result = result
	s.HTMLStatusColorClass = fmt.Sprintf("bg-%s-500", color)
}

type Step struct {
	Title                string
	Status               string
	HTMLStatusColorClass string
	Duration             time.Duration
	FmtDuration          string
}

func NewScenario() Scenario {
	return Scenario{
		StartDate: time.Now(),
	}
}
