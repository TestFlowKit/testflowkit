@API @ALBUMS @PHOTOS
Feature: Albums and Photos API Testing

    Scenario: Retrieve all albums successfully (response body fields)
        Given I prepare a request to "jsonplaceholder.get_albums"
        When I send the request
        Then the response status code should be 200
        And the response should have field "0.userId"
        And the response should have field "0.id"


    Scenario: Retrieve a specific album by ID (with path params)
        Given I prepare a request to "jsonplaceholder.get_album_by_id"
        When I set the following path parameters:
            | id | 1 |
        And I send the request
        Then the response status code should be 200
        And the response field "title" should be "quidem molestiae enim"
        And the response should have field "id"
        And the response should have field "userId"
        And the response should have field "title"

    Scenario: Retrieve albums for a specific user (with query params)
        Given I prepare a request to "jsonplaceholder.get_albums"
        When I set the following query parameters:
            | userId | 1 |
        And I send the request
        Then the response status code should be 200
        And the response should have field "0.userId"
        And the response should have field "0.id"
      

    Scenario: Retrieve all photos successfully
        Given I prepare a request to "jsonplaceholder.get_photos"
        When I send the request
        Then the response status code should be 200
        And the response should have field "0.albumId"



    Scenario: Retrieve photos for a specific album
        Given I prepare a request to "jsonplaceholder.get_photos_by_album"
        When I set the path parameter "id" to "1"
        And I send the request
        Then the response status code should be 200
        And the response should have field "0.albumId"


    Scenario: Retrieve photos with query parameters
        Given I prepare a request to "jsonplaceholder.get_photos"
        When I set the following query parameters:
            | albumId | 1 |
            | _limit  | 5 |
        And I send the request
        Then the response status code should be 200
        And the response should have field "0.id"
    Scenario: Retrieve albums with pagination
        Given I prepare a request to "jsonplaceholder.get_albums"
        When I set the following query parameters:
            | _page  | 1  |
            | _limit | 10 |
        And I send the request
        Then the response status code should be 200
        And the response should have field "0.id"

    Scenario: Verify album structure for specific user
        Given I prepare a request to "jsonplaceholder.get_albums"
        And I send the request
        Then the response status code should be 200
        And the response should contain "quidem molestiae enim"
        And the response should contain "sunt qui excepturi placeat culpa"
        And the response should have field "0.userId"
        And the response should have field "0.id"


    Scenario: Verify photo URLs are valid
        Given I prepare a request to "jsonplaceholder.get_photos"
        And I set the query parameter "_limit" to "1"
        And I send the request
        Then the response status code should be 200
        Then the response field "0.url" should contain "https://via.placeholder.com"
        Then the response field "0.thumbnailUrl" should contain "https://via.placeholder.com"