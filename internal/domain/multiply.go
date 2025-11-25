package domain

import "log"

func (m *mathService) Multiply(a, b float64) float64 {
	log.Printf("Debug: b value is %f", b)

	return a * a
}
