package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to pointers in golang")

	var ptr *int
	fmt.Println("Value of pointer is: ", ptr)

	num := 10
	ptr = &num
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of number is ", *ptr)

	*ptr = *ptr * 5
	fmt.Println("New value of pointer is ", *ptr)
}
