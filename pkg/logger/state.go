package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// Verbosity levels for --debug / --verbosity flags.
const (
	VerbosityOff     = 0 // no debug output
	VerbositySummary = 1 // scenario/step flow with timings
	VerbosityDefault = 2 // + request/response headers and variable substitutions (--debug default)
	VerbosityTrace   = 3 // + full bodies and all internal state (--verbosity=3)
)

// Debug scope names for --debug-scope.
const (
	ScopeHTTP      = "http"
	ScopeBrowser   = "browser"
	ScopeVariables = "variables"
	ScopeConfig    = "config"
)

// Log format names for --log-format.
const (
	LogFormatText = "text"
	LogFormatJSON = "json"
)

type loggerState struct {
	verbosity int
	scopes    map[string]struct{}
	allScopes bool
	logFormat string
	logFile   io.WriteCloser
}

var (
	mu      sync.RWMutex
	current = loggerState{verbosity: VerbosityOff, allScopes: true, logFormat: LogFormatText}
)

// Init configures the logger for debug output. Call once after the config is loaded.
func Init(verbosity int, scopes, logFile, logFormat string) error {
	mu.Lock()
	defer mu.Unlock()

	current.verbosity = verbosity

	current.logFormat = LogFormatText
	if logFormat == LogFormatJSON {
		current.logFormat = LogFormatJSON
	}

	if scopes == "" {
		current.allScopes = true
		current.scopes = nil
	} else {
		current.allScopes = false
		current.scopes = make(map[string]struct{})
		for _, s := range strings.Split(scopes, ",") {
			if s = strings.TrimSpace(strings.ToLower(s)); s != "" {
				current.scopes[s] = struct{}{}
			}
		}
	}

	if current.logFile != nil {
		_ = current.logFile.Close()
		current.logFile = nil
	}
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			return fmt.Errorf("failed to open log file %q: %w", logFile, err)
		}
		current.logFile = f
	}

	return nil
}

// IsDebugEnabled reports whether any debug output is currently enabled.
func IsDebugEnabled() bool {
	mu.RLock()
	defer mu.RUnlock()
	return current.verbosity > VerbosityOff
}

// GetVerbosity returns the current verbosity level (0–3).
func GetVerbosity() int {
	mu.RLock()
	defer mu.RUnlock()
	return current.verbosity
}

// IsScopeEnabled reports whether the given debug scope is active.
func IsScopeEnabled(scope string) bool {
	mu.RLock()
	defer mu.RUnlock()
	if current.allScopes {
		return true
	}
	_, ok := current.scopes[strings.ToLower(scope)]
	return ok
}

func getLogFormat() string {
	mu.RLock()
	defer mu.RUnlock()
	return current.logFormat
}

func getLogFile() io.Writer {
	mu.RLock()
	defer mu.RUnlock()
	return current.logFile
}

type jsonLogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

func writeJSONLog(level logLevel, message string, fileWriter io.Writer) {
	entry := jsonLogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     string(level),
		Message:   message,
	}
	data, _ := json.Marshal(entry)
	line := string(data)
	fmt.Fprintln(os.Stdout, line)
	if fileWriter != nil {
		fmt.Fprintln(fileWriter, line)
	}
}
