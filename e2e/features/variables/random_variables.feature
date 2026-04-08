@TEST_VARIABLES @RANDOM_DATA
Feature: random data generation in e2e tests

    @FRONTEND @FORM
    Scenario: Use a random email in a form field
        Given the user goes to the "form e2e" page
        When I store the value "{{ rand:email:domain=kil.com }}" into "randomEmail" variable
        And the user enters "{{randomEmail}}" into the "text" field
        Then the value of the text field should be "{{randomEmail}}"

    @FRONTEND @FORM
    Scenario: Reuse a generated French phone number in E.164 format
        Given the user goes to the "form e2e" page
        When I store the value "{{ rand:phone:country=FR,format=e164 }}" into "phoneNumber" variable
        And the user enters "{{phoneNumber}}" into the "text" field
        Then the value of the text field should be "{{phoneNumber}}"

    @FRONTEND @FORM
    Scenario: Use a regex-generated reference code in the UI
        Given the user goes to the "form e2e" page
        When I store the value "{{ rand:regex:pattern=[A-Z]{3}-\d{4} }}" into "referenceCode" variable
        And the user enters "{{referenceCode}}" into the "text" field
        Then the value of the text field should be "{{referenceCode}}"
