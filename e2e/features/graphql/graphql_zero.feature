@GRAPHQL @API
Feature: GraphQL Zero API Testing

  Scenario: Fetch a user successfully
    Given I prepare a request to "almansi.get_user"
    And I set the following GraphQL variables:
      | id | 1 |
    When I send the request
    Then the GraphQL response should not have errors
    And the response should have field "user.username"
    And the response should contain "Bret"

  Scenario: Create a post successfully with random content
    Given I prepare a request to "almansi.create_post"
    And I store the value "{{ rand:words:count=4 }}" into "postTitle" variable
    And I store the value "{{ rand:words:count=10 }}" into "postBody" variable
    And I set the following GraphQL variables:
      | input | {"title": "{{postTitle}}", "body": "{{postBody}}"} |
    When I send the request
    Then the GraphQL response should not have errors
    And the response should have field "createPost.id"
    And the response should contain "{{postTitle}}"

  Scenario: Fetch a user with error handling (Simulated)
    # GraphQLZero might not error easily on ID, but let's try to verify we can check for errors if they occurred.
    # Since we can't easily force an error on this public API without changing the query structure (which is fixed in config),
    # we will just verify the success path again but using error checking steps negatively.
    Given I prepare a request to "almansi.get_user"
    And I set the following GraphQL variables:
      | id | 1 |
    When I send the request
    Then the GraphQL response should not have errors
# If we could force an error, we would use:
# Then the GraphQL response should have errors
# And the GraphQL error message should contain "Some error"
