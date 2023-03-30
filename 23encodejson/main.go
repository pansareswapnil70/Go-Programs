package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to encode json")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcocourses := []course{
		{"Reactjs", 399, "lcoonline.com", "abc123", []string{"web-dev"}},
		{"Golang", 399, "lcoonline.com", "bcd123", []string{"dev"}},
		{"Python", 2399, "lcoonline.com", "xyz123", nil},
	}
	fmt.Printf("%v\n", lcocourses)

	finaljson, _ := json.Marshal(lcocourses)
	fmt.Printf("%s\n", finaljson)
	Indentedfinaljson, _ := json.MarshalIndent(lcocourses, "", "\t")
	fmt.Printf("%s\n", Indentedfinaljson)
}

func DecodeJson() {
	datafromweb := []byte(`
		{
                "coursename": "Reactjs",
                "Price": 399,
                "website": "lcoonline.com",
                "tags": [ "web-dev" ]
        }
	`)
	validatejson := json.Valid(datafromweb)
	fmt.Println(validatejson)
}
