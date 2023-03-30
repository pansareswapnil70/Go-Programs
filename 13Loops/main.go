package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days)
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Print("\n")
	for i := range days {
		fmt.Println(days[i])
	}
	for index, value := range days {
		if index == 2 {
			continue
		}
		fmt.Printf("For key %v, value is %v\n", index, value)
	}
}
