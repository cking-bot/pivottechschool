package calculator

import "log"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

type DivideError string

func (err DivideError) Error() string {
	return string(err)
}

func Divide(a, b int) (int, error) {
	err := b == 0

	if err {
		log.Printf("Error dividing by zero, %v", err)

	}
	return a / b, nil
}
