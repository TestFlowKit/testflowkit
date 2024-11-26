@TABLE
Feature: Table e2e tests

  Background:
      Given I open a new browser tab
      When I navigate to table e2e page


  Scenario: User should see a specific table row
    Then I should see a row containing the following elements
      | name      | description                        | price   |
      | Produit 1 | Description détaillée du produit 1 | 19.99 € |


  Scenario: User should not see a specific table row
    Then I should not see a row containing the following elements
      | name      | description                        | price   |
      | Produit 1 | Description détaillée du produit 0 | 19.99 € |


  Scenario: User should click a specific table row
    When I click on the row containing the following elements
      | name      | description                        | price   |
      | Produit 1 | Description détaillée du produit 1 | 19.99 € |
    Then I should see "Description détaillée du produit 1 clicked !" on the page


  Scenario: User should see a table with following headers
    Then I should see a table with the following headers
      | product name        | Produit     |
      | product description | Description |
      | product price       | Prix        |
