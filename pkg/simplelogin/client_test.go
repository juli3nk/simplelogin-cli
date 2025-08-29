package simplelogin

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		apiKey  string
		wantErr bool
	}{
		{
			name:    "valid client",
			url:     "https://api.example.com",
			apiKey:  "test-key",
			wantErr: false,
		},
		{
			name:    "empty URL",
			url:     "",
			apiKey:  "test-key",
			wantErr: true,
		},
		{
			name:    "empty API key",
			url:     "https://api.example.com",
			apiKey:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(nil, tt.apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && client == nil {
				t.Error("NewClient() returned nil client when no error expected")
			}
		})
	}
}

func TestValidationError_Error(t *testing.T) {
	err := &ValidationError{
		Field:   "test_field",
		Message: "test message",
	}
	expected := "validation error for field 'test_field': test message"
	if err.Error() != expected {
		t.Errorf("ValidationError.Error() = %v, want %v", err.Error(), expected)
	}
}

func TestAPIError_Error(t *testing.T) {
	tests := []struct {
		name     string
		apiError *APIError
		expected string
	}{
		{
			name: "with message",
			apiError: &APIError{
				StatusCode: 400,
				Message:    "Bad Request",
				Body:       "error body",
			},
			expected: "API error (status 400): Bad Request",
		},
		{
			name: "without message",
			apiError: &APIError{
				StatusCode: 500,
				Body:       "Internal Server Error",
			},
			expected: "API error (status 500): Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.apiError.Error() != tt.expected {
				t.Errorf("APIError.Error() = %v, want %v", tt.apiError.Error(), tt.expected)
			}
		})
	}
}
