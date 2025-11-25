package domain

// Subtract performs subtraction of b from a (a - b)
func (m *mathService) Subtract(a, b float64) float64 {
	return a - b
}
