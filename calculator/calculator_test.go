package calculator_test

import (
	"testing"

	"github.com/cking-bot/pivottechschool/calculator"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "test one", a: 1, b: 2, want: 3},
		{name: "test two", a: 4, b: 10, want: 14},
		{name: "test three", a: 25, b: 10, want: 35},
	}

	for _, test := range tests {
		//t.Run uses subtest
		t.Run(test.name, func(t *testing.T) {
			got := calculator.Add(test.a, test.b)
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "test one", a: 1, b: 2, want: -1},
		{name: "test two", a: 100, b: 100, want: 0},
		{name: "test three", a: 25, b: 10, want: 15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculator.Subtract(test.a, test.b)
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "test one", a: 1, b: 2, want: 2},
		{name: "test two", a: -25, b: 5, want: -125},
		{name: "test three", a: 25, b: 0, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculator.Multiply(test.a, test.b)
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{name: "test one", a: 2, b: 3, want: 8},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculator.Pow(test.a, test.b)
			if got != test.want {
				t.Errorf("got %g, want %g", got, test.want)
			}

		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
		err  error
	}{
		{name: "test one", a: 4, b: 2, want: 2},
		{name: "test two", a: 25, b: 0, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := calculator.Divide(test.a, test.b)
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
			if test.err != nil {
				if err.Error() != test.err.Error() {
					t.Errorf("got %q, want %q", err, test.err)
				}
			}
		})
	}

}
