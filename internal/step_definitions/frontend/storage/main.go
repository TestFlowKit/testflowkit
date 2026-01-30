package storage

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct{}

func GetSteps() []stepbuilder.Step {
	handlers := steps{}

	return []stepbuilder.Step{
		handlers.setLocalStorageItem(),
		handlers.setSessionStorageItem(),
		handlers.storeLocalStorageItemIntoVariable(),
		handlers.storeSessionStorageItemIntoVariable(),
		handlers.deleteLocalStorageItem(),
		handlers.deleteSessionStorageItem(),
		handlers.clearLocalStorage(),
		handlers.clearSessionStorage(),
	}
}
