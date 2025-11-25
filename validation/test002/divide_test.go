package test002_test

import (
	"io"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

func TestDivideEndpointExists(t *testing.T) {
	resp, err := http.Get(baseURL + "/div?a=10&b=2")
	if err != nil {
		t.Fatalf("Server request failed - ensure server is running: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		t.Error("Divide endpoint not implemented. Expected /div route to exist")
	}
}

func TestDivideEndpointValid(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     string
		expected string
	}{
		{"simple division", "10", "2", "5.00"},
		{"decimal result", "7", "2", "3.50"},
		{"division by one", "42", "1", "42.00"},
		{"small numbers", "1", "4", "0.25"},
		{"negative dividend", "-15", "3", "-5.00"},
		{"negative divisor", "15", "-3", "-5.00"},
		{"both negative", "-12", "-4", "3.00"},
		{"decimal inputs", "7.5", "2.5", "3.00"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + "/div?a=" + tc.a + "&b=" + tc.b)
			if err != nil {
				t.Fatalf("Server request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusNotFound {
				t.Skip("Divide endpoint not implemented yet")
			}

			if resp.StatusCode != http.StatusOK {
				body, _ := io.ReadAll(resp.Body)
				t.Errorf("Divide endpoint returned status %d. Response: %s", resp.StatusCode, string(body))
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			result := string(body)
			if result != tc.expected {
				t.Errorf("Divide(%s, %s): expected '%s', got '%s'", tc.a, tc.b, tc.expected, result)
			}
		})
	}
}

func TestDivideByZero(t *testing.T) {
	resp, err := http.Get(baseURL + "/div?a=10&b=0")
	if err != nil {
		t.Fatalf("Server request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		t.Skip("Divide endpoint not implemented yet")
	}

	// Should return an error status (400) for division by zero
	if resp.StatusCode != http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		t.Errorf("Division by zero should return BadRequest status, got %d. Response: %s", resp.StatusCode, string(body))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if len(body) == 0 {
		t.Error("Division by zero should return an error message")
	}
}

func TestDivideEndpointInvalid(t *testing.T) {
	testCases := []struct {
		name       string
		url        string
		expectCode int
	}{
		{"missing parameter a", "/div?b=5", http.StatusBadRequest},
		{"missing parameter b", "/div?a=5", http.StatusBadRequest},
		{"invalid parameter a", "/div?a=invalid&b=5", http.StatusBadRequest},
		{"invalid parameter b", "/div?a=5&b=invalid", http.StatusBadRequest},
		{"no parameters", "/div", http.StatusBadRequest},
		{"empty parameters", "/div?a=&b=", http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + tc.url)
			if err != nil {
				t.Fatalf("Server request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusNotFound {
				t.Skip("Divide endpoint not implemented yet")
			}

			if resp.StatusCode != tc.expectCode {
				body, _ := io.ReadAll(resp.Body)
				t.Errorf("Expected status %d, got %d. Response: %s", tc.expectCode, resp.StatusCode, string(body))
				return
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

func TestDivideEndpointFollowsPattern(t *testing.T) {
	// Test that divide endpoint behaves like other math endpoints
	resp, err := http.Get(baseURL + "/div?a=20&b=4")
	if err != nil {
		t.Fatalf("Server request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		t.Skip("Divide endpoint not implemented yet")
	}

	// Check content type
	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/plain" {
		t.Errorf("Expected Content-Type 'text/plain', got '%s'", contentType)
	}

	// Check that successful response returns formatted number
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		result := string(body)
		if result != "5.00" {
			t.Errorf("Expected formatted result '5.00', got '%s'", result)
		}
	}
}
