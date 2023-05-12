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
	fmt.Println("Welcome to Json data")
	lcocourse := []course{
		{"ReactJs", 499, "lco.dev", "abc123", []string{"web-dev", "front-end"}},
		{"Golang", 399, "linkedIn.com", "xyz123", nil},
	}
	fmt.Println(lcocourse)
	finalJson, _ := json.MarshalIndent(&lcocourse, "", "\t")
	fmt.Println(string(finalJson))
	DecodeJson()
}

func DecodeJson() {
	jsondatafromweb := []byte(`
	   {"menu": {
  "id": "file",
  "value": "File",
  "popup": {
    "menuitem": [
      {"value": "New", "onclick": "CreateNewDoc()"},
      {"value": "Open", "onclick": "OpenDoc()"},
      {"value": "Close", "onclick": "CloseDoc()"}
    ]
  }
}}
	`)

	checkValid := json.Valid(jsondatafromweb)
	if checkValid {
		fmt.Println("Json is valid")
	}
}
