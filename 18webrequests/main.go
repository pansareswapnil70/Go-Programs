package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Learn to handle web requests")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of %T\n", response)
	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}
