package main

import "fmt"

const LoginToken string = "hfuefhweu"

func main() {
	fmt.Println("Variables")
	var username string = "Swapnil"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T\n", username)
	var number int = 10
	fmt.Println(number)
	fmt.Printf("Variable is of type : %T\n", number)
	var smallnumber uint8 = 255
	fmt.Println(smallnumber)
	fmt.Printf("Variable is of type : %T\n", smallnumber)
	var floatNumber float32 = 255.45666666666
	fmt.Println(floatNumber)
	fmt.Printf("Variable is of type : %T\n", floatNumber)
	var longfloatNumber float64 = 255.45666666666
	fmt.Println(longfloatNumber)
	fmt.Printf("Variable is of type : %T\n", longfloatNumber)
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T\n", isLoggedIn)
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type %T\n", anothervariable)
	var website = "Learncodeonline"
	fmt.Println(website)
	fmt.Printf("Variable is of type %T\n", website)
	numberofusers := 78.76
	fmt.Println(numberofusers)
	fmt.Printf("Variable is of type %T\n", numberofusers)
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T\n", LoginToken)

}
