package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	result := add(3, 4)
	fmt.Println(result)
	proresult := proadder(3, 5, 7, 2, 5, 3)
	fmt.Println(proresult)
}
func add(a int, b int) int {
	return a + b
}

func proadder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
