package testsConfig

import (
	"os"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup(_ *testing.T) {
	file, err := os.ReadFile("frontend.test.yml")
	if err != nil {
		panic(err)
	}

	content = string(file)
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
	if selectors, _ := GetHTMLElementSelectors("login"); !slices.Equal(selectors, expected) {
		t.Errorf("error ")
	}
}

func TestGetBaseUrl(t *testing.T) {
	setup(t)

	expected := "http://etoole.test"
	assert.Equal(t, expected, getBaseURL())
}

func TestFFormatPageUrlWithPath(t *testing.T) {
	setup(t)
	pageURL, _ := formatPageURL("/home")
	assert.Equal(t, "http://etoole.test/home", pageURL)
}

func TestFFormatPageUrlWithUrl(t *testing.T) {
	setup(t)
	pageURL, _ := formatPageURL("http://home.com/home")
	assert.Equal(t, "http://home.com/home", pageURL)
}
