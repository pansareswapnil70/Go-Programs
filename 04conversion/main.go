package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app !")
	fmt.Println("Please enter a rating for pizza:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Hey, You entered the rating ", input)
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Hey something went wrong")
	}
	fmt.Println("Your converted rating is ", num+1)
}
