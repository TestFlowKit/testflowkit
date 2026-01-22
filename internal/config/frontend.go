package config

import (
	"path/filepath"
	"strings"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/variables"
	"time"
)

func (c *Config) GetElementSelectors(page, elementName string) []Selector {
	elementName = stringutils.SnakeCase(elementName)
	var selectors []Selector

	chainOfResponsibility := []func(string) []Selector{
		c.getElementSelectors(page),
		c.getElementSelectors("common"),
	}

	for _, chain := range chainOfResponsibility {
		selectors = append(selectors, chain(elementName)...)
		if len(selectors) > 0 {
			return selectors
		}
	}

	return []Selector{}
}

func (c *Config) getElementSelectors(key string) func(string) []Selector {
	return func(elementName string) []Selector {
		var selectors []Selector
		if pageElements, isFound := c.GetFrontendElements()[key]; isFound {
			if selectorStrs, isSelectorsFound := pageElements[elementName]; isSelectorsFound {
				for _, selectorStr := range selectorStrs {
					selectors = append(selectors, NewSelector(selectorStr))
				}
				return selectors
			}
		}
		return selectors
	}
}

func (c *Config) IsHeadlessModeEnabled() bool {
	if c.IsFrontendDefined() {
		return c.Frontend.Headless
	}
	return false
}

// GetTimeout returns the element search timeout as a time.Duration.
// This timeout is used when searching for elements by CSS selectors or XPath expressions.
func (c *Config) GetTimeout() time.Duration {
	return c.GetFrontendTimeout()
}

func (c *Config) IsScreenshotOnFailureEnabled() bool {
	if c.IsFrontendDefined() {
		return c.Frontend.ScreenshotOnFailure
	}
	return false
}

func (c *Config) GetFrontendBaseURL() string {
	val, _ := variables.GetEnvVariable("frontend_base_url")
	return val
}

func (c *Config) IsFrontendDefined() bool {
	return c.Frontend != nil
}

func (c *Config) GetFrontendElements() FrontendElements {
	if c.IsFrontendDefined() {
		return c.Frontend.Elements
	}
	return make(FrontendElements)
}

func (c *Config) GetFrontendPages() FrontendPages {
	if c.IsFrontendDefined() {
		return c.Frontend.Pages
	}
	return make(FrontendPages)
}

func (c *Config) GetFrontendTimeout() time.Duration {
	if !c.IsFrontendDefined() {
		return 0
	}
	return time.Duration(c.Frontend.DefaultTimeout) * time.Millisecond
}

func (c *Config) GetFrontendURL(page string) (string, error) {
	if pagePath, ok := c.GetFrontendPages()[stringutils.SnakeCase(page)]; ok {
		if strings.HasPrefix(pagePath, "http://") || strings.HasPrefix(pagePath, "https://") {
			return pagePath, nil
		}

		fullURL := filepath.Join(c.GetFrontendBaseURL(), pagePath)
		return fullURL, nil
	}

	return c.GetFrontendBaseURL(), nil
}
