package main

import (
	"fmt"
	"sync"
)

func main() {
	myChan := make(chan int, 1)
	wg := &sync.WaitGroup{}
	fmt.Println("Channels in Golang")
	wg.Add(2)
	// Read Only Channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChan
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myChan, wg)
	// Write Only Channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChan <- 5
		close(myChan)
		wg.Done()
	}(myChan, wg)
	wg.Wait()
}
