@ASSERTIONS @FRONTEND
Feature: Text Assertions

  Background:
    Given the user opens a new browser tab

  Scenario: Verify that product name is displayed and matches exactly the expected value
    Given the user goes to the details e2e page
    Then the product name element should contain the text "Ordinateur de Bord pour Rameur"

  Scenario: Verify that product name is displayed and matches partially
    Given the user goes to the details e2e page
    Then the product name element should contain the text "Ordinateur de Bord"

  Scenario: Verify that product description is displayed and matches partially
    Given the user goes to the details e2e page
    Then the product description element should contain the text "Cet ordinateur de rameur vous permet de suivre vos performances en temps réel"

  Scenario: Verify that product name displayed not matches with an incorrect value
    Given the user goes to the details e2e page
    Then the product name element should not contain the text "Nintendo switch 2"

  @TEXT_FIELD
  Scenario Outline: a user can type into <type> field
    Given the user goes to the "form e2e" page
    When the user enters "<value>" into the "<type>" field
    Then the value of the <type> field should be "<value>"

    Examples:
      | type     | value             |
      | text     | Hello Test !      |
      | textarea | Hello Test area ! |
