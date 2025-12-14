@API @POSTS
Feature: Posts API Testing


    @macro
    Scenario: I try to create a new post with the following details:
        this macro was created for test a bug in macro application
        the macro did not include the docstring

        Given I prepare a request to "create_post"
        When I set the request body to:
            """
            {
            "title": "|title|",
            "body": "|body|",
            "userId": |userId|
            }
            """
        And I send the request

    Scenario: Retrieve all posts successfully
        Given I prepare a request to "get_posts"
        When I send the request
        Then the response status code should be 200
        And the response should have field "0.userId"
        And the response should have field "0.id"
        And the response should have field "0.title"

    Scenario: Retrieve a specific post by ID
        Given I prepare a request to "get_post_by_id"
        When I set the following path parameters:
            | id | 1 |
        And I send the request
        Then the response status code should be 200
        And the response should contain "sunt aut facere repellat provident occaecati excepturi optio reprehenderit"
        And the response should have field "id"
        And the response should have field "userId"
        And the response should have field "title"
        And the response should have field "body"

    Scenario: Create a new post
        When I try to create a new post with the following details:
            | title           | body                                     | userId |
            | Test Post Title | This is a test post body for API testing | 1      |
        Then the response status code should be 201


    Scenario: Delete a post
        Given I prepare a request to "delete_post"
        When I set the following path parameters:
            | id | 1 |
        And I send the request
        Then the response status code should be 200

    Scenario: Retrieve posts with query parameters
        Given I prepare a request to "get_posts"
        When I set the following query parameters:
            | _limit | 5 |
        And I send the request
        Then the response status code should be 200
        And the response should have field "0.userId"
        And the response should have field "0.id"
        And the response should have field "0.title"

    Scenario: Retrieve posts for specific user
        Given I prepare a request to "get_posts"
        When I set the following query parameters:
            | userId | 1 |
        And I send the request
        Then the response status code should be 200
        And the response should have field "0.userId"
        And the response should have field "0.id"
        And the response should have field "0.title"