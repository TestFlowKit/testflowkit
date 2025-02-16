package testsconfig

import (
	"errors"
	"etoolse/pkg/logger"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/goccy/go-yaml"
)

var content string

const pagesYamlKey = "$.global.pages"
const elementsYamlKey = "$.global.elements"

func HasAlreadyInitialized() bool {
	return len(content) > 0
}

func Init(fileContent string) {
	if HasAlreadyInitialized() {
		logger.Info("frontend config already initialized")
		return
	}
	content = fileContent
}

func IsElementDefined(elementName string) bool {
	_, err := GetHTMLElementSelectors(elementName)
	return err == nil
}

func IsPageDefined(pageName string) bool {
	_, err := GetPageURL(pageName)
	return err == nil
}

func GetHTMLElementSelectors(name string) ([]string, error) {
	var selectors []string

	name = GetLabelKey(name)

	path, err := yaml.PathString(fmt.Sprintf("%s.%s", elementsYamlKey, name))
	if err == nil {
		err = path.Read(strings.NewReader(content), &selectors)
	}

	if len(selectors) == 0 {
		return selectors, fmt.Errorf("no selectors found for %s", name)
	}
	return selectors, err
}

func GetPageURL(page string) (string, error) {
	page = GetLabelKey(page)

	path, err := yaml.PathString(fmt.Sprintf("%s.%s", pagesYamlKey, page))
	if err != nil {
		return "", err
	}

	var pageURL string
	err = path.Read(strings.NewReader(content), &pageURL)
	if err != nil {
		return "", errors.New("page not found in config")
	}

	return formatPageURL(pageURL)
}

func formatPageURL(pathOrURL string) (string, error) {
	parsedURL, err := url.Parse(pathOrURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Host == "" {
		pageUR, _ := url.JoinPath(getBaseURL(), pathOrURL)
		return pageUR, nil
	}

	return pathOrURL, err
}

func getBaseURL() string {
	path, err := yaml.PathString("$.global.base_url")
	if err != nil {
		log.Panic("base_url path format error")
	}

	var baseURL string
	err = path.Read(strings.NewReader(content), &baseURL)
	if err != nil {
		logger.Info("base_url not found in config")
		return ""
	}

	return baseURL
}

func GetLabelKey(label string) string {
	return strings.ToLower(strings.ReplaceAll(label, " ", "_"))
}
