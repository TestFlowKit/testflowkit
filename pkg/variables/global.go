package variables

import (
	"maps"
	"sync"
)

var (
	globalStore = make(map[string]any)
	mutex       sync.RWMutex
)

func SetGlobalVariable(name string, value any) {
	mutex.Lock()
	defer mutex.Unlock()
	globalStore[name] = value
}

func GetGlobalVariables() map[string]any {
	mutex.RLock()
	defer mutex.RUnlock()

	vars := make(map[string]any)
	maps.Copy(vars, globalStore)
	return vars
}

func ResetGlobalVariables() {
	mutex.Lock()
	defer mutex.Unlock()
	globalStore = make(map[string]any)
}
