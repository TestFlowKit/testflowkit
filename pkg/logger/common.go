package logger

import (
	"fmt"
	coreLogger "log"
	"strings"

	"github.com/fatih/color"
)

func log(level logLevel, message string) {
	logColor := getLevelColor(level)
	coreLogger.Printf("[%s] %s\n", logColor(string(level)), logColor(message))
}

func getLevelColor(level logLevel) func(format string, a ...interface{}) string {
	switch level {
	case info:
		return color.BlueString
	case success:
		return color.GreenString
	case warn:
		return color.YellowString
	case erro:
		return color.RedString
	case fatal:
		return color.HiRedString
	default:
		return color.WhiteString
	}
}

type logLevel string

const (
	info    logLevel = "INFO"
	warn    logLevel = "WARN"
	success logLevel = "SUCCESS"
	erro    logLevel = "ERROR"
	fatal   logLevel = "FATAL"
)

const indent = "  "

func GetIndents(number int) string {
	idents := ""
	for range number {
		idents += indent
	}
	return idents
}

func formatList(elements []string) string {
	for i, element := range elements {
		elements[i] = fmt.Sprintf("%s- %s", indent, element)
	}

	return strings.Join(elements, "\n")
}
