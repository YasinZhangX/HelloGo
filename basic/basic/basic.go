package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "a"

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d, %q\r\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	a := cmplx.Pow(math.E, 1i*math.Pi)
	b := 1 + 0i
	c := cmplx.Exp(1i * math.Pi)
	fmt.Println(a + b)
	fmt.Println(c + b)
	fmt.Printf("%.3f\r\n", a+b)
}

func triganle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64((a*a + b*b))))
	return c
}

func consts() {
	const filename = "abs.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss)
	euler()
	triganle()
	consts()
	enums()
}
