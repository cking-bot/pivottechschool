package main

import (
	"fmt"

	"github.com/cking-bot/pivottechschool/calculator"
)

func main() {

	fmt.Printf("Add(1, 2) = %d\n", calculator.Add(5, 10))
	fmt.Printf("Subtract(10, 1) = %d\n", calculator.Subtract(10, 1))
	fmt.Printf("Multiply(10, 2) = %d\n", calculator.Multiply(10, 2))
	fmt.Printf("Pow(2,3) = %v\n", calculator.Pow(2, 3))

	r, err := calculator.Divide(10, 5)
	fmt.Printf("Divide(10, 5) = %d, %v\n", r, err)

	r, err = calculator.Divide(10, 0)
	fmt.Printf("Divide(10,0) = %d, %v\n", r, err)
}
