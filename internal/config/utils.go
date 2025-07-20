package config

import (
	"strings"
)

func IsElementDefined(elementName string) bool {
	for _, pageElements := range cfg.Frontend.Elements {
		if _, ok := pageElements[elementName]; ok {
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

func GetLabelKey(label string) string {
	return strings.ToLower(strings.ReplaceAll(label, " ", "_"))
}
