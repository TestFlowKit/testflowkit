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

	for _, sc := range details.Scenarios {
		printScenario(sc, green, red)
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
