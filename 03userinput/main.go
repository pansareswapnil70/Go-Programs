package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating", input)
	fmt.Printf("Type of your rating %T\n", input)
}
