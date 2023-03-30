package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang")
	num := 10
	if num > 10 {
		fmt.Println("Number is greater than 10")
	} else if num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("NUmber is equal to 10")
	}
}
