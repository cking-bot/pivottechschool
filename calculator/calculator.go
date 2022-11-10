package calculator

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

type ErrDivideByZero struct{}

func (ErrDivideByZero) Error() string {
	return "divide by zero"
}

func Divide(a, b int) (int, error) {
	zero := b == 0

	if zero {
		return 0, ErrDivideByZero{}
	}
	return a / b, nil
}
