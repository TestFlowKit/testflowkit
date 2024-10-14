package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

var yamlContent string

func InitializeFrontConfig() {
	file, err := os.ReadFile("frontend.yml")
	if err != nil {
		log.Fatal("frontend config file not found")
	}
	yamlContent = string(file)
}

func GetPageURL(page string) (string, error) {
	page = wildcardToKey(page)
	var pageURL string
	path, err := yaml.PathString(fmt.Sprintf("$.global.pages.%s", page))
	if err != nil {
		return "", err
	}

	err = path.Read(strings.NewReader(yamlContent), &pageURL)
	return pageURL, err
}

func GetElementSelectors(label string) ([]string, error) {
	var selectors []string
	label = wildcardToKey(label)

	path, err := yaml.PathString(fmt.Sprintf("$.global.elements.%s", label))
	if err == nil {
		err = path.Read(strings.NewReader(yamlContent), &selectors)
	}

	return selectors, err
}

func GetInputSelectors(name string) ([]string, error) {
	var selectors []string

	name = wildcardToKey(name)

	path, err := yaml.PathString(fmt.Sprintf("$.global.inputs.%s", name))
	if err == nil {
		err = path.Read(strings.NewReader(yamlContent), &selectors)
	}
	return selectors, err
}

func wildcardToKey(label string) string {
	return strings.ToLower(strings.ReplaceAll(label, " ", "_"))
}
