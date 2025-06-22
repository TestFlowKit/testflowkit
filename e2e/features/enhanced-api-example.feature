Feature: Enhanced API and Frontend Testing
  As a tester using TestFlowKit Enhanced Configuration
  I want to test both API endpoints and frontend interactions
  So that I can ensure comprehensive application testing

  Background:
    Given I prepare a request for the "setup_test_data" endpoint
    When I send the request
    Then the response status code should be 200

  @api @smoke
  Scenario: User Management API with Enhanced Configuration
    # Create a new user via API
    Given I prepare a request for the "create_user" endpoint
    And I set the request body to:
      """
      {
        "name": "John Doe",
        "email": "john.doe@example.com",
        "password": "SecurePassword123",
        "role": "user"
      }
      """
    When I send the request
    Then the response status code should be 201
    And the response body should contain "success"
    And I store the response body path "user.id" as "created_user_id"

    # Retrieve the created user
    Given I prepare a request for the "get_user" endpoint
    And I set the path parameter "id" to "${created_user_id}"
    When I send the request
    Then the response status code should be 200
    And the response body should contain "john.doe@example.com"
    And the response header "Content-Type" should be "application/json"

    # Update the user information
    Given I prepare a request for the "update_user" endpoint
    And I set the path parameter "id" to "${created_user_id}"
    And I set the request body to:
      """
      {
        "name": "John Updated Doe",
        "email": "john.updated@example.com"
      }
      """
    When I send the request
    Then the response status code should be 200
    And the response body should contain "john.updated@example.com"

  @frontend @smoke
  Scenario: Login Process with Enhanced Element Selectors
    # Navigate to login page using configured page name
    When the user goes to the "login" page
    Then the user should be navigated to "login" page

    # Fill login form using fallback selectors
    When the user fills the "email_field" on "login_page" with "john.updated@example.com"
    And the user fills the "password_field" on "login_page" with "SecurePassword123"
    And the user clicks on the "login_button" element on "login_page"

    # Verify successful login
    Then the user should be navigated to "dashboard" page
    And the user waits for the "welcome_message" element on "dashboard_page"
    And the user should see on the page the "user_profile_dropdown" element

  @integration @api @frontend
  Scenario: End-to-End User Journey with API Setup and Frontend Verification
    # API Setup: Create user and product data
    Given I prepare a request for the "create_user" endpoint
    And I set the request body to:
      """
      {
        "name": "Test Customer",
        "email": "customer@example.com",
        "password": "TestPassword123",
        "role": "customer"
      }
      """
    When I send the request
    Then the response status code should be 201
    And I store the response body path "user.id" as "customer_id"

    # Create a product via API
    Given I prepare a request for the "create_product" endpoint
    And I set the request body to:
      """
      {
        "name": "Test Product",
        "description": "A product for testing",
        "price": 29.99,
        "category": "electronics"
      }
      """
    When I send the request
    Then the response status code should be 201
    And I store the response body path "product.id" as "product_id"

    # Frontend Verification: Login and verify product is visible
    When the user goes to the "login" page
    And the user fills the "email_field" on "login_page" with "customer@example.com"
    And the user fills the "password_field" on "login_page" with "TestPassword123"
    And the user clicks on the "login_button" element on "login_page"
    
    # Navigate to products page
    When the user goes to the "products" page
    Then the user should see on the page "Test Product"
    And the user clicks on an element which contains "Test Product"
    
    # Verify product details page
    Then the user should be navigated to "details" page
    And the user should see details on the page:
      | element                        | value                   |
      | product_name_element          | Test Product            |
      | product_description_element   | A product for testing   |
      | product_price                 | $29.99                  |

    # Add to cart functionality
    When the user clicks on the "add_to_cart_button" element on "details_page"
    And the user waits for the "success_message" element on "common"
    Then the user should see on the page "Product added to cart"

    # API Verification: Check cart contents
    Given I prepare a request for the "get_orders" endpoint
    And I set the query parameter "user_id" to "${customer_id}"
    And I set the query parameter "status" to "cart"
    When I send the request
    Then the response status code should be 200
    And the response body should contain "${product_id}"

  @api @error-handling
  Scenario: API Error Handling with Enhanced Configuration
    # Test invalid endpoint access
    Given I prepare a request for the "get_user" endpoint
    And I set the path parameter "id" to "99999"
    When I send the request
    Then the response status code should be 404
    And the response body should contain "User not found"

    # Test unauthorized access
    Given I prepare a request for the "delete_user" endpoint
    And I set the path parameter "id" to "1"
    And I set the request header "Authorization" to "Bearer invalid-token"
    When I send the request
    Then the response status code should be 401
    And the response body should contain "Unauthorized"

  @frontend @responsive
  Scenario: Responsive Design Testing with Enhanced Selectors
    # Test on mobile viewport
    When the user sets the browser viewport to 375x667
    And the user goes to the "dashboard" page
    Then the user should see on the page the "navigation_menu" element on "common"
    
    # Test mobile-specific elements with fallback selectors
    When the user clicks on the "navigation_menu" element on "common"
    Then the user should see on the page the "logout_button" element on "common"

    # Test on desktop viewport
    When the user sets the browser viewport to 1920x1080
    And the user refreshes the page
    Then the user should see on the page the "user_profile_dropdown" element on "dashboard_page"
    And the user should see on the page the "stats_cards" element on "dashboard_page"

  @visual @screenshots
  Scenario: Visual Testing with Enhanced Configuration
    # Navigate to visual test page
    When the user goes to the "visual" page
    Then the user should be navigated to "visual" page

    # Test visual elements with fallback selectors
    And the user should see on the page the "visual_test_button" element on "common"
    When the user clicks on the "visual_test_button" element on "common"
    
    # Wait for visual changes
    Then the user waits 2 seconds
    And the user should see on the page "Visual test completed"

  @data-driven @api
  Scenario Outline: Data-Driven API Testing with Multiple Users
    Given I prepare a request for the "create_user" endpoint
    And I set the request body to:
      """
      {
        "name": "<name>",
        "email": "<email>",
        "password": "<password>",
        "role": "<role>"
      }
      """
    When I send the request
    Then the response status code should be <expected_status>
    And the response body should contain "<expected_content>"

    Examples:
      | name           | email                  | password        | role  | expected_status | expected_content |
      | Admin User     | admin@example.com      | AdminPass123    | admin | 201             | success          |
      | Regular User   | user@example.com       | UserPass123     | user  | 201             | success          |
      | Invalid User   | invalid-email          | weak            | user  | 400             | Invalid email    |

  @cleanup
  Scenario: Test Data Cleanup
    # Clean up test data after scenarios
    Given I prepare a request for the "cleanup_test_data" endpoint
    When I send the request
    Then the response status code should be 200
    And the response body should contain "Cleanup completed" 