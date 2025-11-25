package domain

type MathService interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
}

type mathService struct{}

func NewMathService() MathService {
	return &mathService{}
}
