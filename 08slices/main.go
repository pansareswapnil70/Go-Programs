package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices in golang")
	var fruitList = []string{"Apple", "Banana", "Orange"}
	fmt.Printf("FruitList of %v Type of fruitlist is %T\n", fruitList, fruitList)
	fruitList = append(fruitList, "mango", "peach", "Kiwi")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	var courses = []string{"React", "Golang", "C", "Java", "Python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
