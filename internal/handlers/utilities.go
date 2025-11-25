package handlers

import (
	"errors"
	"net/http"
	"strconv"
)

// ParseQueryParams extracts and validates 'a' and 'b' query parameters
func ParseQueryParams(r *http.Request) (*float64, *float64, error) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	if aStr == "" || bStr == "" {
		return nil, nil, errors.New("both 'a' and 'b' query parameters are required")
	}

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return nil, nil, errors.New("parameter 'a' must be a valid number")
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return nil, nil, errors.New("parameter 'b' must be a valid number")
	}

	return &a, &b, nil
}
