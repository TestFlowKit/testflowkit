@FORM
Feature: Form e2e tests

  Background:
    Given I open a new private browser tab
    And the user goes to the "form e2e" page

  @DROPDOWN
  Scenario Outline: a user can select dropdown value
    When I select "<selection>" into the <dropdown> dropdown
    Then the <dropdown> dropdown should have "<selection>" selected

    Examples:
      | dropdown | selection          |
      | test     | Option 2           |
      | multiple | Option 2, Option 1 |

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
    Then I type "<value>" into the <type> field
    When the <type> field should be contain "<value>"

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
