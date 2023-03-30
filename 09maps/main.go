package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"
	fmt.Println(languages)
	delete(languages, "RB")
	fmt.Println(languages)
	for _, val := range languages {
		fmt.Printf("For key, value is %v\n", val)
	}
}
