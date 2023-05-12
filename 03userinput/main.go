package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var username string
	var age int
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	fmt.Println("Enter your text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 := scanner.Text()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza:")
	input, _ := reader.ReadString('\n')
	//Other way of getting user input
	fmt.Println("Enter your name:")
	fmt.Scan(&username)
	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	fmt.Println("Thanks for your rating", input)
	fmt.Printf("Your entered details are: Name: %v and Age: %v\n", username, age)
	fmt.Printf("Type of your rating %T\n", input)

	fmt.Printf("Entered input is %v", input1)
}
