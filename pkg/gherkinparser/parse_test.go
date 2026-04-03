package gherkinparser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse_SelectContactFeatureFile(t *testing.T) {
	tempDir := t.TempDir()
	featurePath := filepath.Join(tempDir, "select-contact.feature")
	featureContent := `Feature: Selects contact by id

    @TEST
    Scenario: Select contact by id successfully
        Given I prepare a request to "contact.select_by_id"
        And I store the value "FRA90-00000353" into "contact_id" variable
        And I set the following path parameters:
            | id | {{ contact_id }} |
        When I send the request
        Then the response status code should be 200

    Scenario: ok
        Given I prepare a request to "contact.select_by_id"
        And I store the value "non_existent_contact_id" into "contact_id" variable
        And I set the following path parameters:
            | id | {{ contact_id }} |
        When I send the request
        Then the response status code should be 404
`

	require.NoError(t, os.WriteFile(featurePath, []byte(featureContent), 0o600))

	features := Parse(tempDir)
	require.Len(t, features, 1)

	feature := features[0]
	require.NotNil(t, feature)
	assert.Equal(t, "Selects contact by id", feature.Name)
	require.Len(t, feature.scenarios, 2)

	firstScenario := feature.scenarios[0]
	require.NotNil(t, firstScenario)
	assert.Equal(t, "Select contact by id successfully", firstScenario.Name)
	assert.ElementsMatch(t, []string{"@TEST"}, collectTagNames(firstScenario.Tags))
	require.Len(t, firstScenario.Steps, 5)
	assert.Equal(t, `I prepare a request to "contact.select_by_id"`, firstScenario.Steps[0].Text)
	require.NotNil(t, firstScenario.Steps[2].DataTable)
	require.Len(t, firstScenario.Steps[2].DataTable.Rows, 1)
	require.Len(t, firstScenario.Steps[2].DataTable.Rows[0].Cells, 2)
	assert.Equal(t, "id", firstScenario.Steps[2].DataTable.Rows[0].Cells[0].Value)
	assert.Equal(t, "{{ contact_id }}", firstScenario.Steps[2].DataTable.Rows[0].Cells[1].Value)
	assert.Equal(t, "the response status code should be 200", firstScenario.Steps[4].Text)

	secondScenario := feature.scenarios[1]
	require.NotNil(t, secondScenario)
	assert.Equal(t, "ok", secondScenario.Name)
	assert.Empty(t, collectTagNames(secondScenario.Tags))
	require.Len(t, secondScenario.Steps, 5)
	assert.Equal(t, `I store the value "non_existent_contact_id" into "contact_id" variable`, secondScenario.Steps[1].Text)
	require.NotNil(t, secondScenario.Steps[2].DataTable)
	assert.Equal(t, "the response status code should be 404", secondScenario.Steps[4].Text)
}
