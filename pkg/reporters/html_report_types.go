package reporters

import "fmt"

type htmlTestSuiteDetails struct {
	testSuiteDetails
	Scenarios []htmlScenario
}

type htmlScenario struct {
	Scenario
	FmtDuration          string
	HTMLStatusColorClass string
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
