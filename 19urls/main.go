package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gjhygfd"

func main() {
	fmt.Println("Learning to handle urls in go")
	fmt.Println(myurl)
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryparams := result.Query()
	fmt.Printf("Query params are of type %T \n", queryparams)
	fmt.Println(queryparams["coursename"])
	fmt.Println(queryparams["paymentid"])

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "?coursename=reactjs",
	}
	anotherurl := partsOfUrl.String()
	fmt.Println(anotherurl)

}
