package fib

import (
	"testing"
)

func TestFib(t *testing.T)  {
	type Case struct {
		n int
		answer int
	}
	var cases = []Case{
		{0, 0},
		{1, 1},
		{10, 55},
	}
	for _, c := range cases {

	}
}
