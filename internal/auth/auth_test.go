package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}

	// Test case: No Authorization header
	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}

	// Test case: Malformed Authorization header
	headers.Set("Authorization", "Bearer some_token")
	_, err = GetAPIKey(headers)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("expected malformed authorization header error, got %v", err)
	}

	// Test case: Correct Authorization header
	headers.Set("Authorization", "ApiKey my_secret_api_key")
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if apiKey != "my_secret_api_keyXXX" {
		t.Errorf("expected api key 'my_secret_api_key', got '%s'", apiKey)
	}
}
