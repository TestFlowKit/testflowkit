package logger

import "fmt"

// Debug logs at verbosity ≥ 2 (enabled by --debug or --verbosity=2+).
func Debug(message string) {
	if GetVerbosity() >= VerbosityDefault {
		log(debug, message)
	}
}

// Debugf logs a formatted message at verbosity ≥ 2.
func Debugf(format string, args ...interface{}) {
	if GetVerbosity() >= VerbosityDefault {
		log(debug, fmt.Sprintf(format, args...))
	}
}

// DebugInScope logs at verbosity ≥ 2 only when the given debug scope is active.
func DebugInScope(scope, message string) {
	if GetVerbosity() >= VerbosityDefault && IsScopeEnabled(scope) {
		log(debug, message)
	}
}

// DebugInScopef logs a formatted message at verbosity ≥ 2 only when the given scope is active.
func DebugInScopef(scope, format string, args ...interface{}) {
	if GetVerbosity() >= VerbosityDefault && IsScopeEnabled(scope) {
		log(debug, fmt.Sprintf(format, args...))
	}
}

// Trace logs at verbosity ≥ 3 (--verbosity=3). Use for full bodies and internal state.
func Trace(message string) {
	if GetVerbosity() >= VerbosityTrace {
		log(debug, message)
	}
}

// Tracef logs a formatted message at verbosity ≥ 3.
func Tracef(format string, args ...interface{}) {
	if GetVerbosity() >= VerbosityTrace {
		log(debug, fmt.Sprintf(format, args...))
	}
}
