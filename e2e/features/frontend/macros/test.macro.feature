Feature: macro

  @macro
  Scenario: the user already checked test checkbox
    Given the user checks the "test" checkbox
    And the test checkbox should be checked


  @macro
  Scenario: the user is on page
    Given the user opens a new browser tab
    When the user goes to the |page_name| page
