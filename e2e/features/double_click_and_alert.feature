Feature: Double Click and Alert Handling
  As a user
  I want to interact with elements using double-click
  And handle alerts that appear
  So that I can perform actions that require these interactions

  Scenario: Double click on a document and handle confirmation alert
    Given I navigate to "Test" page
    When I double click on "Document 1"
    Then I should see "Are you sure you want to open Document 1?" in the alert and "accept" it
    And I should see "Opening Document 1" in the alert and "accept" it

  Scenario: Double click on a document and dismiss confirmation alert
    Given I navigate to "Test" page
    When I double click on "Document 2"
    Then I should see "Are you sure you want to open Document 2?" in the alert and "dismiss" it
    And I should not see "Opening Document 2" on the page 