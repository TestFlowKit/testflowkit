package gherkinparser

import (
	"testing"

	messages "github.com/cucumber/messages/go/v21"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ── helpers ───────────────────────────────────────────────────────────────────

func mustParseFeature(t *testing.T, content string) *Feature {
	t.Helper()
	f, err := parseFeatureContent(content)
	require.NoError(t, err)
	return f
}

// ── collectTagNames ───────────────────────────────────────────────────────────

func Test_CollectTagNames_NilInputs(t *testing.T) {
	assert.Empty(t, collectTagNames())
	assert.Empty(t, collectTagNames(nil))
}

func Test_CollectTagNames_SingleSlice(t *testing.T) {
	tags := []*messages.Tag{{Name: "@smoke"}, {Name: "@critical"}}
	result := collectTagNames(tags)
	assert.Equal(t, []string{"@smoke", "@critical"}, result)
}

func Test_CollectTagNames_MultipleSlicesMerged(t *testing.T) {
	featureTags := []*messages.Tag{{Name: "@featureTag"}}
	scenarioTags := []*messages.Tag{{Name: "@smoke"}}
	result := collectTagNames(featureTags, scenarioTags)
	assert.ElementsMatch(t, []string{"@featureTag", "@smoke"}, result)
}

// ── filterFeatures ────────────────────────────────────────────────────────────

func Test_FilterFeatures_EmptyFilter_ReturnsAllUnchanged(t *testing.T) {
	content := `Feature: Login

  @smoke
  Scenario: Successful login
    Given I am on the login page

  @regression
  Scenario: Failed login
    Given I am on the login page`

	result := filterFeatures([]*Feature{mustParseFeature(t, content)}, "")

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 2)
}

func Test_FilterFeatures_EmptyFeaturesList(t *testing.T) {
	result := filterFeatures([]*Feature{}, "@smoke")
	assert.Empty(t, result)
}

func Test_FilterFeatures_Include_ScenarioLevelTag_OnlyMatchingSurvives(t *testing.T) {
	content := `Feature: Login

  @smoke
  Scenario: Successful login
    Given I am on the login page

  @regression
  Scenario: Failed login
    Given I am on the login page`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"@smoke",
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 1)
	assert.Equal(t, "Successful login", result[0].scenarios[0].Name)
	assert.Contains(t, string(result[0].Contents), "Successful login")
	assert.NotContains(t, string(result[0].Contents), "Failed login")
}

func Test_FilterFeatures_Include_AND_BothTagsRequired(t *testing.T) {
	content := `Feature: Login

  @smoke @critical
  Scenario: Important login
    Given I am on the login page

  @smoke
  Scenario: Less important login
    Given I am on the login page`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"@smoke and @critical",
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 1)
	assert.Equal(t, "Important login", result[0].scenarios[0].Name)
}

func Test_FilterFeatures_Include_FeatureLevelTag_AllScenariosSurvive(t *testing.T) {
	content := `@smoke
Feature: Login

  Scenario: Successful login
    Given I am on the login page

  Scenario: Failed login
    Given I am on the login page`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"@smoke",
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 2)
}

func Test_FilterFeatures_IncludeAndExclude_SingleExpression(t *testing.T) {
	content := `Feature: Login

  @smoke
  Scenario: Successful login
    Given I am on the login page

  @smoke @regression
  Scenario: Flaky login
    Given I am on the login page

  @regression
  Scenario: Failed login
    Given I am on the login page`

	// @smoke and not @regression: must have @smoke AND must NOT have @regression
	// → only Successful login matches
	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"@smoke and not @regression",
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 1)
	assert.Equal(t, "Successful login", result[0].scenarios[0].Name)
}

func Test_FilterFeatures_AllScenariosExcluded_FeatureDropped(t *testing.T) {
	content := `Feature: Login

  @smoke
  Scenario: Successful login
    Given I am on the login page`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"@nonexistent",
	)

	assert.Empty(t, result)
}

func Test_FilterFeatures_MultipleFeatures_OnlyMatchingSurvive(t *testing.T) {
	smokeContent := `Feature: Smoke Tests

  @smoke
  Scenario: Smoke test
    Given I run a smoke test`

	regressionContent := `Feature: Regression Tests

  @regression
  Scenario: Regression test
    Given I run a regression test`

	features := []*Feature{
		mustParseFeature(t, smokeContent),
		mustParseFeature(t, regressionContent),
	}

	result := filterFeatures(features, "@smoke")

	assert.Len(t, result, 1)
	assert.Equal(t, "Smoke Tests", result[0].Name)
}

func Test_FilterFeatures_FeatureHeaderAndBackgroundPreservedAfterFilter(t *testing.T) {
	content := `Feature: Login

  Background:
    Given I open the browser

  @smoke
  Scenario: Successful login
    When I enter valid credentials
    Then I am logged in

  @regression
  Scenario: Failed login
    When I enter invalid credentials
    Then I see an error`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"@smoke",
	)

	assert.Len(t, result, 1)
	resultContent := string(result[0].Contents)
	assert.Contains(t, resultContent, "Feature: Login")
	assert.Contains(t, resultContent, "Background:")
	assert.Contains(t, resultContent, "Successful login")
	assert.NotContains(t, resultContent, "Failed login")
}

func Test_FilterFeatures_RebuiltContentIsValidGherkin(t *testing.T) {
	content := `Feature: API Tests

  @api
  Scenario: Create resource
    Given the API is available
    When I send a POST request
    Then the resource is created

  @ui
  Scenario: View resource
    Given the UI is available
    When I navigate to the resource
    Then I see it`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"@api",
	)

	assert.Len(t, result, 1)
	// Rebuilt content must itself be valid parseable gherkin.
	rebuilt, err := parseFeatureContent(string(result[0].Contents))
	require.NoError(t, err)
	assert.Len(t, rebuilt.scenarios, 1)
	assert.Equal(t, "Create resource", rebuilt.scenarios[0].Name)
}

// ── Filter (public API) ───────────────────────────────────────────────────────

func Test_Filter_EmptyExpression_ReturnsAllFeatures(t *testing.T) {
	content := `Feature: Login

  @smoke
  Scenario: Successful login
    Given I am on the login page`

	features := []*Feature{mustParseFeature(t, content)}
	assert.Equal(t, features, Filter(features, ""))
	assert.Equal(t, features, Filter(features, "   "))
}

// ── Cucumber tag-expression syntax ───────────────────────────────────────────

func Test_FilterFeatures_Cucumber_And(t *testing.T) {
	content := `Feature: Login

  @smoke @critical
  Scenario: Important login
    Given I am on the login page

  @smoke
  Scenario: Less important login
    Given I am on the login page`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"@smoke and @critical",
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 1)
	assert.Equal(t, "Important login", result[0].scenarios[0].Name)
}

func Test_FilterFeatures_Cucumber_Or(t *testing.T) {
	content := `Feature: Login

  @smoke
  Scenario: Successful login
    Given I am on the login page

  @regression
  Scenario: Failed login
    Given I am on the login page`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"@smoke or @regression",
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 2)
}

func Test_FilterFeatures_Cucumber_Not(t *testing.T) {
	content := `Feature: Login

  @smoke
  Scenario: Successful login
    Given I am on the login page

  @regression
  Scenario: Failed login
    Given I am on the login page`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"not @regression",
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 1)
	assert.Equal(t, "Successful login", result[0].scenarios[0].Name)
}

func Test_FilterFeatures_Cucumber_Parentheses(t *testing.T) {
	content := `Feature: Auth

  @login @fast
  Scenario: Login fast
    Given I am on the login page

  @signup @fast
  Scenario: Signup fast
    Given I am on the signup page

  @login @slow
  Scenario: Login slow
    Given I am on the login page slowly`

	// (@login or @signup) and not @slow
	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		"(@login or @signup) and not @slow",
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 2)
	names := []string{result[0].scenarios[0].Name, result[0].scenarios[1].Name}
	assert.Contains(t, names, "Login fast")
	assert.Contains(t, names, "Signup fast")
	assert.NotContains(t, names, "Login slow")
}
