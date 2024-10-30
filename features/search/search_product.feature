Feature: Products search (JUMIA)

  Background:
    Given I open a new private browser tab
    And I navigate to jumia home page
    And I click on "newsletter modal close" button

  Scenario: A user search an iphone 15 on the web shop
    When I fill the search input with "Iphone 15"
    And I press the enter button
    Then I am redirected to iphone 15 results page
     And I must see on page 15 iphone 15


  Scenario: A user search an unexciting product on the web shop
    When I fill the search input with "yyyyy"
    And I press the enter button
    And I am redirected to yyyyy results page
    Then I must see on page a link with text "Retour à l'accueil"