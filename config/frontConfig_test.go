package config

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup(_ *testing.T) *FrontConfig {
	config := FrontConfig{}

	config.init("frontend.test.yml")

	return &config
}

func TestGetPageUrl(t *testing.T) {
	config := setup(t)

	url, err := config.GetPageURL("home")
	if err != nil {
		t.Error(err)
	}

	expected := "https://google.com"
	if url != expected {
		t.Errorf("url got %s / url expected: %s", url, expected)
	}
}

func TestGetElementSelectors(t *testing.T) {
	config := setup(t)

	expected := []string{"#login", ".login", "button .login"}
	if selectors, _ := config.GetHTMLElementSelectors("login"); !slices.Equal(selectors, expected) {
		t.Errorf("error ")
	}
}

func TestGetBaseUrl(t *testing.T) {
	config := setup(t)

	expected := "http://etoole.test"
	assert.Equal(t, expected, config.getBaseURL())
}

func TestFFormatPageUrlWithPath(t *testing.T) {
	config := setup(t)
	pageURL, _ := config.formatPageURL("/home")
	assert.Equal(t, "http://etoole.test/home", pageURL)
}

func TestFFormatPageUrlWithUrl(t *testing.T) {
	config := setup(t)
	pageURL, _ := config.formatPageURL("http://home.com/home")
	assert.Equal(t, "http://home.com/home", pageURL)
}
