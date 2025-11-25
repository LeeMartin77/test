package core_test

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

const baseURL = "http://localhost:8080"

func TestPingEndpoint(t *testing.T) {
	resp, err := http.Get(baseURL + "/ping")
	if err != nil {
		t.Fatalf("Failed to make request to ping endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expected := "pong"
	if string(body) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, string(body))
	}
}

func TestAddEndpointValid(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     string
		expected string
	}{
		{"positive integers", "5", "3", "8.00"},
		{"positive decimals", "2.5", "1.5", "4.00"},
		{"negative numbers", "-5", "3", "-2.00"},
		{"zero values", "0", "5", "5.00"},
		{"large numbers", "999", "1", "1000.00"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + "/add?a=" + tc.a + "&b=" + tc.b)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status 200, got %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if string(body) != tc.expected {
				t.Errorf("Add(%s, %s): expected '%s', got '%s'", tc.a, tc.b, tc.expected, string(body))
			}
		})
	}
}

func TestAddEndpointInvalid(t *testing.T) {
	testCases := []struct {
		name       string
		url        string
		expectCode int
	}{
		{"missing parameter a", "/add?b=5", http.StatusBadRequest},
		{"missing parameter b", "/add?a=5", http.StatusBadRequest},
		{"invalid parameter a", "/add?a=invalid&b=5", http.StatusBadRequest},
		{"invalid parameter b", "/add?a=5&b=invalid", http.StatusBadRequest},
		{"no parameters", "/add", http.StatusBadRequest},
		{"empty parameters", "/add?a=&b=", http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + tc.url)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectCode {
				t.Errorf("Expected status %d, got %d", tc.expectCode, resp.StatusCode)
			}

			// For error cases, just verify we got some error text
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if len(body) == 0 {
				t.Error("Expected error message, got empty response")
			}
		})
	}
}

func TestSubtractEndpointValid(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     string
		expected string
	}{
		{"positive result", "10", "3", "7.00"},
		{"negative result", "3", "10", "-7.00"},
		{"zero result", "5", "5", "0.00"},
		{"decimal numbers", "7.5", "2.3", "5.20"},
		{"negative inputs", "-5", "-3", "-2.00"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + "/sub?a=" + tc.a + "&b=" + tc.b)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status 200, got %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if string(body) != tc.expected {
				t.Errorf("Subtract(%s, %s): expected '%s', got '%s'", tc.a, tc.b, tc.expected, string(body))
			}
		})
	}
}

func TestSubtractEndpointInvalid(t *testing.T) {
	testCases := []struct {
		name       string
		url        string
		expectCode int
	}{
		{"missing parameter a", "/sub?b=5", http.StatusBadRequest},
		{"missing parameter b", "/sub?a=5", http.StatusBadRequest},
		{"invalid parameter a", "/sub?a=text&b=5", http.StatusBadRequest},
		{"invalid parameter b", "/sub?a=5&b=text", http.StatusBadRequest},
		{"no parameters", "/sub", http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + tc.url)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectCode {
				t.Errorf("Expected status %d, got %d", tc.expectCode, resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if len(body) == 0 {
				t.Error("Expected error message, got empty response")
			}
		})
	}
}

func TestMethodNotAllowed(t *testing.T) {
	endpoints := []string{"/ping", "/add", "/sub"}

	for _, endpoint := range endpoints {
		t.Run("POST "+endpoint, func(t *testing.T) {
			resp, err := http.Post(baseURL+endpoint, "application/json", strings.NewReader("{}"))
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusMethodNotAllowed {
				t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
			}
		})
	}
}
