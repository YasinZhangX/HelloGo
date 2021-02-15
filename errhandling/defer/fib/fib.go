package fib

import (
	"fmt"
	"strings"
)

// Fibonacci generator
func Fibonacci() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// IntGen Generator function
type IntGen func() int

func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
