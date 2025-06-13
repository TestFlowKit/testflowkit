@FORM
Feature: Form e2e tests

  Background:
    Given I open a new private browser tab
    And the user goes to the "form e2e" page

  @DROPDOWN @SELECT_BY_TEXT
  Scenario: a user can select dropdown value by text
    When the user selects the option with text "Option 2" from the "test" dropdown

  @DROPDOWN @MULTIPLE @SELECT_BY_TEXT
  Scenario: a user can select multiple dropdown values by text
    When the user selects the options with text "Option 2, Option 1" from the "multiple" dropdown

  @DROPDOWN @SELECT_BY_VALUE
  Scenario: a user can select dropdown value
    When the user selects the option with value "option2" from the "test" dropdown

  @DROPDOWN @MULTIPLE @SELECT_BY_VALUE
  Scenario: a user can select multiple dropdown values by values
    When the user selects the options with values "option2, option1" from the "multiple" dropdown

  @DROPDOWN @MULTIPLE @SELECT_BY_INDEX
  Scenario: a user can select multiple dropdown values by index
    When the user selects the option at index 2 from the "test" dropdown

  @CHECKBOX @CHECKED
  Scenario: a user can check a checkbox
    When the user checks the "test" checkbox
    Then the test checkbox should be checked

  @CHECKBOX @UNCHECKED
  Scenario: a user can uncheck a checkbox
    Given I already checked test checkbox
    When the user unchecks the "test" checkbox
    Then the test checkbox should be unchecked

  @TEXT_FIELD
  Scenario Outline: a user can type into <type> field
    When the user enters "<value>" into the "<type>" field
    Then the <type> field should be contain "<value>"

    Examples:
      | type     | value             |
      | text     | Hello Test !      |
      | textarea | Hello Test area ! |

  @RADIO
  Scenario: a user can select radio button
    Given the test radio button should be unselected
    When the user selects the test radio button
    Then the test radio button should be selected

  @RADIO @UNSELECTED
  Scenario: a user can unselect radio button
    Given the user selects the test radio button
    And the test radio button should be selected
    When the user selects the second test radio button
    Then the test radio button should be unselected
