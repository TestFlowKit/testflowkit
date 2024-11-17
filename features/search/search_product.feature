Feature: Products search (JUMIA)

  Background:
    Given I open a new browser tab
    And I navigate to jumia home page
    And I click on newsletter modal close button

  Scenario: A user search an iphone 15 on the web shop
    When I type "Iphone 15" into the search field
    And I press the "enter" button
    Then I should be navigated to iphone 15 results page
     And I should see 13 iphone 15


  Scenario: A user search an unexciting product on the web shop
    When I type "yyyyy" into the search field
       And I press the "enter" button
    Then I should be navigated to yyyyy results page
    And I should see a link which contains "Retour à l'accueil"