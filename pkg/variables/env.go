package variables

import (
	"maps"
	"regexp"
	"sort"
	"sync"
)

var (
	envStore = make(map[string]string)
	envLock  sync.RWMutex
)

// SetEnvVariables sets the environment variables store.
// This replaces all existing environment variables with the provided map.
// Thread-safe for concurrent access.
func SetEnvVariables(vars map[string]string) {
	envLock.Lock()
	defer envLock.Unlock()
	envStore = make(map[string]string)
	if vars == nil {
		return
	}
	maps.Copy(envStore, vars)
}

// GetEnvVariable retrieves an environment variable by key.
// Returns the value and true if found, empty string and false otherwise.
// Thread-safe for concurrent access.
func GetEnvVariable(key string) (string, bool) {
	envLock.RLock()
	defer envLock.RUnlock()
	val, ok := envStore[key]
	return val, ok
}

// GetAllEnvVariables returns a copy of all environment variables.
// Returns a new map to prevent external modification of the internal store.
// Thread-safe for concurrent access.
func GetAllEnvVariables() map[string]string {
	envLock.RLock()
	defer envLock.RUnlock()
	vars := make(map[string]string)
	maps.Copy(vars, envStore)
	return vars
}

// FindUndefinedEnvReferences finds all {{ env.key }} references in content that are not defined in the store.
// Returns a sorted list of unique undefined keys.
// Thread-safe for concurrent access.
func FindUndefinedEnvReferences(content string) []string {
	if content == "" {
		return []string{}
	}
	re := regexp.MustCompile(`\{\{\s*env\.([\w\.-]+)\s*\}\}`)
	matches := re.FindAllStringSubmatch(content, -1)

	undefinedMap := make(map[string]bool)

	envLock.RLock()
	defer envLock.RUnlock()

	for _, match := range matches {
		if len(match) > 1 {
			key := match[1]
			if _, exists := envStore[key]; !exists {
				undefinedMap[key] = true
			}
		}
	}

	undefined := make([]string, 0, len(undefinedMap))
	for k := range undefinedMap {
		undefined = append(undefined, k)
	}
	sort.Strings(undefined)

	return undefined
}

// ReplaceEnvVariables finds all {{ env.key }} references in content and replaces them with their values.
// If a variable is not found, the placeholder is left unchanged.
// Thread-safe for concurrent access.
func ReplaceEnvVariables(content string) string {
	if content == "" {
		return content
	}
	re := regexp.MustCompile(`\{\{\s*env\.([\w\.-]+)\s*\}\}`)
	return re.ReplaceAllStringFunc(content, func(match string) string {
		submatch := re.FindStringSubmatch(match)
		if len(submatch) > 1 {
			key := submatch[1]
			if val, exists := GetEnvVariable(key); exists {
				return val
			}
		}
		return match
	})
}

// ResetEnvVariables clears all environment variables from the store.
// Useful for testing purposes to ensure clean state between tests.
func ResetEnvVariables() {
	envLock.Lock()
	defer envLock.Unlock()
	envStore = make(map[string]string)
}
