package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectErr   bool
	}{
		{
			name: "Valid APIKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey 1234abcd"},
			},
			expectedKey: "1234abcd",
			expectErr:   false,
		},
		{
			name: "Missing Authorization header",
			headers: http.Header{
				"Content-Type": []string{"application/json"},
			},
			expectedKey: "",
			expectErr:   true,
		},
		{
			name: "Malformed Authorization header 1",
			headers: http.Header{
				"Authorization": []string{"Bearer 12345"},
			},
			expectedKey: "",
			expectErr:   true,
		},
		{
			name: "Malformed Authorization header 2",
			headers: http.Header{
				"Authorization": []string{"12345"},
			},
			expectedKey: "",
			expectErr:   true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)
			if (err != nil) != tc.expectErr {
				t.Errorf("expect error: %v, got: %v", tc.expectErr, err)
			}
			if apiKey != tc.expectedKey {
				t.Errorf("expected APIKey: %s, got: %s", tc.expectedKey, apiKey)
			}
		})
	}
}
