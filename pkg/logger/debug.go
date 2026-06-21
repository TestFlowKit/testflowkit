package logger

import "fmt"

// Debug logs a message at DEBUG level (cyan). Messages are only meaningful
// when debug mode is enabled in the config; callers are responsible for
// gating calls behind the config flag to avoid the formatting overhead.
func Debug(message string) {
	log(debug, message)
}

// Debugf logs a formatted message at DEBUG level.
func Debugf(format string, args ...interface{}) {
	log(debug, fmt.Sprintf(format, args...))
}
