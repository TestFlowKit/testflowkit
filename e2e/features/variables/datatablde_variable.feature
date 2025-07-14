@TEST_VARIABLES
Feature: variables testing with table data

    @tableData
    Scenario: Store table data into variable in order to use it in another step
        Given I store the "LCD rétroéclairé" into "screenType" variable
        And the user opens a new private browser tab
        When the user goes to the details e2e page
        Then the user should see "computer" details on the page
            | name        | Ordinateur de Bord pour Rameur                                                |
            | description | Cet ordinateur de rameur vous permet de suivre vos performances en temps réel |
            | price       | 149,99 €                                                                      |
            | screen type | {{ screenType }}                                                              |
