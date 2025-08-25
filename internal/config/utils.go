package config

import (
	"testflowkit/internal/utils/stringutils"
)

func IsElementDefined(elementName string) bool {
	key := stringutils.SnakeCase(elementName)
	for _, pageElements := range cfg.GetFrontendElements() {
		if _, ok := pageElements[key]; ok {
			return true
		}
	}

	return false
}

func IsFileDefined(fileName string) bool {
	defs := cfg.GetFileDefinitions()
	if defs == nil {
		return false
	}

	_, exists := defs[fileName]
	return exists
}

func IsPageDefined(pageName string) bool {
	pageURL, getFrontendURLErr := cfg.GetFrontendURL(pageName)
	if getFrontendURLErr != nil {
		return false
	}

	return pageURL != ""
}
