@VISUAL
Feature: visual e2e tests

  Background:
    Given the user opens a new browser tab
    Then the user goes to the visual e2e page

  Scenario: User should see certain things on page
    Given the user should not see "L'élément a été caché." on the page
    And the user should see "Cet élément va disparaître quand vous cliquerez sur le bouton." on the page
    When the user clicks on the button which contains "Cacher l'élément"
    Then the user should not see "Cet élément va disparaître quand vous cliquerez sur le bouton." on the page
    And the user should see "L'élément a été caché." on the page

  @doubleClick
  Scenario: double click on element which contains
    Given the user should not see "Vous avez double cliqué sur le bouton." on the page
    When the user double clicks on the button which contains "double click"
    Then the user should see "Vous avez double cliqué sur le bouton." on the page

  @doubleClick
  Scenario: double click on element
    Given the user should not see "Vous avez double cliqué sur le bouton." on the page
    When the user double clicks on double click button
    Then the user should see "Vous avez double cliqué sur le bouton." on the page

  @visibility
  Scenario: element should exist but not visible
    Then the hidden button should exist
    And the hidden button should not be visible

  @visibility
  Scenario: element should not exist so not visible
    Then the non-existent button should not exist
