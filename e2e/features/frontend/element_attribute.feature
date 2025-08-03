@ELEMENT_ATTRIBUTE_ASSERTIONS
Feature: Element Attribute Assertions

    Scenario: Verify that an element attribute matches expected value
        When the user goes to the "form e2e" page
        Then the "id" attribute of the "page_title" element should be "page-title"

    Scenario: Verify that an element type attribute matches expected value
        When the user goes to the "form e2e" page
        Then the "type" attribute of the "text_field" element should be "text"

    Scenario: Verify that an element name attribute matches expected value
        When the user goes to the "form e2e" page
        Then the "name" attribute of the "text_field" element should be "text-input"