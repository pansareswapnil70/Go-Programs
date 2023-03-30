package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Json data")
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
	fmt.Println(checkValid)
}
