package config

import (
	"os"
	"slices"
	"testing"
)

func setup(t *testing.T) {
	bytes, fileErr := os.ReadFile("frontend.test.yml")
	if fileErr != nil {
		t.Error(fileErr)
	}
	yamlContent = string(bytes)
}

func TestGetPageUrl(t *testing.T) {
	setup(t)

	url, err := GetPageURL("home")
	if err != nil {
		t.Error(err)
	}

	expected := "https://google.com"
	if url != expected {
		t.Errorf("url got %s / url expected: %s", url, expected)
	}
}

func TestGetElementSelectors(t *testing.T) {
	setup(t)

	expected := []string{"#login", ".login", "button .login"}
	if selectors, _ := GetElementSelectors("login"); !slices.Equal(selectors, expected) {
		t.Errorf("error ")
	}
}

func TestGetInputSelectors(t *testing.T) {
	setup(t)

	expected := []string{"#username"}

	if selectors, _ := GetInputSelectors("username"); !slices.Equal(selectors, expected) {
		t.Errorf("error ")
	}
}
