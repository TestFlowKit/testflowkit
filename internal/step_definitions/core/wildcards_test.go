package core

import (
	"fmt"
	"testing"
)

func TestShouldReplaceTheOnlyStringWildcard(t *testing.T) {
	expected := fmt.Sprintf("^I am redirected to %s page$", wildcard)
	result := ConvertWildcards("^I am redirected to {string} page$")

	if result != expected {
		t.Fatalf(`"%s" expected but "%s" received`, expected, result)
	}
}

func TestShouldReplaceTheOnlyNumberWildcard(t *testing.T) {
	expected := fmt.Sprintf("^I must see %s links", wildcard)
	result := ConvertWildcards("^I must see {number} links")

	if result != expected {
		t.Fatalf(`"%s" expected but "%s" received`, expected, result)
	}
}

func TestShouldReplaceManyWildcard(t *testing.T) {
	expected := fmt.Sprintf("^I must see %s links which contains %s", wildcard, wildcard)
	result := ConvertWildcards("^I must see {number} links which contains {string}")

	if result != expected {
		t.Fatalf(`"%s" expected but "%s" received`, expected, result)
	}
}
