@NAVIGATION
Feature: navigation e2e tests

  Scenario: a user can navigate between pages
    Given I open a new browser tab
    When the user goes to the google page
    Then the user should be navigated to google page
