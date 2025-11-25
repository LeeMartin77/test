package test003_test

import (
	"io"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

func TestBusinessLogicEndpointExists(t *testing.T) {
	resp, err := http.Get(baseURL + "/businesslogic?a=10&b=5&c=3&d=2")
	if err != nil {
		t.Fatalf("Server request failed - ensure server is running: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		t.Error("Business logic endpoint not implemented. Expected /businesslogic route to exist")
	}
}

func TestBusinessLogicCalculation(t *testing.T) {
	testCases := []struct {
		name        string
		a, b, c, d  string
		expected    string
		description string
	}{
		{
			"basic calculation",
			"10", "5", "3", "2",
			"2.40",
			"((10 + 5) - 3) * 2 / 10 = 2.40",
		},
		{
			"different values",
			"20", "8", "4", "3",
			"3.60",
			"((20 + 8) - 4) * 3 / 20 = 3.60",
		},
		{
			"decimal inputs",
			"5.0", "2.5", "1.5", "4.0",
			"4.80",
			"((5.0 + 2.5) - 1.5) * 4.0 / 5.0 = 4.80",
		},
		{
			"result with zero",
			"8", "2", "10", "5",
			"0.00",
			"((8 + 2) - 10) * 5 / 8 = 0.00",
		},
		{
			"negative intermediate result",
			"5", "2", "10", "3",
			"-1.80",
			"((5 + 2) - 10) * 3 / 5 = -1.80",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := baseURL + "/businesslogic?a=" + tc.a + "&b=" + tc.b + "&c=" + tc.c + "&d=" + tc.d
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Server request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusNotFound {
				t.Skip("Business logic endpoint not implemented yet")
			}

			if resp.StatusCode != http.StatusOK {
				body, _ := io.ReadAll(resp.Body)
				t.Errorf("Business logic endpoint returned status %d. Response: %s", resp.StatusCode, string(body))
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			result := string(body)
			if result != tc.expected {
				t.Errorf("BusinessLogic(%s,%s,%s,%s): expected '%s', got '%s'. Formula: %s",
					tc.a, tc.b, tc.c, tc.d, tc.expected, result, tc.description)
			}
		})
	}
}

func TestBusinessLogicDivisionByZero(t *testing.T) {
	// Test division by zero when a=0
	resp, err := http.Get(baseURL + "/businesslogic?a=0&b=5&c=3&d=2")
	if err != nil {
		t.Fatalf("Server request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		t.Skip("Business logic endpoint not implemented yet")
	}

	// Should return an error status for division by zero
	if resp.StatusCode != http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		t.Errorf("Division by zero (a=0) should return BadRequest status, got %d. Response: %s",
			resp.StatusCode, string(body))
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

func TestBusinessLogicMissingParameters(t *testing.T) {
	testCases := []struct {
		name       string
		url        string
		expectCode int
	}{
		{"missing a", "/businesslogic?b=5&c=3&d=2", http.StatusBadRequest},
		{"missing b", "/businesslogic?a=10&c=3&d=2", http.StatusBadRequest},
		{"missing c", "/businesslogic?a=10&b=5&d=2", http.StatusBadRequest},
		{"missing d", "/businesslogic?a=10&b=5&c=3", http.StatusBadRequest},
		{"missing multiple", "/businesslogic?a=10&b=5", http.StatusBadRequest},
		{"no parameters", "/businesslogic", http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + tc.url)
			if err != nil {
				t.Fatalf("Server request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusNotFound {
				t.Skip("Business logic endpoint not implemented yet")
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

func TestBusinessLogicInvalidParameters(t *testing.T) {
	testCases := []struct {
		name string
		url  string
	}{
		{"invalid a", "/businesslogic?a=invalid&b=5&c=3&d=2"},
		{"invalid b", "/businesslogic?a=10&b=invalid&c=3&d=2"},
		{"invalid c", "/businesslogic?a=10&b=5&c=invalid&d=2"},
		{"invalid d", "/businesslogic?a=10&b=5&c=3&d=invalid"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseURL + tc.url)
			if err != nil {
				t.Fatalf("Server request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusNotFound {
				t.Skip("Business logic endpoint not implemented yet")
			}

			if resp.StatusCode != http.StatusBadRequest {
				body, _ := io.ReadAll(resp.Body)
				t.Errorf("Invalid parameter should return BadRequest, got %d. Response: %s",
					resp.StatusCode, string(body))
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
