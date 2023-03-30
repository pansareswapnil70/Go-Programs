package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer function")
	defer fmt.Println("One")
	defer fmt.Println("three")
	fmt.Println("two")
	mydefer()
}
func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
