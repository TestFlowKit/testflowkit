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

// ── parseTagExpression ────────────────────────────────────────────────────────

func Test_ParseTagExpression_Empty(t *testing.T) {
	assert.Nil(t, parseTagExpression(""))
	assert.Nil(t, parseTagExpression("   "))
}

func Test_ParseTagExpression_SingleTag(t *testing.T) {
	assert.Equal(t, []tagGroup{{"@smoke"}}, parseTagExpression("@smoke"))
	assert.Equal(t, []tagGroup{{"@smoke"}}, parseTagExpression("  @smoke  "))
}

func Test_ParseTagExpression_OR(t *testing.T) {
	assert.Equal(t, []tagGroup{{"@a"}, {"@b"}}, parseTagExpression("@a || @b"))
}

func Test_ParseTagExpression_AND(t *testing.T) {
	assert.Equal(t, []tagGroup{{"@a", "@b"}}, parseTagExpression("@a && @b"))
}

func Test_ParseTagExpression_ORofAND(t *testing.T) {
	result := parseTagExpression("@smoke || @critical && @api")
	assert.Equal(t, []tagGroup{{"@smoke"}, {"@critical", "@api"}}, result)
}

func Test_ParseTagExpression_MultipleANDGroups(t *testing.T) {
	result := parseTagExpression("@a && @b || @c && @d")
	assert.Equal(t, []tagGroup{{"@a", "@b"}, {"@c", "@d"}}, result)
}

func Test_ParseTagExpression_Negation(t *testing.T) {
	assert.Equal(t, []tagGroup{{"~@smoke"}}, parseTagExpression("~@smoke"))
	assert.Equal(t, []tagGroup{{"@smoke", "~@regression"}}, parseTagExpression("@smoke && ~@regression"))
}

// ── matchesFilter ─────────────────────────────────────────────────────────────

func Test_MatchesFilter_EmptyGroups_AlwaysTrue(t *testing.T) {
	assert.True(t, matchesFilter([]string{"@smoke"}, nil))
	assert.True(t, matchesFilter([]string{}, nil))
	assert.True(t, matchesFilter([]string{"@smoke"}, []tagGroup{}))
}

func Test_MatchesFilter_SingleTagGroup_Match(t *testing.T) {
	assert.True(t, matchesFilter([]string{"@smoke", "@regression"}, []tagGroup{{"@smoke"}}))
}

func Test_MatchesFilter_SingleTagGroup_NoMatch(t *testing.T) {
	assert.False(t, matchesFilter([]string{"@regression"}, []tagGroup{{"@smoke"}}))
}

func Test_MatchesFilter_OR_SecondGroupMatches(t *testing.T) {
	groups := []tagGroup{{"@smoke"}, {"@regression"}}
	assert.True(t, matchesFilter([]string{"@regression"}, groups))
}

func Test_MatchesFilter_AND_AllTagsPresent(t *testing.T) {
	assert.True(t, matchesFilter([]string{"@smoke", "@critical"}, []tagGroup{{"@smoke", "@critical"}}))
}

func Test_MatchesFilter_AND_MissingOneTag(t *testing.T) {
	assert.False(t, matchesFilter([]string{"@smoke"}, []tagGroup{{"@smoke", "@critical"}}))
}

func Test_MatchesFilter_ORofAND_FirstGroupFails_SecondPasses(t *testing.T) {
	groups := []tagGroup{
		{"@smoke", "@critical"}, // needs both – only @smoke → false
		{"@regression"},         // only @regression needed → true
	}
	assert.True(t, matchesFilter([]string{"@smoke", "@regression"}, groups))
}

func Test_MatchesFilter_ORofAND_BothGroupsFail(t *testing.T) {
	groups := []tagGroup{
		{"@smoke", "@critical"},
		{"@regression", "@flaky"},
	}
	assert.False(t, matchesFilter([]string{"@smoke", "@regression"}, groups))
}

func Test_MatchesFilter_Negation_TagAbsent(t *testing.T) {
	// ~@regression: passes when @regression is NOT present
	assert.True(t, matchesFilter([]string{"@smoke"}, []tagGroup{{"~@regression"}}))
}

func Test_MatchesFilter_Negation_TagPresent(t *testing.T) {
	// ~@regression: fails when @regression IS present
	assert.False(t, matchesFilter([]string{"@smoke", "@regression"}, []tagGroup{{"~@regression"}}))
}

func Test_MatchesFilter_AndWithNegation(t *testing.T) {
	// @smoke && ~@regression: must have @smoke AND not have @regression
	groups := []tagGroup{{"@smoke", "~@regression"}}
	assert.True(t, matchesFilter([]string{"@smoke"}, groups))
	assert.False(t, matchesFilter([]string{"@smoke", "@regression"}, groups))
	assert.False(t, matchesFilter([]string{"@regression"}, groups))
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

	result := filterFeatures([]*Feature{mustParseFeature(t, content)}, nil)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 2)
}

func Test_FilterFeatures_EmptyFeaturesList(t *testing.T) {
	result := filterFeatures([]*Feature{}, parseTagExpression("@smoke"))
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
		parseTagExpression("@smoke"),
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 1)
	assert.Equal(t, "Successful login", result[0].scenarios[0].Name)
	assert.Contains(t, string(result[0].Contents), "Successful login")
	assert.NotContains(t, string(result[0].Contents), "Failed login")
}

func Test_FilterFeatures_Include_OR_BothTagsMatchBothScenarios(t *testing.T) {
	content := `Feature: Login

  @smoke
  Scenario: Successful login
    Given I am on the login page

  @regression
  Scenario: Failed login
    Given I am on the login page`

	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		parseTagExpression("@smoke || @regression"),
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 2)
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
		parseTagExpression("@smoke && @critical"),
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
		parseTagExpression("@smoke"),
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 2)
}

func Test_FilterFeatures_Exclude_RemovesMatchingScenario(t *testing.T) {
	content := `Feature: Login

  @smoke
  Scenario: Successful login
    Given I am on the login page

  @regression
  Scenario: Failed login
    Given I am on the login page`

	// ~@regression means: exclude scenarios that have @regression
	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		parseTagExpression("~@regression"),
	)

	assert.Len(t, result, 1)
	assert.Len(t, result[0].scenarios, 1)
	assert.Equal(t, "Successful login", result[0].scenarios[0].Name)
	assert.NotContains(t, string(result[0].Contents), "Failed login")
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

	// @smoke && ~@regression: must have @smoke AND must NOT have @regression
	// → only Successful login matches
	result := filterFeatures(
		[]*Feature{mustParseFeature(t, content)},
		parseTagExpression("@smoke && ~@regression"),
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
		parseTagExpression("@nonexistent"),
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

	result := filterFeatures(features, parseTagExpression("@smoke"))

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
		parseTagExpression("@smoke"),
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
		parseTagExpression("@api"),
	)

	assert.Len(t, result, 1)
	// Rebuilt content must itself be valid parseable gherkin.
	rebuilt, err := parseFeatureContent(string(result[0].Contents))
	require.NoError(t, err)
	assert.Len(t, rebuilt.scenarios, 1)
	assert.Equal(t, "Create resource", rebuilt.scenarios[0].Name)
}

// ── ParseWithFilter (public API) ──────────────────────────────────────────────

func Test_ParseWithFilter_EmptyExpression_NilGroups(t *testing.T) {
	// Verify that an empty expression produces nil groups (no filtering).
	assert.Nil(t, parseTagExpression(""))
	assert.Nil(t, parseTagExpression("   "))
}
