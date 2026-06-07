package stepexpr

import (
	"fmt"
	"regexp"
	"testing"
)

func TestShouldReplaceTheOnlyStringWildcard(t *testing.T) {
	expected := fmt.Sprintf("^I am redirected to %s page$", stringWildcard)
	result := ConvertWildcards("^I am redirected to {string} page$")

	if result != expected {
		t.Fatalf(`"%s" expected but "%s" received`, expected, result)
	}
}

func TestShouldReplaceTheOnlyNumberWildcard(t *testing.T) {
	expected := fmt.Sprintf("^I must see %s links", numberWildcard)
	result := ConvertWildcards("^I must see {number} links")

	if result != expected {
		t.Fatalf(`"%s" expected but "%s" received`, expected, result)
	}
}

func TestShouldReplaceManyWildcard(t *testing.T) {
	expected := fmt.Sprintf("^I must see %s links which contains %s", numberWildcard, stringWildcard)
	result := ConvertWildcards("^I must see {number} links which contains {string}")

	if result != expected {
		t.Fatalf(`"%s" expected but "%s" received`, expected, result)
	}
}

func TestWildcardShouldMatchEscapedQuotes(t *testing.T) {
	pattern := ConvertWildcards(`^the response field "{string}" should contain "{string}"$`)
	re := regexp.MustCompile(pattern)

	step := `the response field "errors.0.message" should contain "` +
		`Field \"priority\" of required type \"String!\" was not provided."`
	if !re.MatchString(step) {
		t.Fatalf("pattern did not match step with escaped quotes\npattern: %s\nstep: %s", pattern, step)
	}
}

func TestNumberWildcardShouldRejectText(t *testing.T) {
	pattern := ConvertWildcards(`^I must see {number} links$`)
	re := regexp.MustCompile(pattern)

	if re.MatchString(`I must see many links`) {
		t.Fatalf("number wildcard unexpectedly matched non-numeric text\npattern: %s", pattern)
	}
}

func TestNumberWildcardShouldMatchNumericValues(t *testing.T) {
	pattern := ConvertWildcards(`^I must see {number} links$`)
	re := regexp.MustCompile(pattern)

	cases := []string{
		`I must see 10 links`,
		`I must see -3 links`,
		`I must see 4.5 links`,
		`I must see "12" links`,
	}

	for _, step := range cases {
		if !re.MatchString(step) {
			t.Fatalf("number wildcard did not match numeric step\npattern: %s\nstep: %s", pattern, step)
		}
	}
}

func TestIntWildcardShouldRejectNonIntegerValues(t *testing.T) {
	pattern := ConvertWildcards(`^I must see {int} links$`)
	re := regexp.MustCompile(pattern)

	if re.MatchString(`I must see 4.5 links`) {
		t.Fatalf("int wildcard unexpectedly matched a float value\npattern: %s", pattern)
	}
	if re.MatchString(`I must see "12" links`) {
		t.Fatalf("int wildcard unexpectedly matched a quoted number\npattern: %s", pattern)
	}
	if !re.MatchString(`I must see 10 links`) {
		t.Fatalf("int wildcard did not match an integer value\npattern: %s", pattern)
	}
	if !re.MatchString(`I must see -3 links`) {
		t.Fatalf("int wildcard did not match a negative integer value\npattern: %s", pattern)
	}
}

func TestIntWildcardShouldMatchIntegerValues(t *testing.T) {
	pattern := ConvertWildcards(`^I must see {int} links$`)
	re := regexp.MustCompile(pattern)

	cases := []string{
		`I must see 10 links`,
		`I must see -3 links`,
	}

	for _, step := range cases {
		if !re.MatchString(step) {
			t.Fatalf("int wildcard did not match integer step\npattern: %s\nstep: %s", pattern, step)
		}
	}
}
