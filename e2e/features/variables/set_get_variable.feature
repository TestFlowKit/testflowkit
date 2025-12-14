@TEST_VARIABLES
Feature: variables testing


    Scenario: Write API response field into another field
        Given I prepare a request to "get_post_by_id"
        And I set the following path parameters:
            | id | 1 |
        And I send the request
        And I store the JSON path "title" from the response into "postTitle" variable
        When the user goes to the "form e2e" page
        And the user enters "{{postTitle}}" into the "text" field
        Then the value of the text field should be "sunt aut facere repellat provident occaecati excepturi optio reprehenderit"


    Scenario: Write html element context into another field
        The page title is "Formulaire de test E2E"

        Given the user goes to the "form e2e" page
        And I store the content of "page title" into "pageTitle" variable
        And the user enters "{{ pageTitle }}" into the "text" field
        Then the value of the text field should be "Formulaire de test E2E"