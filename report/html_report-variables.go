package report

type reportVariables struct {
	executionDate, totalTests, succeededTests, failedTests, successRate, scenariosTemplate string
}

func getReportVariables() reportVariables {
	return reportVariables{
		executionDate:     "EXECUTION_DATE",
		totalTests:        "TOTAL_TESTS",
		succeededTests:    "SUCCEEDED_TESTS",
		failedTests:       "FAILED_TESTS",
		successRate:       "SUCCESS_RATE",
		scenariosTemplate: "SCENARIOS",
	}
}

type scenarioVariables struct {
	name, duration, result, statusColor, errorMessage, stepsTemplate string
}

func getScenarioVariables() scenarioVariables {
	return scenarioVariables{
		name:          "SCENARIO_NAME",
		duration:      "SCENARIO_DURATION",
		result:        "SCENARIO_RESULT",
		statusColor:   "SCENARIO_STATUS_COLOR",
		errorMessage:  "SCENARIO_ERROR_MESSAGE",
		stepsTemplate: "STEPS",
	}
}

type stepVariables struct {
	title, status, statusColor string
}

func getStepVariables() stepVariables {
	return stepVariables{
		title:       "STEP_TITLE",
		status:      "STEP_STATUS",
		statusColor: "STEP_STATUS_COLOR",
	}
}
