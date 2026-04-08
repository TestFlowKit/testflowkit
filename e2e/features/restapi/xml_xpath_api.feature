@API @XML @XPATH
Feature: XML API XPath validation

    Scenario: Validate an XML customer profile response using XPath
        Given I prepare a request to "jsonplaceholder.get_xml_customer_profile"
        When I set the path parameter "id" to "42"
        And I set the query parameter "ok" to "true"
        And I send the request
        Then the response status code should be 200
        And the response should have field "//customer/@id"
        And the response field "//customer/@id" should be "42"
        And the response field "//status" should be "ok"
        And the response field "//query/ok" should be "true"
        And the response field "//info/email" should be "contact@example.com"
        And the response field "//info/email" should match pattern "^[^@]+@example\.com$"
        And the response field "//info/verified" should have type "boolean"
        And the response field "//tags/tag" should have type "list"
        And the response path "//tags/tag" should have 2 elements
        And the response field "//tags/tag[@priority='high']" should be "vip"
        And the response field "//orders/order[id='1002']/total" should be "10.00"
