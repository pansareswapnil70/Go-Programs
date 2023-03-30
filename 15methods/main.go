package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in golang")
	swapnil := User{"Swapnil", "swapnil@gmail.com", true}
	fmt.Println(swapnil)
	swapnil.GetStatus()
	swapnil.NewEmail()
}

type User struct {
	name   string
	email  string
	status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.status)
}
func (u User) NewEmail() {
	u.email = "test@go.dev"
	fmt.Println("Email of user is: ", u.email)
}
