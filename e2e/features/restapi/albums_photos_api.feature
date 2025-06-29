@API @ALBUMS @PHOTOS
Feature: Albums and Photos API Testing


    Scenario: Retrieve all albums successfully (response body fields)
        Given I prepare a request for the "get_albums" endpoint
        When I send the request
        Then the response status code should be 200
        And the response body should contain "userId"
        And the response body should contain "id"
        And the response body should contain "title"
        And the response body path "0.userId" should exist
        And the response body path "0.id" should exist
        And the response body path "0.title" should exist


    Scenario: Retrieve a specific album by ID (with path params)
        Given I prepare a request for the "get_album_by_id" endpoint
        When I set the following path params:
            | id | 1 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "quidem molestiae enim"
        And the response body path "id" should exist
        And the response body path "userId" should exist
        And the response body path "title" should exist

    Scenario: Retrieve albums for a specific user (with query params)
        Given I prepare a request for the "get_albums" endpoint
        When I have the following query parameters:
            | userId | 1 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "userId"
        And the response body should contain "id"
        And the response body should contain "title"
        And the response body path "0.userId" should exist
        And the response body path "0.id" should exist
        And the response body path "0.title" should exist

    Scenario: Retrieve all photos successfully
        Given I prepare a request for the "get_photos" endpoint
        When I send the request
        Then the response status code should be 200
        And the response body should contain "albumId"
        And the response body should contain "id"
        And the response body should contain "title"
        And the response body should contain "url"
        And the response body should contain "thumbnailUrl"
        And the response body path "0.albumId" should exist
        And the response body path "0.id" should exist
        And the response body path "0.title" should exist
        And the response body path "0.url" should exist
        And the response body path "0.thumbnailUrl" should exist

    Scenario: Retrieve photos for a specific album
        Given I prepare a request for the "get_photos_by_album" endpoint
        When I set the following path params:
            | id | 1 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "albumId"
        And the response body should contain "id"
        And the response body should contain "title"
        And the response body should contain "url"
        And the response body should contain "thumbnailUrl"
        And the response body path "0.albumId" should exist
        And the response body path "0.id" should exist
        And the response body path "0.title" should exist
        And the response body path "0.url" should exist
        And the response body path "0.thumbnailUrl" should exist

    Scenario: Retrieve photos with query parameters
        Given I prepare a request for the "get_photos" endpoint
        When I have the following query parameters:
            | albumId | 1 |
            | _limit  | 5 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "albumId"
        And the response body should contain "id"
        And the response body should contain "title"
        And the response body should contain "url"
        And the response body should contain "thumbnailUrl"

    Scenario: Retrieve albums with pagination
        Given I prepare a request for the "get_albums" endpoint
        When I have the following query parameters:
            | _page  | 1  |
            | _limit | 10 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "userId"
        And the response body should contain "id"
        And the response body should contain "title"

    Scenario: Verify album structure for specific user
        Given I prepare a request for the "get_albums" endpoint
        When I have the following query parameters:
            | userId | 1 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "quidem molestiae enim"
        And the response body should contain "sunt qui excepturi placeat culpa"
        And the response body should contain "omnis laborum odio"
        And the response body path "0.userId" should exist
        And the response body path "0.id" should exist
        And the response body path "0.title" should exist

    Scenario: Verify photo URLs are valid
        Given I prepare a request for the "get_photos" endpoint
        When I have the following query parameters:
            | _limit | 1 |
        And I send the request
        Then the response status code should be 200
        And the response body should contain "https://"
        And the response body should contain "via.placeholder.com"
        And the response body path "0.url" should exist
        And the response body path "0.thumbnailUrl" should exist