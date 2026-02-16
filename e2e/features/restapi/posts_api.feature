@API @POSTS
Feature: Posts API Testing


    @macro
    Scenario: I try to create a new post with the following details:
        this macro was created for test a bug in macro application
        the macro did not include the docstring

        Given I prepare a request to "jsonplaceholder.create_post"
        When I set the request body to:
            """
            {
            "title": "${title}",
            "body": "${body}",
            "userId": ${userId}
            }
            """
        And I send the request

    Scenario: Retrieve all posts successfully
        Given I prepare a request to "jsonplaceholder.get_posts"
        When I send the request
        Then the response status code should be 200
        And the response should have field "0.userId"
        And the response should have field "0.id"
        And the response should have field "0.title"

    Scenario: Retrieve a specific post by ID
        Given I prepare a request to "jsonplaceholder.get_post_by_id"
        When I set the path parameter "id" to "1"
        And I send the request
        Then the response status code should be 200
        And the response should contain "{{ env.post_content }}"
        And the response should have field "id"
        And the response should have field "userId"
        And the response should have field "title"
        And the response should have field "body"

    Scenario: Create a new post
        When I try to create a new post with the following details:
            | title           | body                                     | userId |
            | Test Post Title | This is a test post body for API testing | 1      |
        Then the response status code should be 201
        And the response header "test" should equal "micro"
        And the response header "Location" should match pattern "/posts/\d+"



    Scenario: Delete a post
        Given I prepare a request to "jsonplaceholder.delete_post"
        When I set the path parameter "id" to "1"
        And I send the request
        Then the response status code should be 200

    Scenario: Retrieve posts with query parameters
        Given I prepare a request to "jsonplaceholder.get_posts"
        When I set the query parameter "_limit" to "5"
        And I send the request
        Then the response status code should be 200
        And the response should have field "0.userId"
        And the response should have field "0.id"
        And the response should have field "0.title"

    Scenario: Retrieve posts for specific user
        Given I prepare a request to "jsonplaceholder.get_posts"
        When I set the query parameter "userId" to "1"
        And I send the request
        Then the response status code should be 200
        And the response should have field "0.userId"
        And the response should have field "0.id"
        And the response should have field "0.title"


    @from_file
    Scenario: Create a new post with json file
        Given I prepare a request to "jsonplaceholder.create_post"
        When I set the request body from file "features/restapi/new_post.json"
        And I send the request
        Then the response status code should be 201
        

    Scenario: Validate response field types and patterns
        Given I prepare a request to "jsonplaceholder.get_post_by_id"
        When I set the following path parameters:
            | id | 1 |
        And I send the request
        Then the response status code should be 200
        And the response field "id" should have type "integer"
        And the response field "userId" should have type "integer"
        And the response field "body" should have type "string"
        And the response field "title" should match pattern "\w+"

