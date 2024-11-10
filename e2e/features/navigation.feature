@NAVIGATION
Feature: navigation e2e tests

      Scenario: a user can navigate between pages
        Given I open a new browser tab
        When I navigate to google page
        Then I should be navigated to google page

