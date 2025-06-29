@API @POSTS
Feature: Posts API Testing


    Scenario: Retrieve all posts successfully
        Given I prepare a request for the "get_posts" endpoint
        When I send the request
        Then the response status code should be 200
        And the response body should contain "userId"
        And the response body should contain "id"
        And the response body should contain "title"
        And the response body should contain "body"
        And the response body path "0.userId" should exist
        And the response body path "0.id" should exist
        And the response body path "0.title" should exist

    Scenario: Retrieve a specific post by ID
        Given I prepare a request for the "get_post_by_id" endpoint
        When I set the following path params:
            | id | 1 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "sunt aut facere repellat provident occaecati excepturi optio reprehenderit"
        And the response body path "id" should exist
        And the response body path "userId" should exist
        And the response body path "title" should exist
        And the response body path "body" should exist

    Scenario: Create a new post
        Given I prepare a request for the "create_post" endpoint
        When I set the request body to:
            """
            {
                "title": "Test Post Title",
                "body": "This is a test post body for API testing",
                "userId": 1
            }
            """
        And I send the request
        Then the response status code should be 201


    Scenario: Delete a post
        Given I prepare a request for the "delete_post" endpoint
        When I set the following path params:
            | id | 1 |
        And I send the request
        Then the response status code should be 200

    Scenario: Retrieve posts with query parameters
        Given I prepare a request for the "get_posts" endpoint
        When I have the following query parameters:
            | _limit | 5 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "userId"
        And the response body should contain "id"
        And the response body should contain "title"

    Scenario: Retrieve posts for specific user
        Given I prepare a request for the "get_posts" endpoint
        When I have the following query parameters:
            | userId | 1 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "userId"
        And the response body should contain "id"
        And the response body should contain "title"