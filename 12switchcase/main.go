package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch & Case in golang")
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	//fmt.Println("Number on dice is :", dice)

	switch dice {
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
	case 4:
		fmt.Println("Dice value is 4")
	case 5:
		fmt.Println("Dice value is 5")
	case 6:
		fmt.Println("Dice value is 6")
	}
}
