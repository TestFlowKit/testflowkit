package logger

import (
	"fmt"
	coreLogger "log"
	"strings"
	"time"

	"github.com/fatih/color"
)

func log(level logLevel, message string) {
	format := getLogFormat()
	fileWriter := getLogFile()

	if format == LogFormatJSON {
		writeJSONLog(level, message, fileWriter)
		return
	}

	logColor := getLevelColor(level)
	coreLogger.Printf("[%s] %s\n", logColor(string(level)), logColor(message))

	if fileWriter != nil {
		ts := time.Now().Format("2006/01/02 15:04:05")
		fmt.Fprintf(fileWriter, "%s [%s] %s\n", ts, string(level), message)
	}
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
		return color.RedString
	case debug:
		return color.CyanString
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
	debug   logLevel = "DEBUG"
)

const indent = "  "

func GetIndents(number int) string {
	var identsSb47 strings.Builder
	for range number {
		identsSb47.WriteString(indent)
	}
	return identsSb47.String()
}

func formatList(elements []string) string {
	for i, element := range elements {
		elements[i] = fmt.Sprintf("%s- %s", indent, element)
	}

	return strings.Join(elements, "\n")
}
