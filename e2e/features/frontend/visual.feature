@VISUAL @FRONTEND
Feature: visual e2e tests

  Background:
    Given the user is on page
      | page_name  |
      | visual e2e |

  Scenario: User should see certain things on page
    Given the visual disappearing element should contain the text "Cet élément va disparaître quand vous cliquerez sur le bouton."
    And the visual hidden message should not exist
    When the user clicks on the button which contains "Cacher l'élément"
    Then the visual disappearing element should not be visible
    And the visual hidden message should exist

  @doubleClick
  Scenario: double click on element which contains
    Given the visual double click message should not exist
    When the user double clicks on the button which contains "double click"
    Then the visual double click message should exist

  @doubleClick
  Scenario: double click on element
    Given the visual double click message should not exist
    When the user double clicks on double click button
    Then the visual double click message should exist

  @visibility
  Scenario: element should exist but not visible
    Then the hidden button should exist
    And the hidden button should not be visible

  @visibility
  Scenario: element should not exist so not visible
    Then the non-existent button should not exist
