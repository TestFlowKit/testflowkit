package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

var content string

type FrontConfig struct{}

func (c FrontConfig) init(filePath string) {
	if len(content) != 0 {
		return
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("frontend config file not found")
	}
	content = string(file)
}

func (c FrontConfig) GetPageURL(page string) (string, error) {
	page = c.wildcardToKey(page)
	var pageURL string
	path, err := yaml.PathString(fmt.Sprintf("$.global.pages.%s", page))
	if err != nil {
		return "", err
	}

	err = path.Read(strings.NewReader(content), &pageURL)
	return pageURL, err
}

func (c FrontConfig) GetHTMLElementSelectors(name string) ([]string, error) {
	var selectors []string

	name = c.wildcardToKey(name)

	path, err := yaml.PathString(fmt.Sprintf("$.global.elements.%s", name))
	if err == nil {
		err = path.Read(strings.NewReader(content), &selectors)
	}

	if len(selectors) == 0 {
		return selectors, fmt.Errorf("no selectors found for %s", name)
	}
	return selectors, err
}

func (c FrontConfig) wildcardToKey(label string) string {
	return strings.ToLower(strings.ReplaceAll(label, " ", "_"))
}
