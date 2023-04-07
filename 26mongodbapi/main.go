package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/pansareswapnil70/mongodbapi/Router"
)

func main() {
	fmt.Println("MONGO DB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Started Listening on port 4000")
}
