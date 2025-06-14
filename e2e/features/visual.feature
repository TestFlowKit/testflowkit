@VISUAL
Feature: visual e2e tests

  Background:
    Given I open a new browser tab
    Then the user goes to the visual e2e page

  Scenario: User should see certain things on page
    Given I should not see "L'élément a été caché." on the page
    And I should see "Cet élément va disparaître quand vous cliquerez sur le bouton." on the page
    When I click on the button which contains "Cacher l'élément"
    Then I should not see "Cet élément va disparaître quand vous cliquerez sur le bouton." on the page
    And I should see "L'élément a été caché." on the page

  @doubleClick
  Scenario: double click on element which contains
    Given I should not see "Vous avez double cliqué sur le bouton." on the page
    When I double click on the button which contains "double click"
    Then I should see "Vous avez double cliqué sur le bouton." on the page

  @doubleClick
  Scenario: double click on element
    Given I should not see "Vous avez double cliqué sur le bouton." on the page
    When I double click on double click button
    Then I should see "Vous avez double cliqué sur le bouton." on the page

  @visibility
  Scenario: element should exist but not visible
    Then the hidden button should exist
    And the hidden button should not be visible

  @visibility
  Scenario: element should not exist so not visible
    Then the non-existent button should not exist
