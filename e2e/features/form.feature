@FORM
Feature: Form e2e tests

  Background:
    Given I open a new private browser tab
    And I navigate to "form e2e" page



  @DROPDOWN
  Scenario Outline: a user can select dropdown value
    When I select "<selection>" into the <dropdown> dropdown
    Then the <dropdown> dropdown should have "<selection>" selected
    Examples:
      | dropdown | selection          |
      | test     | Option 2           |
      | multiple | Option 2, Option 1 |


  @CHECKBOX
  Scenario: a user can check a checkbox
    When I click on test checkbox
    Then the test checkbox should be checked

  @CHECKBOX
  Scenario: a user can uncheck a checkbox
    Given I click on test checkbox
    And the test checkbox should be checked
    When I click on test checkbox
    Then the test checkbox should be unchecked

  @TEXT_FIELD
  Scenario Outline: a user can type into <type> field
    Then I type "<value>" into the <type> field
    When the <type> field should be contain "<value>"
    Examples:
      | type     | value             |
      | text     | Hello Test !      |
      | textarea | Hello Test area ! |

