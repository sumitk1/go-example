package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println(foo(12, 11))
	fmt.Println(multiNamedReturn(100, "last", "first"))

	//var x, y int = 1, 2
	//z := 3 // implicit int

	fmt.Printf("Type: %T Value: %v\n", someBool, someBool)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", someComplex, someComplex)
}

func foo(x, y int) int {
	return x + y
}

func multiNamedReturn(x int, y, z string) (a string, b string, c int) {
	a = y
	b = z
	c = x
	return
}

var (
	someBool    bool       = false
	maxInt      uint64     = 1<<64 - 1
	someComplex complex128 = cmplx.Sqrt(-5 + 12i)
)
