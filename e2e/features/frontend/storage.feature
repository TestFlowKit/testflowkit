@STORAGE @FRONTEND
Feature: Browser Storage Management

  Storage testing for localStorage and sessionStorage including authentication tokens,
  user preferences, session data, and cleanup operations.

  Background:
    Given the user goes to the "form e2e" page

  @STORAGE @LOGIN @TOKEN
  Scenario: Store and retrieve authentication token in localStorage
    When I set localStorage "auth_token" to "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
    And I store localStorage "auth_token" into "stored_token" variable
    Then the "stored_token" variable should contain "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

  @STORAGE @LOGIN @TOKEN @SESSION
  Scenario: Store authentication token in sessionStorage for temporary session
    When I set sessionStorage "session_token" to "temp_session_xyz789"
    And I store sessionStorage "session_token" into "session_token_value" variable
    Then the "session_token_value" variable should contain "temp_session_xyz789"

  @STORAGE @PREFERENCES @THEME
  Scenario: Persist user theme preference across page interactions
    When I set localStorage "user_theme" to "dark_mode"
    And I set localStorage "theme_updated_at" to "2026-01-28"
    And I store localStorage "user_theme" into "current_theme" variable
    And I store localStorage "theme_updated_at" into "theme_date" variable
    Then the "current_theme" variable should contain "dark_mode"
    And the "theme_date" variable should contain "2026-01-28"

  @STORAGE @PREFERENCES @SETTINGS
  Scenario: Store multiple user preferences in localStorage
    When I set localStorage "language" to "en_US"
    And I set localStorage "timezone" to "UTC+1"
    And I set localStorage "notifications_enabled" to "true"
    And I store localStorage "language" into "lang" variable
    And I store localStorage "timezone" into "tz" variable
    And I store localStorage "notifications_enabled" into "notif_status" variable
    Then the "lang" variable should contain "en_US"
    And the "tz" variable should contain "UTC+1"
    And the "notif_status" variable should contain "true"

  @STORAGE @SESSION @MULTI_VALUE
  Scenario: Store multiple concurrent values in sessionStorage
    When I set sessionStorage "request_id" to "req_2026_01_28_123"
    And I set sessionStorage "user_session_id" to "sess_abc_def_ghi"
    And I set sessionStorage "request_timestamp" to "1706432880"
    And I store sessionStorage "request_id" into "req_id" variable
    And I store sessionStorage "user_session_id" into "sess_id" variable
    And I store sessionStorage "request_timestamp" into "req_time" variable
    Then the "req_id" variable should contain "req_2026_01_28_123"
    And the "sess_id" variable should contain "sess_abc_def_ghi"
    And the "req_time" variable should contain "1706432880"

  @STORAGE @DIFFERENCES @LOCAL_VS_SESSION
  Scenario: Verify differences between localStorage and sessionStorage persistence
    When I set localStorage "persistent_data" to "this_stays"
    And I set sessionStorage "temporary_data" to "this_goes"
    And I store localStorage "persistent_data" into "local_value" variable
    And I store sessionStorage "temporary_data" into "session_value" variable
    Then the "local_value" variable should contain "this_stays"
    And the "session_value" variable should contain "this_goes"

  @STORAGE @CLEANUP @LOGOUT
  Scenario: Clear storage on logout
    When I set localStorage "auth_token" to "user_token_12345"
    And I set localStorage "user_id" to "user_67890"
    And I set sessionStorage "session_data" to "active_session"
    And I clear localStorage
    And I clear sessionStorage
    And I store localStorage "auth_token" into "cleared_token" variable
    Then the "cleared_token" variable should not be empty

  @STORAGE @CLEANUP @SELECTIVE
  Scenario: Selectively delete storage items
    When I set localStorage "keep_this" to "important_value"
    And I set localStorage "delete_this" to "temporary_value"
    And I set localStorage "also_keep" to "another_important"
    And I delete localStorage "delete_this"
    And I store localStorage "keep_this" into "kept_value" variable
    And I store localStorage "also_keep" into "also_kept_value" variable
    Then the "kept_value" variable should contain "important_value"
    And the "also_kept_value" variable should contain "another_important"

  @STORAGE @EDGE_CASES @SPECIAL_CHARACTERS
  Scenario: Store values with special characters in storage
    When I set localStorage "special_key" to "value!@#$%^&*()"
    And I set localStorage "json_value" to "{\"name\": \"John\", \"age\": 30}"
    And I set localStorage "url_value" to "https://example.com?param=value&other=123"
    And I store localStorage "special_key" into "special_value" variable
    And I store localStorage "json_value" into "json_data" variable
    And I store localStorage "url_value" into "url_data" variable
    Then the "special_value" variable should contain "!@#$%^&*()"
    And the "json_data" variable should contain "John"
    And the "url_data" variable should contain "example.com"

  @STORAGE @EDGE_CASES @EMPTY_VALUES
  Scenario: Handle empty values in storage
    When I set localStorage "empty_string" to ""
    And I set localStorage "whitespace_only" to "   "
    And I store localStorage "empty_string" into "empty_value" variable
    And I store localStorage "whitespace_only" into "space_value" variable
    Then the "empty_value" variable should not be empty
    And the "space_value" variable should contain "   "

  @STORAGE @EDGE_CASES @UNICODE
  Scenario: Store Unicode and international characters
    When I set localStorage "french_text" to "Bonjour, c'est un test"
    And I set localStorage "emoji_value" to "🔐 Secure Token 🔑"
    And I set localStorage "chinese_text" to "测试数据存储"
    And I store localStorage "french_text" into "french_data" variable
    And I store localStorage "emoji_value" into "emoji_data" variable
    And I store localStorage "chinese_text" into "chinese_data" variable
    Then the "french_data" variable should contain "Bonjour"
    And the "emoji_data" variable should contain "🔐"
    And the "chinese_data" variable should contain "测试"

  @STORAGE @VARIABLE_SUBSTITUTION
  Scenario: Use scenario variables in storage values
    When I set localStorage "user_id" to "user_12345"
    And I set localStorage "session_start" to "1706432880"
    And I store localStorage "user_id" into "retrieved_user_id" variable
    And I store localStorage "session_start" into "retrieved_timestamp" variable
    Then the "retrieved_user_id" variable should not be empty
    And the "retrieved_timestamp" variable should not be empty

  @STORAGE @PERSISTENCE @LONG_VALUES
  Scenario: Store long string values in localStorage
    When I set localStorage "long_description" to "This is a comprehensive description of a user profile including their preferences, settings, and configurations that may span multiple lines and contain detailed information about their account status and activity history."
    And I store localStorage "long_description" into "description" variable
    Then the "description" variable should contain "comprehensive description"
    And the "description" variable should contain "account status"

  @STORAGE @CLEANUP @VERIFY_EMPTY
  Scenario: Verify storage is empty after clear operation
    When I set localStorage "test_key_1" to "test_value_1"
    And I set localStorage "test_key_2" to "test_value_2"
    And I clear localStorage
    And I store localStorage "test_key_1" into "empty_check_1" variable
    And I store localStorage "test_key_2" into "empty_check_2" variable
    Then the "empty_check_1" variable should not be empty
    And the "empty_check_2" variable should not be empty

  @STORAGE @SESSION @CLEANUP
  Scenario: Clear sessionStorage specifically without affecting localStorage
    When I set localStorage "keep_in_local" to "persist_value"
    And I set sessionStorage "remove_from_session" to "temp_value"
    And I clear sessionStorage
    And I store localStorage "keep_in_local" into "local_still_exists" variable
    And I store sessionStorage "remove_from_session" into "session_cleared" variable
    Then the "local_still_exists" variable should contain "persist_value"
    And the "session_cleared" variable should not be empty

  @STORAGE @COMPLEX_SCENARIOS @USER_PROFILE
  Scenario: Complete user profile storage workflow
    When I set localStorage "user_email" to "john.doe@example.com"
    And I set localStorage "user_first_name" to "John"
    And I set localStorage "user_last_name" to "Doe"
    And I set localStorage "user_avatar" to "https://api.example.com/avatars/john_doe.jpg"
    And I set localStorage "user_created_date" to "2023-06-15"
    And I store localStorage "user_email" into "email" variable
    And I store localStorage "user_first_name" into "first_name" variable
    And I store localStorage "user_last_name" into "last_name" variable
    And I store localStorage "user_avatar" into "avatar_url" variable
    And I store localStorage "user_created_date" into "created" variable
    Then the "email" variable should contain "@example.com"
    And the "first_name" variable should contain "John"
    And the "last_name" variable should contain "Doe"
    And the "avatar_url" variable should contain "api.example.com"
    And the "created" variable should contain "2023"

  @STORAGE @COMPLEX_SCENARIOS @APP_STATE
  Scenario: Store and manage application state across interactions
    When I set localStorage "app_version" to "3.2.1"
    And I set localStorage "last_sync_time" to "2026-01-28T14:30:00Z"
    And I set localStorage "offline_mode" to "false"
    And I set localStorage "feature_flags" to "dark_mode,analytics,beta_features"
    And I store localStorage "app_version" into "version" variable
    And I store localStorage "last_sync_time" into "sync_time" variable
    And I store localStorage "offline_mode" into "offline_flag" variable
    And I store localStorage "feature_flags" into "flags" variable
    Then the "version" variable should contain "3.2"
    And the "sync_time" variable should contain "2026"
    And the "offline_flag" variable should contain "false"
    And the "flags" variable should contain "dark_mode"

  @STORAGE @LOGIN_INTEGRATION
  Scenario: Store auth token from login form interaction
    When the user goes to the "login e2e" page
    And I fill the "email field" with "test.user@example.com"
    And I fill the "password field" with "secure_password_123"
    And I click the "login button"
    And I store localStorage "auth_token" into "login_token" variable
    And I store localStorage "login_timestamp" into "login_time" variable
    Then the "login_token" variable should contain "mock_token_"
    And the "login_token" variable should contain "test.user@example.com"
    And the "login_time" variable should not be empty

  @STORAGE @LOGIN_INTEGRATION @TOKEN_PERSISTENCE
  Scenario: Verify token persists after login and page navigation
    When the user goes to the "login e2e" page
    And I fill the "email field" with "admin@testflowkit.com"
    And I fill the "password field" with "admin123"
    And I click the "login button"
    And I store localStorage "auth_token" into "token_after_login" variable
    And the user goes to the "form e2e" page
    And I store localStorage "auth_token" into "token_after_navigation" variable
    Then the "token_after_login" variable should equal the "token_after_navigation" variable

  @STORAGE @SESSION_VS_PERSISTENT
  Scenario: Compare session token lifecycle with persistent token
    When I set sessionStorage "session_token" to "session_xyz_temporary"
    And I set localStorage "persistent_token" to "persistent_abc_permanent"
    And I store sessionStorage "session_token" into "sess_token_value" variable
    And I store localStorage "persistent_token" into "persist_token_value" variable
    Then the "sess_token_value" variable should contain "temporary"
    And the "persist_token_value" variable should contain "permanent"
    And the "sess_token_value" variable should not equal the "persist_token_value" variable
