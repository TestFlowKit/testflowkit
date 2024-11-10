package config

import (
	"cucumber/utils"
	"slices"
	"testing"
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
	if selectors, _ := config.GetHTMLElementSelectors("login", utils.HTMLElement); !slices.Equal(selectors, expected) {
		t.Errorf("error ")
	}
}

func TestGetInputSelectors(t *testing.T) {
	config := setup(t)

	expected := []string{"#username"}

	if selectors, _ := config.GetHTMLElementSelectors("username", utils.HTMLInput); !slices.Equal(selectors, expected) {
		t.Errorf("error ")
	}
}
