@SAMPLE
Feature: TestFlowKit Documentation Site Sample Test
  This is a sample feature file that demonstrates common TestFlowKit testing patterns.

  Background:
    Given the user opens a new browser tab

  Scenario: Navigate to documentation site and verify homepage
    When the user goes to the "home" page
    Then the page title should be "TestFlowKit - Behavior-Driven Testing Framework"
    And the user should see "TestFlowKit" on the page
    And the "get started button" should be visible

  Scenario: Navigate to Get Started page and verify content
    Given the user goes to the "home" page
    When the user clicks the "get started" button
    Then the current URL should contain "getting-started"
    And the page title should be "Introduction | TestFlowKit"

  Scenario: Explore the Sentences documentation
    When the user goes to the "sentences" page
    Then the current URL should contain "sentences"
    And the user should see "Step Definitions" on the page
    And the "sentence filter field" should be visible
    When the user enters "click" into the "sentence filter" field
    Then the user should see "clicks on an element which contains a specific text." on the page
