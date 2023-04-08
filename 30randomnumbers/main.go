package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome to random numbers in golang")
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5))
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNumber)
}
