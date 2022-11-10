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

type ErrDivideByZero string

func (err ErrDivideByZero) Error() string {
	return string(err)
}

func Divide(a, b int) (int, error) {
	zero := b == 0

	if zero {
		return 0, ErrDivideByZero("divide by zero")
	}
	return a / b, nil
}
