package main

import "fmt"

type User struct {
	id   int
	name string
}

func main() {
	fmt.Println("Welcome to structs in golang")
	Swapnil := User{1, "Swapnil"}
	fmt.Println(Swapnil)
	fmt.Println(Swapnil.id)
	fmt.Printf("Swapnil details are: %+v\n", Swapnil)
}
