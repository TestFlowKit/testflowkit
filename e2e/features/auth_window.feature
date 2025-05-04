Feature: Authentication in popup windows
  As a user
  I want to handle authentication flows that open in popup windows
  So that I can test login functionality

  Scenario: Login through a popup authentication window
    Given I navigate to "form_e2e" page
    When I click on "visual_test_button" element
    And I wait for a new window to open within "5s"
    And I switch to the new window
    # The test can now interact with elements in the new window
    Then I should see "test_checkbox" element
    # After completing authentication actions
    When I switch back to the original window
    # Continue with test in the original window
    Then I should see "visual_test_button" element 