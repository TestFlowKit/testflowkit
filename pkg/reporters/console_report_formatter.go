package reporters

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

const consoleSeparator = "────────────────────────────────────────────────────────────────"

type consoleReportFormatter struct{}

func (f consoleReportFormatter) WriteReport(details testSuiteDetails) {
	green := color.New(color.FgGreen).SprintfFunc()
	red := color.New(color.FgRed).SprintfFunc()
	bold := color.New(color.Bold).SprintfFunc()
	cyan := color.New(color.FgCyan).SprintfFunc()
	yellow := color.New(color.FgYellow).SprintfFunc()

	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout, cyan(consoleSeparator))
	fmt.Fprintf(os.Stdout, "  %s  │  %s  │  %s\n",
		bold("TEST REPORT"),
		details.ExecutionDate,
		details.TotalExecutionTime,
	)
	fmt.Fprintln(os.Stdout, cyan(consoleSeparator))
	fmt.Fprintf(os.Stdout, "  Total: %s  │  Passed: %s  │  Failed: %s  │  Rate: %s%%\n",
		bold(details.TotalTests),
		green(details.SucceededTests),
		colorizeFailCount(details.FailedTests, red),
		colorizeRate(details.SuccessRate, green, red, yellow),
	)
	fmt.Fprintln(os.Stdout, cyan(consoleSeparator))

	// Separate scenarios into categories
	beforeAllScenarios := filterScenarios(details.Scenarios, beforeHook)
	afterAllScenarios := filterScenarios(details.Scenarios, afterHook)
	regularScenarios := filterScenarios(details.Scenarios, scenario)

	// Print BeforeAll hooks
	if len(beforeAllScenarios) > 0 {
		fmt.Fprintln(os.Stdout, cyan("📋 BEFORE ALL HOOKS"))
		fmt.Fprintln(os.Stdout, cyan(consoleSeparator))
		printScenariosGrouped(beforeAllScenarios, green, red)
	}

	// Print regular scenarios
	if len(regularScenarios) > 0 {
		fmt.Fprintln(os.Stdout, cyan("📋 SCENARIOS"))
		fmt.Fprintln(os.Stdout, cyan(consoleSeparator))
		printScenariosGrouped(regularScenarios, green, red)
	}

	// Print AfterAll hooks
	if len(afterAllScenarios) > 0 {
		fmt.Fprintln(os.Stdout, cyan("📋 AFTER ALL HOOKS"))
		fmt.Fprintln(os.Stdout, cyan(consoleSeparator))
		printScenariosGrouped(afterAllScenarios, green, red)
	}

	fmt.Fprintln(os.Stdout, cyan(consoleSeparator))
	fmt.Fprintln(os.Stdout)
}

func printScenario(sc Scenario, green, red func(format string, a ...interface{}) string) {
	duration := fmt.Sprintf("%.2fs", sc.Duration.Seconds())

	if sc.Result == succeeded {
		fmt.Fprintf(os.Stdout, "  %s  %s (%s)\n",
			green("✓"),
			sc.Title,
			duration,
		)
		return
	}

	fmt.Fprintf(os.Stdout, "  %s  %s (%s)\n",
		red("✗"),
		sc.Title,
		duration,
	)

	if sc.ErrorMsg != "" {
		fmt.Fprintf(os.Stdout, "       %s %s\n",
			red("Error:"),
			sc.ErrorMsg,
		)
	}

	failedStep := findFailedStep(sc.Steps)
	if failedStep != "" {
		fmt.Fprintf(os.Stdout, "       %s %s\n",
			red("Failed step:"),
			failedStep,
		)
	}
}

func findFailedStep(steps []Step) string {
	for _, step := range steps {
		if strings.EqualFold(step.Status, "failed") {
			return step.Title
		}
	}
	return ""
}

func colorizeFailCount(count string, red func(format string, a ...interface{}) string) string {
	if count == "0" {
		return count
	}
	return red(count)
}

func colorizeRate(rate string, green, red, yellow func(format string, a ...interface{}) string) string {
	switch rate {
	case "100":
		return green(rate)
	case "0":
		return red(rate)
	default:
		return yellow(rate)
	}
}

// filterScenarios returns scenarios matching the given type.
func filterScenarios(scenarios []Scenario, scType scenarioType) []Scenario {
	var filtered []Scenario
	for _, sc := range scenarios {
		if sc.Type == scType {
			filtered = append(filtered, sc)
		}
	}
	return filtered
}

// printScenariosGrouped prints scenarios grouped by success/failure (succeeded first, then failed).
func printScenariosGrouped(scenarios []Scenario, green, red func(format string, a ...interface{}) string) {
	// Separate succeeded and failed scenarios
	var succeeded, failed []Scenario
	for _, sc := range scenarios {
		if sc.Result == scenarioResult("succeeded") {
			succeeded = append(succeeded, sc)
		} else {
			failed = append(failed, sc)
		}
	}

	// Print succeeded scenarios first
	for _, sc := range succeeded {
		printScenario(sc, green, red)
	}

	// Print failed scenarios
	for _, sc := range failed {
		printScenario(sc, green, red)
	}
}
