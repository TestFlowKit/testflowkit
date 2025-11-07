package graphql

import (
	"testing"
)

func TestNewSchemaValidator(t *testing.T) {
	schema := &Schema{
		QueryType: &Type{Name: "Query"},
		Types: []Type{
			{Name: "Query", Kind: "OBJECT"},
		},
	}

	validator := NewSchemaValidator(schema)
	if validator == nil {
		t.Fatal("Expected validator to be created")
	}

	if validator.schema != schema {
		t.Error("Expected validator to use provided schema")
	}
}

func TestValidateOperation_ValidQuery(t *testing.T) {
	schema := &Schema{
		QueryType: &Type{Name: "Query"},
		Types: []Type{
			{
				Name: "Query",
				Kind: "OBJECT",
				Fields: []Field{
					{Name: "user", Type: Type{Kind: "OBJECT", Name: "User"}},
					{Name: "users", Type: Type{Kind: "LIST", OfType: &Type{Kind: "OBJECT", Name: "User"}}},
				},
			},
			{
				Name: "User",
				Kind: "OBJECT",
				Fields: []Field{
					{Name: "id", Type: Type{Kind: "SCALAR", Name: "ID"}},
					{Name: "name", Type: Type{Kind: "SCALAR", Name: "String"}},
				},
			},
		},
	}

	validator := NewSchemaValidator(schema)

	tests := []struct {
		name      string
		operation string
		wantValid bool
	}{
		{
			name:      "simple query",
			operation: "query { user { id name } }",
			wantValid: true,
		},
		{
			name:      "query without operation keyword",
			operation: "{ user { id } }",
			wantValid: true,
		},
		{
			name:      "query with multiple fields",
			operation: "query { user { id name } users { id } }",
			wantValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.ValidateOperation(tt.operation)
			if result.Valid != tt.wantValid {
				t.Errorf("ValidateOperation() valid = %v, want %v", result.Valid, tt.wantValid)
				if !result.Valid {
					for _, err := range result.Errors {
						t.Logf("Validation error: %v", err)
					}
				}
			}
		})
	}
}

func TestValidateOperation_InvalidQuery(t *testing.T) {
	schema := &Schema{
		QueryType: &Type{Name: "Query"},
		Types: []Type{
			{
				Name: "Query",
				Kind: "OBJECT",
				Fields: []Field{
					{Name: "user", Type: Type{Kind: "OBJECT", Name: "User"}},
				},
			},
		},
	}

	validator := NewSchemaValidator(schema)

	tests := []struct {
		name          string
		operation     string
		expectedError ValidationErrorType
	}{
		{
			name:          "field not found",
			operation:     "query { nonexistentField }",
			expectedError: ErrorTypeFieldNotFound,
		},
		{
			name:          "invalid syntax - missing braces",
			operation:     "query user",
			expectedError: ErrorTypeSyntaxError,
		},
		{
			name:          "invalid syntax - unmatched braces",
			operation:     "query { user { id }",
			expectedError: ErrorTypeSyntaxError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.ValidateOperation(tt.operation)
			if result.Valid {
				t.Error("Expected validation to fail")
			}

			found := false
			for _, err := range result.Errors {
				if err.Type == tt.expectedError {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Expected error type %v, got errors: %v", tt.expectedError, result.Errors)
			}
		})
	}
}

func TestValidateOperation_Mutation(t *testing.T) {
	schema := &Schema{
		QueryType:    &Type{Name: "Query"},
		MutationType: &Type{Name: "Mutation"},
		Types: []Type{
			{
				Name: "Query",
				Kind: "OBJECT",
				Fields: []Field{
					{Name: "user", Type: Type{Kind: "OBJECT", Name: "User"}},
				},
			},
			{
				Name: "Mutation",
				Kind: "OBJECT",
				Fields: []Field{
					{Name: "createUser", Type: Type{Kind: "OBJECT", Name: "User"}},
					{Name: "updateUser", Type: Type{Kind: "OBJECT", Name: "User"}},
				},
			},
		},
	}

	validator := NewSchemaValidator(schema)

	tests := []struct {
		name      string
		operation string
		wantValid bool
	}{
		{
			name:      "valid mutation",
			operation: "mutation { createUser { id } }",
			wantValid: true,
		},
		{
			name:      "invalid mutation field",
			operation: "mutation { deleteUser { id } }",
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.ValidateOperation(tt.operation)
			if result.Valid != tt.wantValid {
				t.Errorf("ValidateOperation() valid = %v, want %v", result.Valid, tt.wantValid)
				if !result.Valid {
					for _, err := range result.Errors {
						t.Logf("Validation error: %v", err)
					}
				}
			}
		})
	}
}

func TestValidateOperation_NoSchema(t *testing.T) {
	validator := NewSchemaValidator(nil)
	result := validator.ValidateOperation("query { user }")

	if result.Valid {
		t.Error("Expected validation to fail with no schema")
	}

	if len(result.Errors) == 0 {
		t.Error("Expected validation errors")
	}

	if result.Errors[0].Type != ErrorTypeSchemaUnavailable {
		t.Errorf("Expected error type %v, got %v", ErrorTypeSchemaUnavailable, result.Errors[0].Type)
	}
}

func TestValidateOperation_UnsupportedOperationType(t *testing.T) {
	schema := &Schema{
		QueryType: &Type{Name: "Query"},
		Types: []Type{
			{Name: "Query", Kind: "OBJECT"},
		},
	}

	validator := NewSchemaValidator(schema)

	tests := []struct {
		name      string
		operation string
	}{
		{
			name:      "mutation not supported",
			operation: "mutation { createUser }",
		},
		{
			name:      "subscription not supported",
			operation: "subscription { userUpdated }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.ValidateOperation(tt.operation)
			if result.Valid {
				t.Error("Expected validation to fail for unsupported operation type")
			}

			found := false
			for _, err := range result.Errors {
				if err.Type == ErrorTypeTypeNotFound {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Expected error type %v, got errors: %v", ErrorTypeTypeNotFound, result.Errors)
			}
		})
	}
}

func TestSchemaValidator_ValidateFieldPath(t *testing.T) {
	schema := &Schema{
		QueryType: &Type{Name: "Query"},
		Types: []Type{
			{
				Name: "Query",
				Kind: "OBJECT",
				Fields: []Field{
					{Name: "user", Type: Type{Kind: "OBJECT", Name: "User"}},
				},
			},
			{
				Name: "User",
				Kind: "OBJECT",
				Fields: []Field{
					{Name: "id", Type: Type{Kind: "SCALAR", Name: "ID"}},
					{Name: "profile", Type: Type{Kind: "OBJECT", Name: "Profile"}},
				},
			},
			{
				Name: "Profile",
				Kind: "OBJECT",
				Fields: []Field{
					{Name: "bio", Type: Type{Kind: "SCALAR", Name: "String"}},
				},
			},
		},
	}

	validator := NewSchemaValidator(schema)

	tests := []struct {
		name          string
		path          string
		operationType string
		wantValid     bool
	}{
		{
			name:          "valid simple path",
			path:          "user",
			operationType: "query",
			wantValid:     true,
		},
		{
			name:          "valid nested path",
			path:          "user.profile.bio",
			operationType: "query",
			wantValid:     true,
		},
		{
			name:          "invalid field",
			path:          "nonexistent",
			operationType: "query",
			wantValid:     false,
		},
		{
			name:          "invalid nested field",
			path:          "user.nonexistent",
			operationType: "query",
			wantValid:     false,
		},
		{
			name:          "invalid operation type",
			path:          "user",
			operationType: "invalid",
			wantValid:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.ValidateFieldPath(tt.path, tt.operationType)
			if result.Valid != tt.wantValid {
				t.Errorf("ValidateFieldPath() valid = %v, want %v", result.Valid, tt.wantValid)
				if !result.Valid {
					for _, err := range result.Errors {
						t.Logf("Validation error: %v", err)
					}
				}
			}
		})
	}
}

func TestExtractOperationType(t *testing.T) {
	validator := NewSchemaValidator(&Schema{})

	tests := []struct {
		name      string
		operation string
		expected  string
	}{
		{
			name:      "explicit query",
			operation: "query { user }",
			expected:  "query",
		},
		{
			name:      "explicit mutation",
			operation: "mutation { createUser }",
			expected:  "mutation",
		},
		{
			name:      "explicit subscription",
			operation: "subscription { userUpdated }",
			expected:  "subscription",
		},
		{
			name:      "implicit query",
			operation: "{ user }",
			expected:  "query",
		},
		{
			name:      "query with whitespace",
			operation: "  query   { user }",
			expected:  "query",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := validator.extractOperationType(tt.operation)
			if err != nil {
				t.Errorf("extractOperationType() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("extractOperationType() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractFields(t *testing.T) {
	validator := NewSchemaValidator(&Schema{})

	tests := []struct {
		name      string
		operation string
		expected  []string
	}{
		{
			name:      "simple fields",
			operation: "query { user name }",
			expected:  []string{"user", "name"},
		},
		{
			name:      "nested fields - only root extracted",
			operation: "query { user { id name } }",
			expected:  []string{"user"},
		},
		{
			name:      "fields with duplicates",
			operation: "query { user user name }",
			expected:  []string{"user", "name"},
		},
		{
			name:      "complex query - only root extracted",
			operation: "query GetUser { user(id: 123) { id name profile { bio } } }",
			expected:  []string{"user"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := validator.extractFields(tt.operation)
			if err != nil {
				t.Errorf("extractFields() error = %v", err)
				return
			}

			// Check that all expected fields are present
			for _, expected := range tt.expected {
				found := false
				for _, actual := range result {
					if actual == expected {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("extractFields() missing expected field %v, got %v", expected, result)
				}
			}
		})
	}
}

func TestCleanOperation(t *testing.T) {
	validator := NewSchemaValidator(&Schema{})

	tests := []struct {
		name      string
		operation string
		expected  string
	}{
		{
			name:      "remove comments",
			operation: "query { # this is a comment\n user }",
			expected:  "query { user }",
		},
		{
			name:      "normalize whitespace",
			operation: "query   {\n  user\n  name\n}",
			expected:  "query { user name }",
		},
		{
			name:      "mixed comments and whitespace",
			operation: "query { # get user\n  user # user field\n  name\n}",
			expected:  "query { user name }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.cleanOperation(tt.operation)
			if result != tt.expected {
				t.Errorf("cleanOperation() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestValidationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      ValidationError
		expected string
	}{
		{
			name: "error with field",
			err: ValidationError{
				Type:    ErrorTypeFieldNotFound,
				Message: "Field not found",
				Field:   "user",
			},
			expected: "field_not_found: Field not found (field: user)",
		},
		{
			name: "error without field",
			err: ValidationError{
				Type:    ErrorTypeSchemaUnavailable,
				Message: "Schema not available",
			},
			expected: "schema_unavailable: Schema not available",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			if result != tt.expected {
				t.Errorf("ValidationError.Error() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestValidationResult_AddError(t *testing.T) {
	result := &ValidationResult{Valid: true}

	result.AddError(ErrorTypeFieldNotFound, "Field not found", "user")

	if result.Valid {
		t.Error("Expected Valid to be false after adding error")
	}

	if len(result.Errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(result.Errors))
	}

	err := result.Errors[0]
	if err.Type != ErrorTypeFieldNotFound {
		t.Errorf("Expected error type %v, got %v", ErrorTypeFieldNotFound, err.Type)
	}

	if err.Message != "Field not found" {
		t.Errorf("Expected message 'Field not found', got %q", err.Message)
	}

	if err.Field != "user" {
		t.Errorf("Expected field 'user', got %q", err.Field)
	}
}

func TestValidationResult_AddDetailedError(t *testing.T) {
	result := &ValidationResult{Valid: true}
	details := map[string]interface{}{
		"type":             "Query",
		"available_fields": []string{"user", "posts"},
	}

	result.AddDetailedError(ErrorTypeFieldNotFound, "Field not found", "invalid", details)

	if result.Valid {
		t.Error("Expected Valid to be false after adding error")
	}

	if len(result.Errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(result.Errors))
	}

	err := result.Errors[0]
	if err.Details == nil {
		t.Error("Expected error to have details")
	}

	if err.Details["type"] != "Query" {
		t.Errorf("Expected details type 'Query', got %v", err.Details["type"])
	}
}
