@PAGE_ASSERTIONS @FRONTEND
Feature: Page Title and URL Assertions

    Scenario: Verify that the page title matches exactly
        When the user goes to the "form e2e" page
        Then the page title should be "Formulaire de test E2E"

    Scenario: Verify that the current URL contains the page path
        When the user goes to the "form e2e" page
        Then the current URL should contain "form"

    Scenario: Verify that the current URL contains the details page path
        When the user goes to the "details e2e" page
        Then the current URL should contain "details"

    Scenario: Verify that the current URL contains the table page path
        When the user goes to the "table e2e" page
        Then the current URL should contain "table"

    Scenario: Verify that the current URL contains the visual page path
        When the user goes to the "visual e2e" page
        Then the current URL should contain "visual"