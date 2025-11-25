package test001_test

import (
	"io"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

func TestMultiplyEndpointValid(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     string
		expected string
	}{
		{"basic multiplication", "5", "3", "15.00"},
		{"multiplication with decimals", "2.5", "4", "10.00"},
		{"multiplication by zero", "7", "0", "0.00"},
		{"multiplication by one", "9", "1", "9.00"},
		{"negative multiplication", "-3", "4", "-12.00"},
		{"both negative", "-2", "-5", "10.00"},
		{"large numbers", "100", "50", "5000.00"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + "/mul?a=" + tc.a + "&b=" + tc.b)
			if err != nil {
				t.Fatalf("Server request failed - ensure server is running: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				body, _ := io.ReadAll(resp.Body)
				t.Errorf("Multiply endpoint returned status %d instead of 200. Response: %s", resp.StatusCode, string(body))
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			result := string(body)
			if result != tc.expected {
				t.Errorf("Multiply(%s, %s): expected '%s', got '%s'. Check multiplication logic.", tc.a, tc.b, tc.expected, result)
			}
		})
	}
}

func TestMultiplyEndpointInvalid(t *testing.T) {
	testCases := []struct {
		name       string
		url        string
		expectCode int
	}{
		{"missing parameter a", "/mul?b=5", http.StatusBadRequest},
		{"missing parameter b", "/mul?a=5", http.StatusBadRequest},
		{"invalid parameter a", "/mul?a=notanumber&b=5", http.StatusBadRequest},
		{"invalid parameter b", "/mul?a=5&b=notanumber", http.StatusBadRequest},
		{"no parameters", "/mul", http.StatusBadRequest},
		{"empty parameters", "/mul?a=&b=", http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + tc.url)
			if err != nil {
				t.Fatalf("Server request failed - ensure server is running: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectCode {
				body, _ := io.ReadAll(resp.Body)
				t.Errorf("Expected status %d, got %d. Response: %s. Check error handling.", tc.expectCode, resp.StatusCode, string(body))
				return
			}

			// Verify we got some error response
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if len(body) == 0 {
				t.Error("Expected error message in response body, got empty response")
			}
		})
	}
}

func TestMultiplyEndpointExists(t *testing.T) {
	resp, err := http.Get(baseURL + "/mul?a=1&b=1")
	if err != nil {
		t.Fatalf("Server request failed - ensure server is running: %v", err)
	}
	defer resp.Body.Close()

	// We expect either a valid response OR a server error, but not 404
	if resp.StatusCode == http.StatusNotFound {
		t.Error("Multiply endpoint not found. Ensure /mul route is registered in main.go")
	}
}

func TestServerPanicRecovery(t *testing.T) {
	// Test basic request that might cause panic due to error handling bugs
	resp, err := http.Get(baseURL + "/mul?a=2&b=3")
	if err != nil {
		t.Fatalf("Server request failed - this might indicate a server crash/panic: %v", err)
	}
	defer resp.Body.Close()

	// If we get here without the test hanging or failing, the server didn't crash
	// We don't care about the exact response, just that the server handled it gracefully
	if resp.StatusCode >= 500 {
		body, _ := io.ReadAll(resp.Body)
		t.Errorf("Server returned 5xx error, possible panic: %d. Response: %s", resp.StatusCode, string(body))
	}
}
