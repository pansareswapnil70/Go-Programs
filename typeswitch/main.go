package main

import (
	"fmt"
)

func main() {
	like("abc")
	like(80)
}

func like(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println(i.(int))
	case string:
		fmt.Println(i.(string))
	default:
		fmt.Println("Invalid type")
	}
}
