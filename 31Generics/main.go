package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Num interface {
	int | int8 | int16 | int64 | float32 | float64
}

func main() {
	fmt.Println("Welcome to Generics in GO")
	fmt.Println(addition1(4, 7))
	fmt.Println(addition2(4.3, 7.8))
	fmt.Println(addition3(244.78, 344.65))
	fmt.Println(addition3(256, 1000))
	fmt.Println(SimpleGeneric(5.555))
	fmt.Println(SimpleGeneric("Kiran"))
	fmt.Println(SimpleGeneric(45678.98345))
}

// Generic Function by constrained by types int & float64
func addition1[T int | float64](a T, b T) T {
	return a + b
}

// Generic Function constrained by types defined in Num type interface
func addition2[T Num](a T, b T) T {
	return a + b
}

// Generic Function constrained by types mentioned in constraints package
func addition3[T constraints.Ordered](a T, b T) T {
	return a + b
}

// Generic function constrained by any type any == interface{}
func SimpleGeneric[T any](a T) T {
	return a
}
