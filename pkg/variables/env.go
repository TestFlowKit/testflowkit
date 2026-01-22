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

var envVarRe = regexp.MustCompile(`\{\{\s*env\.([\w\.-]+)\s*\}\}`)

func SetEnvVariables(vars map[string]string) {
	envLock.Lock()
	defer envLock.Unlock()
	envStore = make(map[string]string)
	if vars == nil {
		return
	}
	maps.Copy(envStore, vars)
}

func GetEnvVariable(key string) (string, bool) {
	envLock.RLock()
	defer envLock.RUnlock()
	val, ok := envStore[key]
	return val, ok
}

func GetAllEnvVariables() map[string]string {
	envLock.RLock()
	defer envLock.RUnlock()
	vars := make(map[string]string)
	maps.Copy(vars, envStore)
	return vars
}

func FindUndefinedEnvReferences(content string) []string {
	if content == "" {
		return []string{}
	}
	matches := envVarRe.FindAllStringSubmatch(content, -1)

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

func ReplaceEnvVariables(content string) string {
	if content == "" {
		return content
	}
	return envVarRe.ReplaceAllStringFunc(content, func(match string) string {
		submatch := envVarRe.FindStringSubmatch(match)
		if len(submatch) > 1 {
			key := submatch[1]
			if val, exists := GetEnvVariable(key); exists {
				return val
			}
		}
		return match
	})
}

func ResetEnvVariables() {
	envLock.Lock()
	defer envLock.Unlock()
	envStore = make(map[string]string)
}
