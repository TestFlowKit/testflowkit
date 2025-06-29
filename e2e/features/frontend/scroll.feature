Feature: scroll e2e tests

  Background:
    Given the user opens a new browser tab
    Then the user goes to the scroll e2e page

  @scroll
  Scenario: scroll to element
    When the user scrolls to the "scroll target" element
