package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// go greeter("hello")
	// greeter("world")

	websiteList := []string{
		"https://www.google.com",
		"https://www.lco.dev",
		"https://go.dev",
		"https://www.fb.com",
		"https://www.github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops something went wrong")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for website %s\n", res.StatusCode, endpoint)
	}
}
