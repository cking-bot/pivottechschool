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
