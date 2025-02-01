package config

import (
	"errors"
	"etoolse/utils"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

var content string

const pagesYamlKey = "$.global.pages"
const elementsYamlKey = "$.global.elements"

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
	page = utils.WildcardToKey(page)

	path, err := yaml.PathString(fmt.Sprintf("%s.%s", pagesYamlKey, page))
	if err != nil {
		return "", err
	}

	var pageURL string
	err = path.Read(strings.NewReader(content), &pageURL)
	if err != nil {
		return "", errors.New("page not found in config")
	}

	return c.formatPageURL(pageURL)
}

func (c FrontConfig) formatPageURL(pathOrURL string) (string, error) {
	parsedURL, err := url.Parse(pathOrURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Host == "" {
		pageUR, _ := url.JoinPath(c.getBaseURL(), pathOrURL)
		return pageUR, nil
	}

	return pathOrURL, err
}

func (c FrontConfig) getBaseURL() string {
	path, err := yaml.PathString("$.global.base_url")
	if err != nil {
		log.Panic("base_url path format error")
	}

	var baseURL string
	err = path.Read(strings.NewReader(content), &baseURL)
	if err != nil {
		log.Println("base_url not found in config")
		return ""
	}

	return baseURL
}

func (c FrontConfig) GetHTMLElementSelectors(name string) ([]string, error) {
	var selectors []string

	name = utils.WildcardToKey(name)

	path, err := yaml.PathString(fmt.Sprintf("%s.%s", elementsYamlKey, name))
	if err == nil {
		err = path.Read(strings.NewReader(content), &selectors)
	}

	if len(selectors) == 0 {
		return selectors, fmt.Errorf("no selectors found for %s", name)
	}
	return selectors, err
}
